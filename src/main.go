package main

import (
	"embed"
	"fabricioveronez/kube-board/handler"
	"io/fs"
	"net/http"
)

//go:embed static/*
var staticFS embed.FS

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/deployments", handler.Deployments)
	mux.HandleFunc("/replicasets", handler.ReplicaSets)
	mux.HandleFunc("/pods", handler.Pods)
	mux.HandleFunc("/daemonsets", handler.DaemonSets)
	mux.HandleFunc("/statefulsets", handler.StatefulSets)
	mux.HandleFunc("/jobs", handler.Jobs)
	mux.HandleFunc("/cronjobs", handler.CronJobs)

	// Configuração para servir arquivos estáticos
	staticFiles, _ := fs.Sub(staticFS, "static")

	// Adiciona o prefixo '/static/' ao caminho
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

	http.ListenAndServe(":3000", mux)
}
