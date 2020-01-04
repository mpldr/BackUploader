pipeline {
  agent any
  stages {
    stage('Linux') {
      parallel {
        stage('Linux') {
          steps {
            sh 'export GOOS=linux'
          }
        }

        stage('Windows') {
          steps {
            sh 'export GOOS=windows'
          }
        }

        stage('MacOS') {
          steps {
            sh 'export GOOS=darwin'
          }
        }

      }
    }

    stage('64-bit') {
      parallel {
        stage('64-bit') {
          steps {
            sh 'export GOARCH=amd64'
          }
        }

        stage('32-bit') {
          steps {
            sh 'export GOARCH=386'
          }
        }

        stage('ARM') {
          steps {
            sh 'export GOARCH=arm'
          }
        }

      }
    }

    stage('get dependencies') {
      steps {
        sh 'make prepare'
      }
    }

    stage('Build') {
      steps {
        sh 'make build'
      }
    }

  }
}