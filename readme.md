## Build and Run
``` sh
cd ./producer
go build -o producer
./producer -cfgpath=./producer.toml

cd ./consumer
go build -o consumer
./consumer -cfgpath=./consumer.toml
```

## Deploy to Dockerhub
When you push to master branch the github action will automatically build image and push it to my dockerhub

## Deploy to k8s
``` sh
kubectl create namespace try
cd helm
helm install --namespace=try -f values.yaml simple-nginx .
helm delete simple-nginx --namespace=try
```

---

## Install Nginx Ingress in K8S
* https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.4/deploy/static/provider/baremetal/deploy.yaml
``` sh
kubectl apply -f ./nginx_ingress/deploy.yaml
# kubectl delete -f ./nginx_ingress/deploy.yaml
```