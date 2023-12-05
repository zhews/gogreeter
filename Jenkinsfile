pipeline {
	agent any
	stages {
		stage("Checkout") {
			steps {
		    		checkout scm
			}
	    	}
	    	stage("Install Verify Tools") {
			steps {
		    		sh "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2"
			}
	    	}
	    	stage("Verify") {
			steps {
		    		sh "./bin/golangci-lint run"
		    		sh "go test ./..."
			}
	    	}
	    	stage("Go Build") {
			steps {
		    		sh "go build -o ./bin/gogreeter ."
			}
	    	}
	    	stage("Container Build") {
			steps {
		    		sh 'docker build -t nexus:8082/gogreeter:latest -t nexus:8082/gogreeter:$GIT_COMMIT -f Containerfile .'
			}
	    	}
		stage("Container Push") {
			when {
				branch 'main'
			}
			environment {
				NEXUS_CREDENTIALS = credentials("nexus-credentials")
			}
			steps {
				sh 'echo $NEXUS_CREDENTIALS_PSW | docker login nexus:8082 -u $NEXUS_CREDENTIALS_USR --password-stdin'
				sh "docker push nexus:8082/gogreeter:latest"
				sh 'docker push nexus:8082/gogreeter:$GIT_COMMIT'
			}
		}
	}
}
