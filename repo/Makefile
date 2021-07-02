
K8S_VERSION := v1.21.0
K8S_WAIT := wait --timeout=120s
K8S_WAIT_ARGOCD := wait --timeout=600s
K8S_DEST_SERVER := https://kubernetes.default.svc
CMDB_REPO := https://github.com/seizadi/cmdb
CMDB_VALUES := https://raw.githubusercontent.com/seizadi/dc-repo/main/build/dev/seizadi-minikube/cmdb.yaml

default: local

local: kind argocd kind_argo_login setup_apps

vm: minikube argocd minikube_argo_login setup_apps

kind:
	kind create cluster --config kind.yaml
	@ echo Add NGINX Ingress
	# FIXME - Bug in deployment can not use master, revert to master when it is fixed.
	# kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/a8408cdb51086a099a2c71ed3e68363eb3a7ae60/deploy/static/provider/kind/deploy.yaml
	kubectl $(K8S_WAIT)  --namespace ingress-nginx --for=condition=ready pod --selector=app.kubernetes.io/component=controller

minikube:
	minikube start --cpus=2 --memory=4g --kubernetes-version $(K8S_VERSION) --driver=virtualbox --addons ingress
	@echo "Waiting on minikube to come up, this can take a long time!"
	kubectl $(K8S_WAIT) -n kube-system --for=condition=Available deployment.apps/coredns

argocd:
	kubectl create namespace argocd
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
	kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
	@echo "Waiting on ArgoCd to come up, this can take a long time!"
	kubectl $(K8S_WAIT_ARGOCD) -n argocd --for=condition=Available deployment.apps/argocd-dex-server
	kubectl $(K8S_WAIT_ARGOCD) -n argocd --for=condition=Available deployment.apps/argocd-redis
	kubectl $(K8S_WAIT_ARGOCD) -n argocd --for=condition=Available deployment.apps/argocd-repo-server
	kubectl $(K8S_WAIT_ARGOCD) -n argocd --for=condition=Available deployment.apps/argocd-server

# FIXME - Merge this with other argo_logins
#argo_login: argocd_password=$(shell kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2)
minikube_argo_login: argocd_password=$(shell kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
minikube_argo_login: argocd_url=$(shell minikube service argocd-server -n argocd --url|tail -1)
minikube_argo_login: argocd_address=$(shell echo $(argocd_url) | awk -F'//' '{print $$2}')
minikube_argo_login:
	@echo "Login to argocd using following command line:"
	@echo "argocd login $(argocd_address) --username admin --password $(argocd_password) --insecure"
	@argocd login $(argocd_address) --username admin --password $(argocd_password) --insecure

#argo_login: argocd_password=$(shell kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2)
kind_argo_login: argocd_password=$(shell kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
kind_argo_login: argocd_address=localhost:30828
kind_argo_login:
	@echo "Login to argocd using following command line:"
	@echo "argocd login $(argocd_address) --username admin --password $(argocd_password) --insecure"
	@argocd login $(argocd_address) --username admin --password $(argocd_password) --insecure

setup_apps:
	argocd repo add $(CMDB_REPO) --name cmdb
	argocd app create cmdb --repo $(CMDB_REPO) --path repo/cmdb --dest-server $(K8S_DEST_SERVER) --dest-namespace cmdb --values-literal-file $(CMDB_VALUES)

clean:
	minikube delete
