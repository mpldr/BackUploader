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
make jenkins'''
            archiveArtifacts(onlyIfSuccessful: true, artifacts: 'BackUploader.linux.amd64', caseSensitive: true)
          }
        }

        stage('Linux32') {
          environment {
            GOOS = 'linux'
            GOARCH = '386'
          }
          steps {
            sh '''make prepare
make jenkins'''
            archiveArtifacts(artifacts: 'BackUploader.linux.386', caseSensitive: true, onlyIfSuccessful: true)
          }
        }

        stage('LinuxARM') {
          environment {
            GOOS = 'linux'
            GOARCH = 'arm'
          }
          steps {
            sh '''make prepare
make jenkins'''
            archiveArtifacts(artifacts: 'BackUploader.linux.arm', caseSensitive: true, onlyIfSuccessful: true)
          }
        }

        stage('Windows64') {
          environment {
            GOOS = 'windows'
            GOARCH = 'amd64'
          }
          steps {
            sh '''make prepare
make jenkins'''
            archiveArtifacts(artifacts: 'BackUploader.windows.amd64.exe', caseSensitive: true, onlyIfSuccessful: true)
          }
        }

        stage('OSX64') {
          environment {
            GOOS = 'darwin'
            GOARCH = 'amd64'
          }
          steps {
            sh '''make prepare
make jenkins'''
            archiveArtifacts(artifacts: 'BackUploader.darwin.amd64', caseSensitive: true, onlyIfSuccessful: true)
          }
        }

        stage('OSX32') {
          environment {
            GOOS = 'darwin'
            GOARCH = '386'
          }
          steps {
            sh '''make prepare
make jenkins'''
            archiveArtifacts(caseSensitive: true, artifacts: 'BackUploader.darwin.386', onlyIfSuccessful: true)
          }
        }

      }
    }

  }
}
