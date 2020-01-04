pipeline {
  agent any
  stages {
    stage('Linux64') {
      parallel {
        stage('Linux64') {
          environment {
            GOOS = 'linux'
            GOARCH = 'amd64'
          }
          steps {
            sh '''make prepare
make build'''
            archiveArtifacts(onlyIfSuccessful: true, artifacts: 'BackUploader', caseSensitive: true)
          }
        }

        stage('Linux32') {
          environment {
            GOOS = 'linux'
            GOARCH = '386'
          }
          steps {
            sh '''make prepare
make build'''
            archiveArtifacts(artifacts: 'BackUploader', caseSensitive: true, onlyIfSuccessful: true)
          }
        }

        stage('LinuxARM') {
          environment {
            GOOS = 'linux'
            GOARCH = 'arm'
          }
          steps {
            sh '''make prepare
make build'''
            archiveArtifacts(artifacts: 'BackUploader', caseSensitive: true, onlyIfSuccessful: true)
          }
        }

        stage('Windows64') {
          environment {
            GOOS = 'windows'
            GOARCH = 'amd64'
          }
          steps {
            sh '''make prepare
make build'''
            archiveArtifacts(artifacts: 'BackUploader', caseSensitive: true, onlyIfSuccessful: true)
          }
        }

        stage('OSX64') {
          environment {
            GOOS = 'darwin'
            GOARCH = 'amd64'
          }
          steps {
            sh '''make prepare
make build'''
            archiveArtifacts(artifacts: 'BackUploader', caseSensitive: true, onlyIfSuccessful: true)
          }
        }

        stage('OSX32') {
          environment {
            GOOS = 'darwin'
            GOARCH = '386'
          }
          steps {
            sh '''make prepare
make build'''
            archiveArtifacts(caseSensitive: true, artifacts: 'BackUploader', onlyIfSuccessful: true)
          }
        }

      }
    }

  }
}