pipeline {
  // Run on an agent where we want to use Go
  agent any

  // Ensure the desired Go version is installed for all stages,
  // using the name defined in the Global Tool Configuration
  tools { go 'go 1.20.7' }


  stages {
    
     stage('Build') {
      steps {
       sh 'rm -rf build && go build'
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
