pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'ngocthuong/server-golang'
        DOCKER_TAG = 'latest'
    }

    stages {4 
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/NgocThuong134/devopsday3.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                echo 'Deploying to DEV...'
                sh 'docker image pull ngocthuong/server-golang:latest'
                sh 'docker container stop server-golang || echo "this container does not exist"'
                sh 'docker network create dev || echo "this network exists"'
                sh 'echo y | docker container prune '

                sh 'docker container run -d --rm --name server-golang -p 5080:3000 --network dev ngocthuong/server-golang:latest'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}