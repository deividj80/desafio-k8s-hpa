# Desafio - K8S - HPA
Projeto com configurações do kubernetes.
### Tecnologias
- GoLang
- Docker
- Docker-compose
- MiniKube
- Google Cloud Platform
### Instrução para teste
- Coamndo para criar pod de teste
```sh
$ kubectl run -it hpa-loader --image=busybox /bin/sh
$ while true; do wget -q -O- http://go-hpa.default.svc.cluster.local:8080/; done;
```
### Container Registry
- link: https://hub.docker.com/repository/docker/deividj80/go-hpa