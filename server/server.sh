#!/bin/bash

#变量
PROJECTNAME="projectName"
PROJECTBASE="."
PROJECTBIN="$PROJECTBASE/bin"
PROJECTLOGS="$PROJECTBASE/log"
prog=$PROJECTNAME
#获取当前目录
CURDIR=$(dirname $0)
cd $CURDIR

ENVIRONMENT=$2

if [ "$ENVIRONMENT" == "" ]; then
  ENVIRONMENT="development"
fi
#运行服务
start() {
        echo -e "Begin to compile the project ---$PROJECTNAME"

        #编译go项目
        go mod tidy
        go build -o $PROJECTNAME main.go
        #赋予权限
        chmod  777  "$CURDIR/$PROJECTNAME"
        echo "Compilation completed"

        echo -e  "starting project $PROJECTNAME,please waiting..."
        sleep 2
        #后台运行项目
        nohup ./$PROJECTNAME -c=$ENVIRONMENT > $PROJECTLOGS/run.log 2>&1 &
        #获取进程
        pid=$(ps -ef | grep $prog | grep -v grep | awk '{print $2}')
        echo -e  "\033[32mthe project is starting at process : $pid \033[0m"
        echo -e "ok"
}
#暂停服务
stop(){
        echo -e $"Stopping the project ---$prog "

        #获取进程
        pid=$(ps -ef | grep $prog | grep -v grep | awk '{print $2}')
        if [ "$pid" ]; then
#                echo -n $"killing process pid: $pid "
                echo -e  "\033[32mkilling process pid: $pid... \033[0m"

                #杀掉进程
                kill -9 $pid
                ret=0
                #多次循环杀掉进程
                for ((i=1;i<=15;i++)); do
                        sleep 1
                        pid=$(ps -ef | grep $prog | grep -v grep | awk '{print $2}')
                        if [ "$pid" ]; then
                                kill -9 $pid
                                ret=0
                        else
                                ret=1
                                break
                        fi
                done

                if [ "$ret" ]; then
                        echo -e $"the process stopped... "
                else
                        echo -e $"no"
                fi
        else
                echo -e $"no program process to stop"
        fi
}

#重启服务
restart(){
        stop
        sleep 2
        start
}
#判断第一个参数
case "$1" in
start)
        $1
        ;;
stop)
        $1
        ;;
restart)
        $1
        ;;
*)
        echo $"Usage: $0 {start|stop|restart} {development|production}"
        exit 2
        ;;
esac