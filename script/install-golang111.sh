#!/usr/bin/env bash

go_v="go1.11.2.linux-amd64.tar.gz"


function install_golang(){

    curl -sfO https://dl.google.com/go/$go_v && tar -C /usr/local -xzf $go_v
    mkdir -p /home/deploy /home/product /home/go /home/go/src /home/go/bin /home/go/pkg
    echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
    echo "export PATH=\$PATH:/home/go/bin" >> /etc/profile
    echo "export GOPATH=/home/go" >> /etc/profile
    source /etc/profile

}

function notify() {
    echo "$(date): $1 $2"
    if [ -n "$notify_url" ];then
        curl -X POST -s "${notify_url}$1" --data-urlencode "progress=$2" >/dev/null 2>&1
    fi
}

notify start install golang
install_golang
notify success install golang