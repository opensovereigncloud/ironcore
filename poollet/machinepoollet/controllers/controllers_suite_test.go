// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/onmetal/controller-utils/buildutils"
	"github.com/onmetal/controller-utils/modutils"
	computev1beta1 "github.com/onmetal/onmetal-api/api/compute/v1beta1"
	corev1beta1 "github.com/onmetal/onmetal-api/api/core/v1beta1"
	ipamv1beta1 "github.com/onmetal/onmetal-api/api/ipam/v1beta1"
	networkingv1beta1 "github.com/onmetal/onmetal-api/api/networking/v1beta1"
	storagev1beta1 "github.com/onmetal/onmetal-api/api/storage/v1beta1"
	computeclient "github.com/onmetal/onmetal-api/internal/client/compute"
	ori "github.com/onmetal/onmetal-api/ori/apis/machine/v1alpha1"
	"github.com/onmetal/onmetal-api/ori/testing/machine"
	machinepoolletclient "github.com/onmetal/onmetal-api/poollet/machinepoollet/client"
	"github.com/onmetal/onmetal-api/poollet/machinepoollet/controllers"
	"github.com/onmetal/onmetal-api/poollet/machinepoollet/mcm"
	"github.com/onmetal/onmetal-api/poollet/orievent"
	utilsenvtest "github.com/onmetal/onmetal-api/utils/envtest"
	"github.com/onmetal/onmetal-api/utils/envtest/apiserver"
	"github.com/onmetal/onmetal-api/utils/envtest/controllermanager"
	"github.com/onmetal/onmetal-api/utils/envtest/process"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	. "sigs.k8s.io/controller-runtime/pkg/envtest/komega"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	metricserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var (
	cfg        *rest.Config
	testEnv    *envtest.Environment
	testEnvExt *utilsenvtest.EnvironmentExtensions
	k8sClient  client.Client
)

const (
	eventuallyTimeout    = 3 * time.Second
	pollingInterval      = 50 * time.Millisecond
	consistentlyDuration = 1 * time.Second
	apiServiceTimeout    = 5 * time.Minute

	controllerManagerService = "controller-manager"

	fooDownwardAPILabel = "custom-downward-api-label"
	fooAnnotation       = "foo"
)

