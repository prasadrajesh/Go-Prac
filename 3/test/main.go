// prototype program to run outside k8s cluster. once functional, move logic into operator.
// authenticate with kubeconfig and query k8s cluster for info on:
// istio VirtualServices, istio Gateways, k8s Services (istio ingress gateways).
// map each VirtualService to gateway k8s service, using the selector
// to run outside k8s cluster:
/*
export NAMESPACE='tts'
export ISTIO_NAMESPACE='istio-system'
export KUBECONFIG='/Users/james_gawley/work/onecloud/sandbox/istio_to_consul/kubeconfig'
*/

// istio client-go package might need a patch with this for AAD integration later:
// import "k8s.io/client-go/plugin/pkg/client/auth/azure"

package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/consul/api"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// structure of Consul service - not all fields to register with Consul
// Gateway* fields will be needed to find Address and Port
type Service struct {
	Id              string
	Name            string
	GatewayNs       string
	GatewayName     string
	GatewaySelector map[string]string
	Address         string
	Port            int
	Tags            []string
	Check           []struct {
		Name     string
		Args     string
		Interval string
		Timeout  string
	}
}

func getGwSelectors(kubeconfig string, namespace string, istionamespace string, service *Service) {
	// standard k8s rest client, to use in both istio and k8s clientsets
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to create k8s rest client: %s", err)
	}
	// istio-specific "clientset"
	ic, err := versionedclient.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}
	gwList, err := ic.NetworkingV1alpha3().Gateways(istionamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get Gateway in %s namespace: %s", namespace, err)
	}
	for g := range gwList.Items {
		gw := gwList.Items[g]
		if gw.ObjectMeta.Name == service.GatewayName {
			service.GatewaySelector = gw.Spec.GetSelector()
		}
	}
	//log.Printf("inside func getGwSelectors. service struct Id: %s", service.Id)
	//log.Printf("inside func getGwSelectors. service struct Name: %s", service.Name)
	//log.Printf("inside func getGwSelectors. service struct GatewayNs: %s", service.GatewayNs)
	//log.Printf("inside func getGwSelectors. service struct GatewayName: %s", service.GatewayName)
	//log.Printf("inside func getGwSelectors. service struct GatewaySelector: %s", service.GatewaySelector)
	//log.Printf("inside func getGwSelectors. service struct Address: %s", service.Address)
	//log.Printf("inside func getGwSelectors. service struct Port: %d\n\n", service.Port)
	return
}

func getGwAddrAndPort(kubeconfig string, istionamespace string, service *Service) {
	// standard k8s rest client, to use in both istio and k8s clientsets
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to create k8s rest client: %s", err)
	}
	// standard k8s clientset, not to be confused with istio clientset
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("Failed to create k8s client: %s", err)
	}
	// k8s services - will be needed to map the gateway to svc IP via the selector
	svcs, err := clientset.CoreV1().Services(istionamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get k8s services: %s", err)
	}
	//
	// need to get the istio label from svc to match Selector:
	//  labels:
	//    app: internal-istio-ingressgateway
	//    fluxcd.io/sync-gc-mark: sha256.8kPkyU3PO1km-dWxWAumByNTH_OlX7Hix5LyjWLWwvs
	//    istio: internal-ingressgateway
	//    release: istio
	//  name: internal-istio-ingressgateway
	for i := range svcs.Items {
		svc := svcs.Items[i]
		if svc.ObjectMeta.Labels["istio"] == service.GatewaySelector["istio"] {
			if len(svc.Status.LoadBalancer.Ingress) == 1 {
				service.Address = svc.Status.LoadBalancer.Ingress[0].IP
				log.Printf("inside func getGwAddrAndPort. service struct Id: %s", service.Id)
				log.Printf("inside func getGwAddrAndPort. service struct Name: %s", service.Name)
				log.Printf("inside func getGwAddrAndPort. service struct GatewayNs: %s", service.GatewayNs)
				log.Printf("inside func getGwAddrAndPort. service struct GatewayName: %s", service.GatewayName)
				log.Printf("inside func getGwAddrAndPort. service struct GatewaySelector: %s", service.GatewaySelector)
				log.Printf("inside func getGwAddrAndPort. service struct Address: %s", service.Address)
				log.Printf("inside func getGwAddrAndPort. service struct Port: %d\n\n", service.Port)
				// handle getting port info first, but eventually registerServiceToConsul(service)
				registerServiceToConsul(service)
			}
		}
	}
	return
}

func registerServiceToConsul(service *Service) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to get new consul API client: %s", err)
	}
	reg := new(api.AgentServiceRegistration)
	reg.ID = service.Id
	reg.Name = service.Name
	reg.Port = service.Port
	reg.Address = service.Address
	// for now, tags will be the name of the istio gw in case we need to trace,
	// but tags can be used for more info, like namespace or maybe cluster ID
	//reg.Tags = //[]string
	// we'll talk about adding check laters
	err1 := client.Agent().ServiceRegister(reg)
	if err1 != nil {
		log.Fatalf("Failed to register service to consul: %s", err1)
	}
	return
}

func main() {
	kubeconfig := os.Getenv("KUBECONFIG")
	namespace := os.Getenv("NAMESPACE")
	istionamespace := os.Getenv("ISTIO_NAMESPACE")
	if len(kubeconfig) == 0 || len(namespace) == 0 {
		log.Fatalf("Environment variables KUBECONFIG and NAMESPACE need to be set")
	}
	// standard k8s rest client, to use in both istio and k8s clientsets
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to create k8s rest client: %s", err)
	}
	// istio-specific "clientset"
	ic, err := versionedclient.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}
	// istio VritualServices
	vsList, err := ic.NetworkingV1alpha3().VirtualServices(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get VirtualService in %s namespace: %s", namespace, err)
	}
	for i := range vsList.Items {
		vs := vsList.Items[i]
		if vs.Spec.GetGateways() != nil {
			vsName := vs.ObjectMeta.Name
			vsGw := vs.Spec.GetGateways()
			if len(vsGw) != 1 {
				log.Fatalf("VirtualService does not have 1 Gateway. Number of Gateways: %d", len(vsGw))
			}
			// vsGw[0] is composed of 'namespace/name', like this: [istio-system/gateway-onecloud-hosting-cerence-net]
			// separate namespace and name
			gwNs := strings.Split(vsGw[0], "/")[0]
			gwName := strings.Split(vsGw[0], "/")[1]
			vsHosts := vs.Spec.GetHosts()
			// on OC integration cluster I see multiple VSs with the SAME NAME (see tts namespace).
			// let's try making the HOST the unique ID for consul service
			service := Service{Id: vsHosts[0], Name: vsName, GatewayNs: gwNs, GatewayName: gwName}
			log.Printf("\n\nvsList loop. Service struct 'service.Id': %s\n", service.Id)
			getGwSelectors(kubeconfig, namespace, istionamespace, &service)
			getGwAddrAndPort(kubeconfig, istionamespace, &service)
		} else {
			log.Printf("Index: %d Found VS WITHOUT gateway. I don't think we need this.\n VS name: %s\n\n", i, vs.ObjectMeta.Name)
		}
	}
	return
}
