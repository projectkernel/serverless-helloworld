pipeline {
    agent none
    stages {
        stage("Build") {
            agent {
                docker { image "golang" }
            }
            steps {
                sh "make"
            }
        }
        stage("Unit Test") {
            agent {
                docker { image "golang" }
            }
            steps {
                sh "./scripts/test.sh"
            }
        }
        stage("Dev Deploy") {
            agent {
                docker { image 'node:7-alpine' }
            }
            steps {
                sh "npm install -g serverless"
                sh "serverless deploy -v"
            }
        }
        stage("Integration Test") {
            agent {
                docker { image "golang" }
            }
            steps {
                sh "./scripts/test.sh"
            }
        }
    }
}