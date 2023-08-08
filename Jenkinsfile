pipeline {
  // Run on an agent where we want to use Go
  agent any

  // Ensure the desired Go version is installed for all stages,
  // using the name defined in the Global Tool Configuration
  tools { go 'go 1.20.7' }


  stages {

    stage('Go init') {
        steps{
        sh 'cd helloworld/'
        sh 'go mod init helloworld'
        }
    }
     stage('Build') {
      steps {
       sh 'cd helloworld/ && go build -buildmode=pie -buildvcs=false'
      }
     }

    stage('Test') {
      steps {
        sh 'cd helloworld/ && go test'
      }
    }
  }
  
  post {
    always {
        dir('helloworld') {
            archiveArtifacts artifacts: 'helloworld', onlyIfSuccessful: true
        }
    }
  }
}
