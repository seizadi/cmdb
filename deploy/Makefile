export AWS_ACCOUNT ?= $(shell aws sts get-caller-identity --output text --query 'Account')
export AWS_REGION		     = us-west-2

K8S_VERSION := v1.21.2
K8S_WAIT := wait --timeout=120s
K8S_WAIT_ARGOCD := wait --timeout=600s
K8S_DEST_SERVER := https://kubernetes.default.svc
CMDB_REPO := https://github.com/seizadi/cmdb
DC_REPO := https://github.com/seizadi/dc-repo
CMDB_MINIKUBE_VALUES := https://raw.githubusercontent.com/seizadi/dc-repo/main/build/dev/minikube/cmdb.yaml
CMDB_EKS_VALUES := https://raw.githubusercontent.com/seizadi/dc-repo/main/build/prod/seizadi-eks/cmdb.yaml
ARGOCD_HOST := locahost

GIT_USER := $(shell cat ~/credentials.txt | grep ^username | awk '{print $$2}')
GIT_PASSWORD := $(shell cat ~/credentials.txt | grep ^password | awk '{print $$2}')
DOCKERHUB_USERNAME := $(shell cat ~/credentials.txt | grep ^dockerhub_username | awk '{print $$2}')
DOCKERHUB_PASSWORD := $(shell cat ~/credentials.txt | grep ^dockerhub_password | awk '{print $$2}')


envs := $(shell cat envs.yaml | yq eval -j| jq .envs | jq 'keys[]' -r)
apps := $(shell cat apps.yaml | yq eval -j| jq .apps | jq 'keys[]' -r)
setup_minikube_apps_app_values := $(foreach E,$(envs),$(addsuffix .yaml,$(addprefix setup_minikube_apps_build^$E^,$(apps))))
sync_minikube_apps_app_values := $(foreach E,$(envs),$(addsuffix .yaml,$(addprefix sync_minikube_apps_build^$E^,$(apps))))

default: vm

setup_minikube_apps: $(setup_minikube_apps_app_values)
sync_minikube_apps: $(sync_minikube_apps_app_values)

vm: minikube docker_login argocd argo_minikube_host argo_login setup_minikube_apps sync_minikube_apps

eks: cluster argocd crossplane crossplane_trust crossplane_aws argo_aws_host argo_login setup_eks_apps

.id:
	git config user.email | awk -F@ '{print $$1}' > .id

.cluster: .id
	echo "$(shell cat .id)-cmdb" > .cluster

cluster.yaml: force .cluster cluster.yaml.in
	sed "s/{{ .Name }}/`cat .id`/g; s/{{ .ClusterName }}/`cat .cluster`/g; s/{{ .Region }}/$(AWS_REGION)/g" cluster.yaml.in > $@

eks-deploy: cluster.yaml
	eksctl create cluster -f cluster.yaml

eks-oidc: cluster.yaml
eks-oidc: cluster_name=$(shell cat .cluster )
eks-oidc:
	eksctl utils associate-iam-oidc-provider --cluster $(cluster_name) --region $(AWS_REGION) --approve

eks-config: cluster.yaml
eks-config: cluster_name=$(shell cat .cluster )
eks-config:
	@echo "Get Cluster Config"
	aws eks update-kubeconfig --name $(cluster_name)

cluster: eks-deploy eks-oidc eks-config
	@echo "Create NGINX Ingress Controller"
	# FIXME - The ingress on main has bug
	# kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/aws/deploy.yaml
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/a8408cdb51086a099a2c71ed3e68363eb3a7ae60/deploy/static/provider/aws/deploy.yaml
	kubectl $(K8S_WAIT) --namespace ingress-nginx --for=condition=ready pod --selector=app.kubernetes.io/component=controller
	@echo 'Done with cluster build'

minikube:
	minikube start --cpus=2 --memory=4g --kubernetes-version $(K8S_VERSION) --driver=virtualbox --addons ingress
	@echo "Waiting on minikube to come up, this can take a long time!"
	kubectl $(K8S_WAIT) -n kube-system --for=condition=Available deployment.apps/coredns
	kubectl create secret generic regcred --from-file=.dockerconfigjson=$$HOME/docker_config.json --type=kubernetes.io/dockerconfigjson

docker_login:
	minikube ssh -- docker login -u ${DOCKERHUB_USERNAME} -p ${DOCKERHUB_PASSWORD}

argocd:
	kubectl create namespace argocd
	kubectl apply -n argocd -f ./argocd-install-e17565ae8124706feb7c17b23195885d8173a26b.yaml
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
	# Found that LoadBalancer was not setup even when the ingress available
	# FATA[0000] dial tcp: lookup ....elb.amazonaws.com: no such host
	# FIXME - Need aws call to make sure LoadBalancer is setup, for now I put argocd login later
	kubectl -n argocd get service argocd-server -o json | jq '.status.loadBalancer.ingress[0].hostname' > .host

