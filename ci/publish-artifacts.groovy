#!/usr/bin/env groovy

@Library('sec_ci_libs@v2-latest') _

pipeline {
  agent none

  options {
    timeout(time: 2, unit: 'HOURS')
  }

  stages {
    stage("Update https://downloads.dcos.io/cli/index.html") {
      agent { label 'py36' }

      steps {
        withCredentials([
          [$class: 'AmazonWebServicesCredentialsBinding',
          credentialsId: 'Team13JenkinsCLIReadOnly',
          accessKeyVariable: 'AWS_READ_ACCESS_KEY_ID',
          secretKeyVariable: 'AWS_READ_SECRET_ACCESS_KEY'],
          [$class: 'AmazonWebServicesCredentialsBinding',
          credentialsId: 'Team13JenkinsCLIWrite',
          accessKeyVariable: 'AWS_WRITE_ACCESS_KEY_ID',
          secretKeyVariable: 'AWS_WRITE_SECRET_ACCESS_KEY']
        ]) {
            sh '''
              bash -exc " \
                cd ci; \
                python -m venv env; \
                source env/bin/activate; \
                pip install --upgrade pip setuptools; \
                pip install -r requirements.txt; \
                cd index; \
                ./publish_artifacts.py"
            '''
        }
      }
    }
  }
}
