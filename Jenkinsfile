pipeline {
  // Run on an agent where we want to use Go
  agent any

  // Ensure the desired Go version is installed for all stages,
  // using the name defined in the Global Tool Configuration
  tools { go 'go 1.20.7' }


  stages {
    
     stage('Build') {
      steps {
       sh 'rm -rf build && go build -o myprogram'
      }
     }

    stage('Test') {
      steps {
        sh 'go test'
      }
    }
  }
  
  post {
    always {
            archiveArtifacts artifacts: 'myprogram', onlyIfSuccessful: true
        }
    }
  }
}
