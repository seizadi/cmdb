#!/bin/bash
update_registry_secret() {
  login_cmd=$(aws ecr get-login)
  username=$(echo $login_cmd | cut -d " " -f 4)
  password=$(echo $login_cmd | cut -d " " -f 6)
  endpoint=$(echo $login_cmd | cut -d " " -f 9)
  auth=$(echo "$username:$password" | /usr/bin/base64)

  configjson="{ \"auths\": { \"${endpoint}\": { \"auth\": \"${auth}\" } } }"

  kubectl apply -f - << EOF
apiVersion: v1
kind: Secret
metadata:
  name: aws-ecr-registry
data:
  .dockerconfigjson: $(echo $configjson | /usr/bin/base64)
type: kubernetes.io/dockerconfigjson
EOF
}

update_registry_secret
