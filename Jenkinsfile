pipeline {
  agent any 

  stages {
    stage('Build') {
      steps {
        sh 'docker build -t myimage .' 
      }
    }
    
    // stage('Test') {
    //   steps {
    //     sh 'docker run myimage testscript.sh'
    //   }
    // }
    
    stage('Push') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
          sh 'docker login -u $USERNAME -p $PASSWORD'          
          sh 'docker push myimage'
        }
      }
    }
  }
}