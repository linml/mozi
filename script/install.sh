#!/usr/bin/env bash

go_v="go1.11.2.linux-amd64.tar.gz"


function install_golang(){
    curl -sfO https://dl.google.com/go/$go_v && tar -C /usr/local -xzf $go_v && echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
    source /etc/profile
}

function notify() {
    echo "$(date): $1 $2"
    if [ -n "$notify_url" ];then
        curl -X POST -s "${notify_url}$1" --data-urlencode "progress=$2" >/dev/null 2>&1
    fi
}
notify start 安装golang
install_golang
notify success 安装golang