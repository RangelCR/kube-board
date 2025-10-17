# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Arquitetura do Projeto

O Kube-Board é uma aplicação web em Go que fornece uma interface para visualizar recursos do Kubernetes. A aplicação é estruturada da seguinte forma:

- **Frontend**: Templates HTML com CSS/imagens estáticos servidos via `embed.FS`
- **Backend**: Servidor HTTP em Go usando `net/http` padrão
- **Cliente Kubernetes**: Utiliza `client-go` para comunicação com clusters K8s (suporta in-cluster e kubeconfig local)

### Estrutura de Diretórios

```
src/
├── main.go              # Ponto de entrada, configuração do servidor HTTP
├── handler/             # Handlers HTTP e templates
│   ├── index.go         # Todos os handlers de rotas
│   └── templates/       # Templates HTML embutidos
├── k8s/                 # Cliente Kubernetes
│   └── client.go        # Wrapper para client-go
├── static/              # Arquivos CSS/imagens (embutidos)
│   ├── css/
│   └── img/
├── go.mod               # Dependências Go
└── Dockerfile           # Build multi-stage
```

### Rotas Disponíveis

- `/` - Página inicial
- `/deployments` - Lista deployments
- `/replicasets` - Lista replica sets  
- `/pods` - Lista pods
- `/daemonsets` - Lista daemon sets
- `/statefulsets` - Lista stateful sets
- `/jobs` - Lista jobs
- `/cronjobs` - Lista cron jobs
- `/static/` - Arquivos estáticos

## Comandos de Desenvolvimento

### Build Local
```bash
cd src
go build -o kube-board
```

### Executar Aplicação
```bash
cd src
go run main.go
# Aplicação roda na porta 3000
```

### Build Docker
```bash
cd src
docker build -t kube-board .
```

### Testes
```bash
cd src
go test ./...
```

### Módulos Go
```bash
cd src
go mod tidy       # Limpa dependências
go mod download   # Baixa dependências
```

## Deploy no Kubernetes

Os manifests estão em `/k8s/`:
- `deployment.yaml` - Deployment e Service
- `rbac.yaml` - Service Account e permissões

```bash
kubectl apply -f k8s/
```

## Configuração K8s

A aplicação detecta automaticamente:
1. **In-cluster**: Usa service account quando rodando dentro do cluster
2. **Local**: Usa `~/.kube/config` para desenvolvimento local

## CI/CD

Pipeline GitHub Actions (`.github/workflows/main.yml`):
- Build automático da imagem Docker
- Push para Docker Hub como `fabricioveronez/kube-board:latest` e `fabricioveronez/kube-board:v1`

## Tecnologias

- **Go 1.22**: Linguagem principal
- **client-go v0.29.3**: Cliente oficial Kubernetes
- **Templates Go**: Renderização HTML
- **Docker multi-stage**: Build otimizado
- **Alpine Linux**: Imagem base de produção