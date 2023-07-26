pipeline {
  agent any 

  stages {
    stage('Build') {
      steps {
        sh 'echo "DB_USERNAME=\"' + ${env.DB_USERNAME} 
          + '\"\nDB_PASSWORD=\"' + ${env.DB_PASSWORD} 
          + '\"\nDB_NAME=\"' + ${env.DB_NAME} 
          + '\"\nDB_HOST=\"' + ${env.DB_HOST} 
          + '\"\nDB_PORT=\"' + ${env.DB_PORT} 
          + '\"" > app.env'
        sh 'docker build -t file_gen .' 
      }
    }
    
    stage('Push') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
          sh 'docker login -u $USERNAME -p $PASSWORD'          
          sh 'docker push file_gen'
        }
      }
    }
  }
}