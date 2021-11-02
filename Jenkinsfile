pipeline {
    agent any
    stages {
        stage('Test Command') {
            steps {
                sh 'whoami'
                sh 'microk8s inspect'
                sh 'pwd'
                // sh 'kubectl run redis --image=redis' 
                // sh 'micr ok8s kubectl get node'
            }
        }
    }
}