#argo_login: argocd_password=$(shell kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2)
argo_login: argocd_password=$(shell kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -D)
argo_login: argocd_address=.......elb.amazonaws.com
argo_login:
	@echo "Login to argocd using following command line:"
	@echo "argocd login $(shell cat .host) --username admin --password $(argocd_password) --insecure"
	@argocd login $(shell cat .host) --username admin --password $(argocd_password) --insecure

argo_password:
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
	@echo "\nUse above password for login"

crossplane-config.yaml: provider_role=$(shell cat .cluster )-provider-role
crossplane-config.yaml: force crossplane-config.yaml.in
	sed "s/{{ .Account }}/$(AWS_ACCOUNT)/g; s/{{ .IamRoleName }}/$(provider_role)/g" crossplane-config.yaml.in > $@

crossplane: crossplane-config.yaml
	@echo "Install Cross-Plane"
	kubectl create namespace crossplane-system
	helm repo add crossplane-stable https://charts.crossplane.io/stable
	helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
	kubectl apply -f crossplane-config.yaml
	kubectl $(K8S_WAIT) -n crossplane-system --for=condition=Available deployment.apps/crossplane
	kubectl $(K8S_WAIT) -n crossplane-system --for=condition=Available deployment.apps/crossplane-rbac-manager
	# Need CRDs to install
	kubectl $(K8S_WAIT) --for=condition=Established crd/providers.pkg.crossplane.io
	@echo "Done with Cross-Plane Installation"

provider-trust.json: OIDC_PROVIDER=$(shell aws eks describe-cluster --name $(shell cat .cluster) --region $(AWS_REGION) --query "cluster.identity.oidc.issuer" --output text | sed -e "s/^https:\/\///")
provider-trust.json: OIDC_PROVIDER_ESCAPED=$(shell echo $(OIDC_PROVIDER) | sed 's/\//\\\//g' )
provider-trust.json: force provider-trust.json.in
	sed "s/{{ .Account }}/$(AWS_ACCOUNT)/g; s/{{ .OidcProvider }}/$(OIDC_PROVIDER_ESCAPED)/g; s/{{ .NameSpace }}/crossplane-system/g" provider-trust.json.in > $@


crossplane_trust: provider_role=$(shell cat .cluster )-provider-role
crossplane_trust: provider-trust.json
	@echo "Setup CrossPlane Trust"
	aws iam create-role --role-name $(provider_role) --assume-role-policy-document file://provider-trust.json --description "IAM role for $(shell cat .cluster) Cluster for CrossPlane pprovider-aws" &> /dev/null
	# FIXME - Setup CrossPlane with minimum access right now it has full admin access
	aws iam attach-role-policy --role-name $(provider_role) --policy-arn=arn:aws:iam::aws:policy/AdministratorAccess &> /dev/null
	kubectl apply -f provider-config.yaml
	@echo "CrossPlane Trust Established"

vpcid: vpcid=$(shell aws eks describe-cluster --region $(AWS_REGION) --name $(shell cat .cluster) --query "cluster.resourcesVpcConfig.vpcId" --output text)
vpcid:
	echo $(vpcid)

rds-subnet.yaml: vpcid=$(shell aws eks describe-cluster --region $(AWS_REGION) --name $(shell cat .cluster) --query "cluster.resourcesVpcConfig.vpcId" --output text)
rds-subnet.yaml: pubSubIds=$(shell aws ec2 describe-subnets --region $(AWS_REGION) --filters '[{"Name":"vpc-id","Values":["$(vpcid)"]}]' | jq '.Subnets | .[] | select(.MapPublicIpOnLaunch == true) | .SubnetId')
rds-subnet.yaml: pubSub1=$(shell echo $(pubSubIds) | awk '{print $$1}')
rds-subnet.yaml: pubSub2=$(shell echo $(pubSubIds) | awk '{print $$2}')
rds-subnet.yaml: pubSub3=$(shell echo $(pubSubIds) | awk '{print $$3}')
rds-subnet.yaml: force rds-subnet.yaml.in
	sed "s/{{ .Name }}/$(shell cat .cluster)-rds-subnetgroup/g; s/{{ .Region }}/$(AWS_REGION)/g; s/{{ .Sub1 }}/$(pubSub1)/g; s/{{ .Sub2 }}/$(pubSub2)/g; s/{{ .Sub3 }}/$(pubSub2)/g; s/{{ .Description }}/$(shell cat .cluster) EKS map vpc to rds/g" rds-subnet.yaml.in > $@

rds-sg.yaml: vpcid=$(shell aws eks describe-cluster --region $(AWS_REGION) --name $(shell cat .cluster) --query "cluster.resourcesVpcConfig.vpcId" --output text)
rds-sg.yaml: force rds-sg.yaml.in
	sed "s/{{ .Name }}/$(shell cat .cluster)-rds-sg/g; s/{{ .Region }}/$(AWS_REGION)/g; s/{{ .VpcId }}/$(vpcid)/g; s/{{ .Port }}/5432/g; s/{{ .Description }}/$(shell cat .cluster) EKS open rds for workloads/g" rds-sg.yaml.in > $@

