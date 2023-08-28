#!/bin/bash

set -x

APP_NAME=jax
SUPERVISOR=supervise.${APP_NAME}
WORKSPACE=/home/anxinsec/goapp
LOG=/home/anxinsec/goapp/${APP_NAME}/status

reload() {
  #  查看supervisor是否存在
  pid=`ps -ef | grep "${SUPERVISOR} " | grep -v "grep" |  awk '{print $2}'`

  if [ ! -d $LOG ]; then
    mkdir -p ${LOG}/${APP_NAME}
    touch ${LOG}/${APP_NAME}/dump.log
    chmod -R 755 $LOG
  fi

  # 设置log文件夹可读写并清空下上次的log日志
  cd ${LOG}/${APP_NAME} && rm -rf supervise.log supervise.log.wf

  # 进入项目目录，由于目录错误   导致go获取配置文件错误
  cd ${WORKSPACE}/${APP_NAME}
  if [ "pid$pid" == "pid" ]; then
    #  不存在启动supervisor进程
    ${WORKSPACE}/${APP_NAME}/${SUPERVISOR} -p ${LOG}/${APP_NAME} -f ${WORKSPACE}/${APP_NAME}/${APP_NAME} -t 60 >${LOG}/${APP_NAME}/dump.log 2>&1
  else
    # 如果存在则直接kill app进程即可
    pid=$(ps -ef | egrep ${APP_NAME}$ |  awk '{print $2}')
    if [ "pid$pid" != "pid" ]; then
      kill "$pid"
    fi
  fi

  # 设置supervisor日志权限可查看
  cd ${LOG} && chmod -R 755 ${APP_NAME}
}

stop() {
  pid=$(ps -ef | egrep ${APP_NAME}$ |  awk '{print $2}')
  if [ "pid$pid" != "pid" ]; then
    kill "$pid"
  fi
}

case "$1" in
stop)
  stop
  ;;
reload)
  reload
  ;;
restart)
  stop
  reload
  ;;
*)
  echo "Usage: $0 {stop|reload|restart}"
  ;;
esac