#!/bin/bash

/opt/yozo/bin/gotpl ${CATALINA_HOME}/conf/server.xml.tpl
/opt/yozo/bin/gotpl ${CATALINA_HOME}/webapps/fcscloud/WEB-INF/classes/application.yml.tpl
/opt/yozo/bin/gotpl ${CATALINA_HOME}/webapps/fcscloud/WEB-INF/classes/config.properties.tpl
/opt/yozo/bin/gotpl ${CATALINA_HOME}/webapps/fcscloud/WEB-INF/classes/bootstrap.yml.tpl
/opt/yozo/bin/gotpl ${CATALINA_HOME}/webapps/fcscloud/WEB-INF/classes/convertConfig.properties.tpl

catalina.sh run