rds-composition.yaml: vpcid=$(shell aws eks describe-cluster --region $(AWS_REGION) --name $(shell cat .cluster) --query "cluster.resourcesVpcConfig.vpcId" --output text)
rds-composition.yaml: force rds-composition.yaml.in
	sed "s/{{ .Region }}/$(AWS_REGION)/g; s/{{ .DbInstanceClass }}/db.t2.small/g; s/{{ .DbEngineVersion }}/11.7/g; s/{{ .VpcSgRef }}/$(shell cat .cluster)-rds-sg/g; s/{{ .SubnetRef }}/$(shell cat .cluster)-rds-subnetgroup/g" rds-composition.yaml.in > $@

postgres-direct.yaml: force postgres-direct.yaml.in
	sed "s/{{ .Region }}/$(AWS_REGION)/g; s/{{ .DbInstanceClass }}/db.t2.small/g; s/{{ .DbEngineVersion }}/11.7/g; s/{{ .VpcSgRef }}/$(shell cat .cluster)-rds-sg/g; s/{{ .SubnetRef }}/$(shell cat .cluster)-rds-subnetgroup/g" postgres-direct.yaml.in > $@

rds-subnetgroup: rds-subnet.yaml
	kubectl apply -f rds-subnet.yaml
	kubectl $(K8S_WAIT) --for=condition=Ready dbsubnetgroup/$(shell cat .cluster)-rds-subnetgroup
	kubectl $(K8S_WAIT) --for=condition=Synced dbsubnetgroup/$(shell cat .cluster)-rds-subnetgroup

rds-sg: rds-sg.yaml
	kubectl apply -f rds-sg.yaml
	kubectl $(K8S_WAIT) --for=condition=Ready securitygroup/$(shell cat .cluster)-rds-sg
	kubectl $(K8S_WAIT) --for=condition=Synced securitygroup/$(shell cat .cluster)-rds-sg

rds-composition: rds-composition.yaml
	kubectl apply -f rds-composition.yaml

rds-comp-resource-def: rds-comp-resource-def.yaml
	kubectl apply -f rds-comp-resource-def.yaml

crossplane_aws: rds-subnetgroup rds-sg rds-composition rds-comp-resource-def
	@echo "Setup CrossPlane AWS Composition Resources"

postgres-direct: postgres-direct.yaml
	kubectl apply -f postgres-direct.yaml

postgres-claim: postgres-claim.yaml
	kubectl apply -f postgres-claim.yaml

$(setup_minikube_apps_app_values): app=$(shell echo $@ | awk -F'^' '{print $$3}' | awk -F'.' '{print $$1}')
$(setup_minikube_apps_app_values): url=$(shell cat apps.yaml | yq eval -j | jq .apps|jq .${app}.repo -r)
$(setup_minikube_apps_app_values): path=$(shell cat apps.yaml | yq eval -j | jq .apps|jq .${app}.path -r)
$(setup_minikube_apps_app_values): namespace=$(shell cat apps.yaml | yq eval -j | jq .apps|jq .${app}.namespace -r)
$(setup_minikube_apps_app_values): manifest=$(shell cat apps.yaml | yq eval -j | jq .apps|jq .${app}.manifest -r)
$(setup_minikube_apps_app_values): force
	argocd repo add $(url) --name ${app} --username "$(GIT_USER)" --password "$(GIT_PASSWORD)"
	argocd app create ${app} --repo $(url) --path ${path} --dest-server $(K8S_DEST_SERVER) --dest-namespace ${namespace} --values-literal-file ${manifest}


$(sync_minikube_apps_app_values): app=$(shell echo $@ | awk -F'^' '{print $$3}' | awk -F'.' '{print $$1}')
$(sync_minikube_apps_app_values): force
	argocd app sync ${app}

setup_eks_apps:
	argocd repo add $(CMDB_REPO) --name cmdb
	argocd repo add $(DC_REPO) --name dc-repo
	argocd app create cmdb --repo $(CMDB_REPO) --path repo/cmdb --dest-server $(K8S_DEST_SERVER) --dest-namespace cmdb --values-literal-file $(CMDB_EKS_VALUES)
	argocd repo add "https://github.com/apptriton/secops" --name secops
	argocd app create secops --repo "https://github.com/apptriton/secops" --path repo/secops --dest-server $(K8S_DEST_SERVER) --dest-namespace default --values-literal-file "minikube.yaml"
	argocd repo add "https://github.com/apptriton/secops-scanner" --name secops-scanner
	argocd app create secops-scanner --repo "https://github.com/apptriton/secops-scanner" --path repo/secops --dest-server $(K8S_DEST_SERVER) --dest-namespace default
	argocd repo add "https://github.com/apptriton/helm-discovery" --name helm-discovery
	argocd app create helm-discovery --repo "https://github.com/apptriton/helm-discovery" --path repo/helm-discovery --dest-server $(K8S_DEST_SERVER) --dest-namespace default

clean_vm:
	minikube delete

clean_eks: provider_role=$(shell cat .cluster )-provider-role
clean_eks:
	# FIXME - Need to cleanup IAM Roles
	# aws iam detach-role-policy --role-name $(provider_role) --policy-arn
	eksctl delete cluster --name $(shell cat .cluster) --region $(AWS_REGION)

force:
