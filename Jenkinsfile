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
		    		sh "docker build -t gogreeter -f Containerfile ."
			}
	    	}
	}
}
