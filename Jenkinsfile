pipeline {
  agent any
  environment {
    name_final = "sgp-processor-svc"
    DB_CREDS = credentials('sgpprocessorsvc')
  }
  stages {
    stage('Docker Build') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          if (result != "") {
            sh '''
            docker stop ${name_final}
            docker rm -vf ${name_final}
            docker build . -t ${name_final}
            docker system prune -f
	    '''
          } else {
            sh '''
            docker build . -t ${name_final}
            docker system prune -f
	    '''
          }
        }
      }
    }
    stage('SonarQube Analysis') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        echo 'SonarQube'
      }
    }
    stage('RUN DB DEV') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          sh '''
          docker run --rm flyway/flyway:9.8.3 version
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:9.8.3 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW migrate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:9.8.3 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW validate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:9.8.3 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW info
	  '''
        }
      }
    }
    stage('Deploy to DEV') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          sh '''
          docker run -dt -p 30006:90 --name ${name_final} ${name_final}
          docker system prune -f
	  '''
        }
      }
    }
    stage('Cucumber Tests DEV') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        echo 'Cucumber Tests'
      }
    }
    stage('RUN DB QA') {
      agent {
        label 'qa'
      }
      when {
        anyOf {
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          sh '''
          docker run --rm flyway/flyway:9.8.3 version
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:9.8.3 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW migrate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:9.8.3 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW validate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:9.8.3 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW info
	  '''
        }
      }
    }
    stage('Deploy to QA') {
      agent {
        label 'qa'
      }
      when {
        anyOf {
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          if (result != "") {
            sh '''
            docker stop ${name_final}
            docker rm -vf ${name_final}
            docker build . -t ${name_final}
            docker run -dt -p 30106:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          } else {
            sh '''
            docker build . -t ${name_final}
            docker run -dt -p 30106:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          }
        }
      }
    }
    stage('QA Approval') {
      agent {
        label 'prd'
      }
      when {
          branch 'master'
      }
      steps {
        input "Aprobacion Tester QA"
      }
    }
    stage('RUN DB PRD') {
      agent {
        label 'prd'
      }
      when {
          branch 'master'
      }
      steps {
        script {
          sh '''
          docker run --rm flyway/flyway:9.8.3 version
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:9.8.3 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW migrate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:9.8.3 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW validate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:9.8.3 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW info
	  '''
        }
      }
    }
    stage('Deploy to PRD') {
      agent {
        label 'prd'
      }
      when {
          branch 'master'
      }
      steps {
        script {
          def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          if (result != "") {
            sh '''
            docker stop ${name_final}
            docker rm -vf ${name_final}
            docker build . -t ${name_final}
            docker run -dt -p 30206:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          } else {
            sh '''
            docker build . -t ${name_final}
            docker run -dt -p 30206:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          }
        }
      }
    }
  }
}
