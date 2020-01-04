pipeline {
  agent any
  stages {
    stage('get dependencies') {
      steps {
        sh 'go get -v ./...'
      }
    }

    stage('Build') {
      steps {
        sh 'make build'
      }
    }

  }
}