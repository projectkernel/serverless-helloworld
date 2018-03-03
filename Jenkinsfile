pipeline {
    agent docker { image "golang" }
    stages {
        stage("Build") {
            // agent {
            //     docker { image "golang" }
            // }
            steps {
                sh "make"
            }
        }
        // stage("Deploy") {
        //     agent {
        //         docker { image 'node:7-alpine' }
        //     }
        //     steps {
        //         sh "npm install -g serverless"
        //         sh "serverless deploy -v"
        //     }
        // }
    }
}