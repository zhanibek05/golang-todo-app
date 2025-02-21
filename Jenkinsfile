pipeline {
    agent any
    
    environment {
        GO111MODULE = "on"
        GOPATH = "${WORKSPACE}/go"
        PATH = "${GOPATH}/bin:/usr/local/go/bin:${PATH}"
    }

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/zhanibek05/golang-todo-app.git'
            }
        }
        
        stage('Set Up Go Environment') {
            steps {
                sh 'go version'
                sh 'mkdir -p ${GOPATH}/src/github.com/zhanibek05 && ln -s ${WORKSPACE} ${GOPATH}/src/github.com/zhanibek05/golang-todo-app'
            }
        }

        stage('Download Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Run Application') {
            steps {
                sh './app &'
            }
        }
    }

    post {
        always {
            archiveArtifacts artifacts: 'app', fingerprint: true
        }
        success {
            echo "Build and Tests Passed Successfully!"
        }
        failure {
            echo "Build or Tests Failed. Check logs."
        }
    }
}
