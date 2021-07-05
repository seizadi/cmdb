export AWS_ACCOUNT ?= $(shell aws sts get-caller-identity --output text --query 'Account')
export AWS_REGION		     = us-west-2

K8S_VERSION := v1.21.0
K8S_WAIT := wait --timeout=120s
K8S_WAIT_ARGOCD := wait --timeout=600s
K8S_DEST_SERVER := https://kubernetes.default.svc
CMDB_REPO := https://github.com/seizadi/cmdb
CMDB_VALUES := https://raw.githubusercontent.com/seizadi/dc-repo/main/build/dev/seizadi-minikube/cmdb.yaml
ARGOCD_HOST := locahost

default: vm

vm: minikube argocd argo_aws_host argo_login setup_apps

eks: cluster argocd argo_minikube_host  argo_login setup_apps
test: argo_aws_host argo_login
.id:
	git config user.email | awk -F@ '{print $$1}' > .id

cluster.yaml: force .id cluster.yaml.in
	sed "s/{{ .Name }}/`cat .id`/g; s/{{ .Region }}/$(AWS_REGION)/g" cluster.yaml.in > $@

eks-deploy: cluster.yaml
	eksctl create cluster -f cluster.yaml

cluster: eks-deploy
	@echo "Get Cluster Config"
	aws eks update-kubeconfig --name $(shell cat .id)-cmdb
	@echo "Create NGINX Ingress Controller"
	# FIXME - The ingress on main has bug
	# kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/aws/deploy.yaml
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/a8408cdb51086a099a2c71ed3e68363eb3a7ae60/deploy/static/provider/aws/deploy.yaml
	kubectl $(K8S_WAIT) --namespace ingress-nginx --for=condition=ready pod --selector=app.kubernetes.io/component=controller
	@echo 'Done with build cluster'

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

argo_minikube_host: argocd_url=$(shell minikube service argocd-server -n argocd --url|tail -1)
argo_minikube_host:
	echo $(argocd_url) | awk -F'//' '{print $$2}' > .host

argo_aws_host:
	kubectl -n argocd get service argocd-server -o json | jq '.status.loadBalancer.ingress[0].hostname' > .host

#argo_login: argocd_password=$(shell kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2)
argo_login: argocd_password=$(shell kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
argo_login: argocd_address=.......elb.amazonaws.com
argo_login:
	@echo "Login to argocd using following command line:"
	@echo "argocd login $(shell cat .host) --username admin --password $(argocd_password) --insecure"
	@argocd login $(shell cat .host) --username admin --password $(argocd_password) --insecure

argo_password:
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
	@echo "\nUse above password for login"


setup_apps:
	argocd repo add $(CMDB_REPO) --name cmdb
	argocd app create cmdb --repo $(CMDB_REPO) --path repo/cmdb --dest-server $(K8S_DEST_SERVER) --dest-namespace cmdb --values-literal-file $(CMDB_VALUES)

clean:
	minikube delete
	 eksctl delete cluster --name $(shell cat .id)-cmdb --region $(AWS_REGION)

force:
