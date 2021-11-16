## Install ingress-nginx with prometheus exporter by single yaml
* https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.4/deploy/static/provider/baremetal/deploy.yaml
``` sh
kubectl apply -f ./deploy.yaml
```
## Install ingress-nginx by helm
``` sh
kubectl create namespace ingress-nginx
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install --set controller.service.type=NodePort --set controller.metrics.enabled=true -n ingress-nginx ingress-nginx ingress-nginx/ingress-nginx
```