func TestControllers(t *testing.T) {
	SetDefaultConsistentlyPollingInterval(pollingInterval)
	SetDefaultEventuallyPollingInterval(pollingInterval)
	SetDefaultEventuallyTimeout(eventuallyTimeout)
	SetDefaultConsistentlyDuration(consistentlyDuration)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	var err error
	By("bootstrapping test environment")
	testEnv = &envtest.Environment{}
	testEnvExt = &utilsenvtest.EnvironmentExtensions{
		APIServiceDirectoryPaths: []string{
			modutils.Dir("github.com/onmetal/onmetal-api", "config", "apiserver", "apiservice", "bases"),
		},
		ErrorIfAPIServicePathIsMissing: true,
		AdditionalServices: []utilsenvtest.AdditionalService{
			{
				Name: controllerManagerService,
			},
		},
	}

	cfg, err = utilsenvtest.StartWithExtensions(testEnv, testEnvExt)
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	DeferCleanup(utilsenvtest.StopWithExtensions, testEnv, testEnvExt)

	Expect(computev1beta1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(networkingv1beta1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(ipamv1beta1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(storagev1beta1.AddToScheme(scheme.Scheme)).To(Succeed())

	// Init package-level k8sClient
	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())
	SetClient(k8sClient)

	apiSrv, err := apiserver.New(cfg, apiserver.Options{
		MainPath:     "github.com/onmetal/onmetal-api/cmd/onmetal-apiserver",
		BuildOptions: []buildutils.BuildOption{buildutils.ModModeMod},
		ETCDServers:  []string{testEnv.ControlPlane.Etcd.URL.String()},
		Host:         testEnvExt.APIServiceInstallOptions.LocalServingHost,
		Port:         testEnvExt.APIServiceInstallOptions.LocalServingPort,
		CertDir:      testEnvExt.APIServiceInstallOptions.LocalServingCertDir,
	})
	Expect(err).NotTo(HaveOccurred())

	Expect(apiSrv.Start()).To(Succeed())
	DeferCleanup(apiSrv.Stop)

	Expect(utilsenvtest.WaitUntilAPIServicesReadyWithTimeout(apiServiceTimeout, testEnvExt, k8sClient, scheme.Scheme)).To(Succeed())

	ctrlMgr, err := controllermanager.New(cfg, controllermanager.Options{
		Args:         process.EmptyArgs().Set("controllers", "*"),
		MainPath:     "github.com/onmetal/onmetal-api/cmd/onmetal-controller-manager",
		BuildOptions: []buildutils.BuildOption{buildutils.ModModeMod},
		Host:         testEnvExt.GetAdditionalServiceHost(controllerManagerService),
		Port:         testEnvExt.GetAdditionalServicePort(controllerManagerService),
	})
	Expect(err).NotTo(HaveOccurred())

	Expect(ctrlMgr.Start()).To(Succeed())
	DeferCleanup(ctrlMgr.Stop)
})

func SetupTest() (*corev1.Namespace, *computev1beta1.MachinePool, *computev1beta1.MachineClass, *machine.FakeRuntimeService) {
	var (
		ns  = &corev1.Namespace{}
		mp  = &computev1beta1.MachinePool{}
		mc  = &computev1beta1.MachineClass{}
		srv = &machine.FakeRuntimeService{}
	)

	BeforeEach(func(ctx SpecContext) {
		*ns = corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-ns-",
			},
		}
		Expect(k8sClient.Create(ctx, ns)).To(Succeed(), "failed to create test namespace")
		DeferCleanup(k8sClient.Delete, ns)

		*mp = computev1beta1.MachinePool{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-mp-",
			},
		}
		Expect(k8sClient.Create(ctx, mp)).To(Succeed(), "failed to create test machine pool")
		DeferCleanup(k8sClient.Delete, mp)

		*mc = computev1beta1.MachineClass{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-mc-",
			},
			Capabilities: corev1beta1.ResourceList{
				corev1beta1.ResourceCPU:    resource.MustParse("1"),
				corev1beta1.ResourceMemory: resource.MustParse("1Gi"),
			},
		}
		Expect(k8sClient.Create(ctx, mc)).To(Succeed(), "failed to create test machine class")
		DeferCleanup(k8sClient.Delete, mc)

		*srv = *machine.NewFakeRuntimeService()
		srv.SetMachineClasses([]*machine.FakeMachineClassStatus{
			{
				MachineClassStatus: ori.MachineClassStatus{
					MachineClass: &ori.MachineClass{
						Name: mc.Name,
						Capabilities: &ori.MachineClassCapabilities{
							CpuMillis:   mc.Capabilities.CPU().MilliValue(),
							MemoryBytes: mc.Capabilities.Memory().Value(),
						},
					},
				},
			},
		})

		k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
			Scheme: scheme.Scheme,
			Metrics: metricserver.Options{
				BindAddress: "0",
			},
		})
		Expect(err).ToNot(HaveOccurred())

		indexer := k8sManager.GetFieldIndexer()
		Expect(machinepoolletclient.SetupMachineSpecNetworkInterfaceNamesField(ctx, indexer, mp.Name)).To(Succeed())
		Expect(machinepoolletclient.SetupMachineSpecVolumeNamesField(ctx, indexer, mp.Name)).To(Succeed())
		Expect(machinepoolletclient.SetupMachineSpecSecretNamesField(ctx, indexer, mp.Name)).To(Succeed())
		Expect(computeclient.SetupMachineSpecMachinePoolRefNameFieldIndexer(ctx, indexer)).To(Succeed())

		machineClassMapper := mcm.NewGeneric(srv, mcm.GenericOptions{
			RelistPeriod: 2 * time.Second,
		})
		Expect(k8sManager.Add(machineClassMapper)).To(Succeed())

		mgrCtx, cancel := context.WithCancel(context.Background())
		DeferCleanup(cancel)

		Expect((&controllers.MachineReconciler{
			EventRecorder:         &record.FakeRecorder{},
			Client:                k8sManager.GetClient(),
			MachineRuntime:        srv,
			MachineRuntimeName:    machine.FakeRuntimeName,
			MachineRuntimeVersion: machine.FakeVersion,
			MachineClassMapper:    machineClassMapper,
			MachinePoolName:       mp.Name,
			DownwardAPILabels: map[string]string{
				fooDownwardAPILabel: fmt.Sprintf("metadata.annotations['%s']", fooAnnotation),
			},
		}).SetupWithManager(k8sManager)).To(Succeed())

		machineEvents := orievent.NewGenerator(func(ctx context.Context) ([]*ori.Machine, error) {
			res, err := srv.ListMachines(ctx, &ori.ListMachinesRequest{})
			if err != nil {
				return nil, err
			}
			return res.Machines, nil
		}, orievent.GeneratorOptions{})

		Expect(k8sManager.Add(machineEvents)).To(Succeed())

		Expect((&controllers.MachineAnnotatorReconciler{
			Client:        k8sManager.GetClient(),
			MachineEvents: machineEvents,
		}).SetupWithManager(k8sManager)).To(Succeed())

		Expect((&controllers.MachinePoolReconciler{
			Client:             k8sManager.GetClient(),
			MachineRuntime:     srv,
			MachineClassMapper: machineClassMapper,
			MachinePoolName:    mp.Name,
		}).SetupWithManager(k8sManager)).To(Succeed())

		Expect((&controllers.MachinePoolAnnotatorReconciler{
			Client:             k8sManager.GetClient(),
			MachineClassMapper: machineClassMapper,
			MachinePoolName:    mp.Name,
		}).SetupWithManager(k8sManager)).To(Succeed())

		go func() {
			defer GinkgoRecover()
			Expect(k8sManager.Start(mgrCtx)).To(Succeed(), "failed to start manager")
		}()
	})

	return ns, mp, mc, srv
}
