pipeline {
  agent {
    label 'cdnsfw_dockerhub_ami'
  }
  tools {
    go "Go 1.10.3"
  }
  options {
    checkoutToSubdirectory('src/github.com/seizadi/cmdb')
  }
  environment {
    GOPATH = "$WORKSPACE"
  }
  stages {
    stage("Test") {
      steps {
        sh "cd src/github.com/seizadi/cmdb && make test"
      }
    }
  }
}
