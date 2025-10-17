package handler

import (
	"embed"
	"fabricioveronez/kube-board/k8s"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates/*
var templatesFS embed.FS

var templates = []string{
	"templates/index.html",
	"templates/header.html",
	"templates/footer.html",
	"templates/menu.html",
	"templates/deployments.html",
	"templates/combobox.html",
}

type Linha struct {
	Coluna01 string
	Coluna02 string
	Coluna03 string
	Coluna04 string
}

type DtoPage struct {
	Namespaces   []string
	Titulo01     string
	Titulo02     string
	Titulo03     string
	Titulo04     string
	LinhasTabela []Linha
}

func Index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func Deployments(w http.ResponseWriter, r *http.Request) {

	k8sClient, err := k8s.NewClient()
	if err != nil {
		panic(err)
	}

	namespace := r.URL.Query().Get("namespace")
	if namespace == "" {
		namespace = "default"
	}

	listaNamespaces, err := k8sClient.ListNamespaces()
	if err != nil {
		panic(err)
	}

	dtoPage := DtoPage{Namespaces: []string{},
		LinhasTabela: []Linha{},
		Titulo01:     "NAME",
		Titulo02:     "READY",
		Titulo03:     "UP-TO-DATE",
		Titulo04:     "AVAILABLE",
	}

	for _, namespace := range listaNamespaces.Items {
		dtoPage.Namespaces = append(dtoPage.Namespaces, namespace.Name)
	}

	listaDeployments, err := k8sClient.ListDeployments(namespace)
	if err != nil {
		panic(err)
	}

	for _, deployment := range listaDeployments.Items {
		dtoPage.LinhasTabela = append(dtoPage.LinhasTabela, Linha{
			Coluna01: deployment.Name,
			Coluna02: fmt.Sprintf("%d / %d", deployment.Status.ReadyReplicas, deployment.Status.Replicas),
			Coluna03: fmt.Sprintf("%d", deployment.Status.UpdatedReplicas),
			Coluna04: fmt.Sprintf("%d", deployment.Status.AvailableReplicas),
		})
	}

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "deployments.html", dtoPage)
	if err != nil {
		panic(err)
	}
}

func ReplicaSets(w http.ResponseWriter, r *http.Request) {

	k8sClient, err := k8s.NewClient()
	if err != nil {
		panic(err)
	}

	namespace := r.URL.Query().Get("namespace")
	if namespace == "" {
		namespace = "default"
	}

	listaNamespaces, err := k8sClient.ListNamespaces()
	if err != nil {
		panic(err)
	}

	dtoPage := DtoPage{Namespaces: []string{},
		LinhasTabela: []Linha{},
		Titulo01:     "NAME",
		Titulo02:     "DESIRED",
		Titulo03:     "CURRENT",
		Titulo04:     "READY",
	}

	for _, namespace := range listaNamespaces.Items {
		dtoPage.Namespaces = append(dtoPage.Namespaces, namespace.Name)
	}

	listaReplicaSets, err := k8sClient.ListReplicaSets(namespace)
	if err != nil {
		panic(err)
	}

	for _, replicaSet := range listaReplicaSets.Items {
		dtoPage.LinhasTabela = append(dtoPage.LinhasTabela, Linha{
			Coluna01: replicaSet.Name,
			Coluna02: fmt.Sprintf("%d", *replicaSet.Spec.Replicas),
			Coluna03: fmt.Sprintf("%d", replicaSet.Status.Replicas),
			Coluna04: fmt.Sprintf("%d", replicaSet.Status.ReadyReplicas),
		})
	}

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "deployments.html", dtoPage)
	if err != nil {
		panic(err)
	}
}

func Pods(w http.ResponseWriter, r *http.Request) {

	k8sClient, err := k8s.NewClient()
	if err != nil {
		panic(err)
	}

	namespace := r.URL.Query().Get("namespace")
	if namespace == "" {
		namespace = "default"
	}

	listaNamespaces, err := k8sClient.ListNamespaces()
	if err != nil {
		panic(err)
	}

	dtoPage := DtoPage{Namespaces: []string{},
		LinhasTabela: []Linha{},
		Titulo01:     "NAME",
		Titulo02:     "QOS",
		Titulo03:     "STATUS",
		Titulo04:     "IP ADDRESS",
	}

	for _, namespace := range listaNamespaces.Items {
		dtoPage.Namespaces = append(dtoPage.Namespaces, namespace.Name)
	}

	listaPods, err := k8sClient.ListPods(namespace)
	if err != nil {
		panic(err)
	}

	for _, pod := range listaPods.Items {
		dtoPage.LinhasTabela = append(dtoPage.LinhasTabela, Linha{
			Coluna01: pod.Name,
			Coluna02: fmt.Sprintf("%v", pod.Status.QOSClass),
			Coluna03: fmt.Sprintf("%v", pod.Status.Phase),
			Coluna04: pod.Status.PodIP,
		})
	}

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "deployments.html", dtoPage)
	if err != nil {
		panic(err)
	}
}

func DaemonSets(w http.ResponseWriter, r *http.Request) {

	k8sClient, err := k8s.NewClient()
	if err != nil {
		panic(err)
	}

	namespace := r.URL.Query().Get("namespace")
	if namespace == "" {
		namespace = "default"
	}

	listaNamespaces, err := k8sClient.ListNamespaces()
	if err != nil {
		panic(err)
	}

	dtoPage := DtoPage{Namespaces: []string{},
		LinhasTabela: []Linha{},
		Titulo01:     "NAME",
		Titulo02:     "DESIRED",
		Titulo03:     "CURRENT",
		Titulo04:     "READY",
	}

	for _, namespace := range listaNamespaces.Items {
		dtoPage.Namespaces = append(dtoPage.Namespaces, namespace.Name)
	}

	listaPodsDaemonSet, err := k8sClient.ListDaemonSets(namespace)
	if err != nil {
		panic(err)
	}

	for _, daemonSet := range listaPodsDaemonSet.Items {
		dtoPage.LinhasTabela = append(dtoPage.LinhasTabela, Linha{
			Coluna01: daemonSet.Name,
			Coluna02: fmt.Sprintf("%v", daemonSet.Status.DesiredNumberScheduled),
			Coluna03: fmt.Sprintf("%v", daemonSet.Status.CurrentNumberScheduled),
			Coluna04: fmt.Sprintf("%v", daemonSet.Status.NumberReady),
		})
	}

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "deployments.html", dtoPage)
	if err != nil {
		panic(err)
	}
}

func StatefulSets(w http.ResponseWriter, r *http.Request) {

	k8sClient, err := k8s.NewClient()
	if err != nil {
		panic(err)
	}

	namespace := r.URL.Query().Get("namespace")
	if namespace == "" {
		namespace = "default"
	}

	listaNamespaces, err := k8sClient.ListNamespaces()
	if err != nil {
		panic(err)
	}

	dtoPage := DtoPage{Namespaces: []string{},
		LinhasTabela: []Linha{},
		Titulo01:     "NAME",
		Titulo02:     "READY",
		Titulo03:     "AVAILABLE",
		Titulo04:     "CURRENT",
	}

	for _, namespace := range listaNamespaces.Items {
		dtoPage.Namespaces = append(dtoPage.Namespaces, namespace.Name)
	}

	listaStatefulSets, err := k8sClient.ListStatefulSets(namespace)
	if err != nil {
		panic(err)
	}

	for _, statefulSet := range listaStatefulSets.Items {
		dtoPage.LinhasTabela = append(dtoPage.LinhasTabela, Linha{
			Coluna01: statefulSet.Name,
			Coluna02: fmt.Sprintf("%v", statefulSet.Status.ReadyReplicas),
			Coluna03: fmt.Sprintf("%v", statefulSet.Status.AvailableReplicas),
			Coluna04: fmt.Sprintf("%v", statefulSet.Status.CurrentReplicas),
		})
	}

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "deployments.html", dtoPage)
	if err != nil {
		panic(err)
	}
}

func Jobs(w http.ResponseWriter, r *http.Request) {

	k8sClient, err := k8s.NewClient()
	if err != nil {
		panic(err)
	}

	namespace := r.URL.Query().Get("namespace")
	if namespace == "" {
		namespace = "default"
	}

	listaNamespaces, err := k8sClient.ListNamespaces()
	if err != nil {
		panic(err)
	}

	dtoPage := DtoPage{Namespaces: []string{},
		LinhasTabela: []Linha{},
		Titulo01:     "NAME",
		Titulo02:     "COMPLETIONS",
		Titulo03:     "FAILEDS",
		Titulo04:     "ACTIVES",
	}

	for _, namespace := range listaNamespaces.Items {
		dtoPage.Namespaces = append(dtoPage.Namespaces, namespace.Name)
	}

	listaJobs, err := k8sClient.ListJobs(namespace)
	if err != nil {
		panic(err)
	}

	for _, job := range listaJobs.Items {
		dtoPage.LinhasTabela = append(dtoPage.LinhasTabela, Linha{
			Coluna01: job.Name,
			Coluna02: fmt.Sprintf("%v", job.Status.CompletionTime),
			Coluna03: fmt.Sprintf("%v", job.Status.Failed),
			Coluna04: fmt.Sprintf("%v", job.Status.Active),
		})
	}

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "deployments.html", dtoPage)
	if err != nil {
		panic(err)
	}
}

func CronJobs(w http.ResponseWriter, r *http.Request) {

	k8sClient, err := k8s.NewClient()
	if err != nil {
		panic(err)
	}

	namespace := r.URL.Query().Get("namespace")
	if namespace == "" {
		namespace = "default"
	}

	listaNamespaces, err := k8sClient.ListNamespaces()
	if err != nil {
		panic(err)
	}

	dtoPage := DtoPage{Namespaces: []string{},
		LinhasTabela: []Linha{},
		Titulo01:     "NAME",
		Titulo02:     "SCHEDULE",
		Titulo03:     "SUSPEND",
		Titulo04:     "ACTIVE",
	}

	for _, namespace := range listaNamespaces.Items {
		dtoPage.Namespaces = append(dtoPage.Namespaces, namespace.Name)
	}

	listaCronJobs, err := k8sClient.ListCronJobs(namespace)
	if err != nil {
		panic(err)
	}

	for _, cronJob := range listaCronJobs.Items {
		dtoPage.LinhasTabela = append(dtoPage.LinhasTabela, Linha{
			Coluna01: cronJob.Name,
			Coluna02: fmt.Sprintf("%v", cronJob.Spec.Schedule),
			Coluna03: fmt.Sprintf("%v", cronJob.Spec.Suspend),
			Coluna04: fmt.Sprintf("%v", len(cronJob.Status.Active)),
		})
	}

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(w, "deployments.html", dtoPage)
	if err != nil {
		panic(err)
	}
}
