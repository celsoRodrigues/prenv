package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cbrgm/githubevents/githubevents"
	"github.com/google/go-github/v47/github"
	"github.com/google/uuid"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	log := log.New(os.Stdout, "REPOWATCHDOG ", log.LstdFlags)

	if err := Run(log); err != nil {
		log.Fatal(err)
	}

}

func Run(logger *log.Logger) error {
	fmt.Println("starting prog...")
	GHSecret := os.Getenv("HOOK")

	//k8s clientset
	clientset, err := CreateK8sClientset()
	if err != nil {
		return err
	}

	// create a new event handler and pass the secret
	handle := githubevents.New(GHSecret)

	handle.OnPullRequestEventOpened(
		func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
			logger.Printf("%s request submitted!", *event.PullRequest.Title)
			//create the event
			if err := CreateEvent(clientset, "Github", "Pull request opened", "opened"); err != nil {
				return err
			}
			return nil
		},
	)

	handle.OnPullRequestEventClosed(
		func(deliveryID string, eventName string, event *github.PullRequestEvent) error {
			logger.Printf("%s request closed!", *event.PullRequest.Title)
			//create the event
			if err := CreateEvent(clientset, "Github", "Pull request closed", "closed"); err != nil {
				return err
			}
			return nil
		},
	)

	// add a http handleFunc
	http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {

		err := handle.HandleEventRequest(r)
		if err != nil {
			logger.Println("erroring out:", err)
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

	})

	// start the server listening on port 8000
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}

	return nil
}

func CreateEvent(clientset *kubernetes.Clientset, evType, evMessage, evReason string) error {

	rname := uuid.New()
	name := "eventCreator-" + rname.String()

	eventClient := clientset.CoreV1().Events("default")

	event := &apiv1.Event{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
		},
		Type:           evType,
		Message:        evMessage,
		Reason:         evReason,
		FirstTimestamp: metav1.NewTime(time.Time{}),
		LastTimestamp:  metav1.NewTime(metav1.Now().Time),
		InvolvedObject: apiv1.ObjectReference{
			Kind: "Pod",
		},
	}

	// Create Deployment

	_, err := eventClient.Create(context.TODO(), event, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil

}

func CreateK8sClientset() (*kubernetes.Clientset, error) {

	var kubeconfig *string
	var config *rest.Config
	if home := homedir.HomeDir(); home != "" {

		var err error
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			var err error
			config, err = rest.InClusterConfig()
			if err != nil {
				panic(err.Error())
			}
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &kubernetes.Clientset{}, err
	}

	return clientset, err

}
