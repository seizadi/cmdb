
K8S_VERSION := v1.21.0
K8S_WAIT := wait --timeout=120s
K8S_WAIT_ARGOCD := wait --timeout=600s
K8S_DEST_SERVER := https://kubernetes.default.svc
CMDB_REPO := https://github.com/seizadi/cmdb
CMDB_VALUES := https://raw.githubusercontent.com/seizadi/dc-repo/main/build/dev/seizadi-minikube/cmdb.yaml

default: vm

vm: minikube argocd argo_login setup_apps

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

#argo_login: argocd_password=$(shell kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2)
argo_login: argocd_password=$(shell kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
argo_login: argocd_url=$(shell minikube service argocd-server -n argocd --url|tail -1)
argo_login: argocd_address=$(shell echo $(argocd_url) | awk -F'//' '{print $$2}')
argo_login:
	@echo "Login to argocd using following command line:"
	@echo "argocd login $(argocd_address) --username admin --password $(argocd_password) --insecure"
	@argocd login $(argocd_address) --username admin --password $(argocd_password) --insecure

setup_apps:
	argocd repo add $(CMDB_REPO) --name cmdb
	argocd app create cmdb --repo $(CMDB_REPO) --path repo/cmdb --dest-server $(K8S_DEST_SERVER) --dest-namespace cmdb --values-literal-file $(CMDB_VALUES)

clean:
	minikube delete
