#!/usr/bin/env bash

openresty_filename="openresty-1.11.2.4.tar.gz"
openresty_dir=$(basename $openresty_filename .tar.gz)

function check_or_download_file() {
    [ -f $openresty_filename ] || curl --fail https://openresty.org/download/$openresty_filename -O
    return $?
}

function init_dir() {
    mkdir -p /var/log/nginx/ \
        /usr/share/nginx/ \
        /usr/local/openresty/nginx/lua \
        /usr/local/openresty/nginx/conf/sites-enabled/ \
        /home/nginx_cache
    return $?
}

function install_lib() {
    apt-get -qy install \
        perl \
        make \
        libreadline-dev \
        libncurses5-dev \
        libpcre3 \
        libpcre3-dev \
        libssl-dev \
        libgeoip-dev \
        build-essential \
        lua5.1 \
        liblua5.1-dev \
        luarocks
    return $?
}

function install_openresty_core(){
    tar xzvf $openresty_filename && cd $openresty_dir && ./configure --add-module=3rd_party_modules/ngx_cache_purge-2.3 --with-pcre-jit --with-ipv6 --with-http_gzip_static_module && make && make install
    return $?
}

function install_geoip() {
    cd /usr/share/nginx
    return $?
}

function change_conf() {
    grep -q openresty ~/.bashrc || echo 'export PATH=/usr/local/openresty/bin:/usr/local/openresty/nginx/sbin:$PATH' >> ~/.bashrc
    cat > /usr/local/openresty/nginx/conf/mime.types <<EOF
types {
	text/html                             html htm shtml;
	text/css                              css;
	text/xml                              xml;
	image/gif                             gif;
	image/jpeg                            jpeg jpg;
	application/javascript                js;
	application/atom+xml                  atom;
	application/rss+xml                   rss;

	text/mathml                           mml;
	text/plain                            txt;
	text/vnd.sun.j2me.app-descriptor      jad;
	text/vnd.wap.wml                      wml;
	text/x-component                      htc;

	image/png                             png;
	image/tiff                            tif tiff;
	image/vnd.wap.wbmp                    wbmp;
	image/x-icon                          ico;
	image/x-jng                           jng;
	image/x-ms-bmp                        bmp;
	image/svg+xml                         svg svgz;
	image/webp                            webp;

	application/font-woff                 woff;
	application/java-archive              jar war ear;
	application/json                      json;
	application/mac-binhex40              hqx;
	application/msword                    doc;
	application/pdf                       pdf;
	application/postscript                ps eps ai;
	application/rtf                       rtf;
	application/vnd.apple.mpegurl         m3u8;
	application/vnd.ms-excel              xls;
	application/vnd.ms-fontobject         eot;
	application/vnd.ms-powerpoint         ppt;
	application/vnd.wap.wmlc              wmlc;
	application/vnd.google-earth.kml+xml  kml;
	application/vnd.google-earth.kmz      kmz;
	application/x-7z-compressed           7z;
	application/x-cocoa                   cco;
	application/x-java-archive-diff       jardiff;
	application/x-java-jnlp-file          jnlp;
	application/x-makeself                run;
	application/x-perl                    pl pm;
	application/x-pilot                   prc pdb;
	application/x-rar-compressed          rar;
	application/x-redhat-package-manager  rpm;
	application/x-sea                     sea;
	application/x-shockwave-flash         swf;
	application/x-stuffit                 sit;
	application/x-tcl                     tcl tk;
	application/x-x509-ca-cert            der pem crt;
	application/x-xpinstall               xpi;
	application/xhtml+xml                 xhtml;
	application/xspf+xml                  xspf;
	application/zip                       zip;

	application/octet-stream              bin exe dll;
	application/octet-stream              deb;
	application/octet-stream              dmg;
	application/octet-stream              iso img;
	application/octet-stream              msi msp msm;

	application/vnd.openxmlformats-officedocument.wordprocessingml.document    docx;
	application/vnd.openxmlformats-officedocument.spreadsheetml.sheet          xlsx;
	application/vnd.openxmlformats-officedocument.presentationml.presentation  pptx;

	audio/x-wav                           wav;
	audio/midi                            mid midi kar;
	audio/mpeg                            mp3;
	audio/ogg                             ogg;
	audio/x-m4a                           m4a;
	audio/x-realaudio                     ra;

	video/3gpp                            3gpp 3gp;
	video/mp2t                            ts;
	video/mp4                             mp4;
	video/mpeg                            mpeg mpg;
	video/quicktime                       mov;
	video/webm                            webm;
	video/x-flv                           flv;
	video/x-m4v                           m4v;
	video/x-mng                           mng;
	video/x-ms-asf                        asx asf;
	video/x-ms-wmv                        wmv;
	video/x-msvideo                       avi;
}
EOF

cpu_cores=`expr $(grep processor /proc/cpuinfo | tail -n 1 | cut -d ":" -f2 | tr -d " ") + 1`
cat > /usr/local/openresty/nginx/conf/nginx.conf << EOF
user www-data;
worker_processes $cpu_cores;
worker_rlimit_nofile 65535;
pid /var/run/nginx.pid;

events {
    use epoll;
    worker_connections 65535;
    multi_accept on;
}

http {
    # lua_package_path "/usr/local/openresty/nginx/lua/yywaf/?.lua;/usr/local/openresty/nginx/lua/yywaf/rule/?.lua;;";
    # lua_shared_dict limit 10m;
    # lua_shared_dict ip_blacklist 10m;

    # init_by_lua_file /usr/local/openresty/nginx/lua/yywaf/init.lua;
    # access_by_lua_file /usr/local/openresty/nginx/lua/yywaf/waf.lua;

    default_type application/octet-stream;
    keepalive_timeout 180s;
    server_names_hash_bucket_size 512;
    server_tokens off;
    sendfile on;
    tcp_nodelay on;
    tcp_nopush on;
    types_hash_max_size 2048;

    access_log off;
    error_log /var/log/nginx/error.log;
    log_format main '\$remote_addr - \$remote_user [\$time_local] "\$request" \$status \$body_bytes_sent \$request_time \$upstream_response_time \$upstream_addr "\$http_referer" "\$http_user_agent" "\$http_x_forwarded_for"';
    log_format cache '\$remote_addr - \$remote_user [\$time_local] "\$request" \$status \$body_bytes_sent \$request_time \$upstream_response_time \$upstream_addr "\$http_referer" "\$http_user_agent" "\$http_x_forwarded_for" "\$upstream_cache_status" "\$upstream_http_cache_control" "\$upstream_http_expires"';

    # Buffer header/body
    client_header_buffer_size 128k;
    large_client_header_buffers 4 128k;
    client_body_buffer_size 1024k;

    # Open file cache
    open_file_cache max=1000 inactive=20s;
    open_file_cache_valid 30s;
    open_file_cache_min_uses 1;

    # Gzip
    gzip on;
    gzip_static on;
    gzip_http_version 1.0;
    gzip_disable "MSIE [1-6]\.";
    gzip_vary on;
    gzip_comp_level 2;
    gzip_proxied any;
    gzip_types text/plain text/css application/x-javascript text/xml application/xml application/xml+rss text/javascript;

    # Proxy cache
    proxy_temp_path  /home/nginx_cache/temp_dir/;

    # Configs
    include /usr/local/openresty/nginx/conf/mime.types;
    include /usr/local/openresty/nginx/conf/sites-enabled/default;
    include /usr/local/openresty/nginx/conf/sites-enabled/*.conf;
}
EOF

cat > /etc/init.d/nginx <<EOF
#!/bin/sh

DAEMON=/usr/local/openresty/nginx/sbin/nginx
NAME=nginx
DESC=nginx

# Include nginx defaults if available
if [ -f /etc/default/nginx ]; then
    . /etc/default/nginx
fi

test -x \$DAEMON || exit 0

set -e

. /lib/lsb/init-functions

test_nginx_config() {
    if \$DAEMON -t \$DAEMON_OPTS >/dev/null 2>&1; then
        return 0
    else
        \$DAEMON -t \$DAEMON_OPTS
        return \$?
    fi
}

case "\$1" in
    start)
        echo -n "Starting \$DESC: "
        test_nginx_config
        # Check if the ULIMIT is set in /etc/default/nginx
        if [ -n "\$ULIMIT" ]; then
            # Set the ulimits
            ulimit \$ULIMIT
        fi
        start-stop-daemon --start --quiet --pidfile /var/run/\$NAME.pid \
            --exec \$DAEMON -- \$DAEMON_OPTS || true
        echo "\$NAME."
        ;;

    stop)
        echo -n "Stopping \$DESC: "
        start-stop-daemon --stop --quiet --pidfile /var/run/\$NAME.pid \
            --exec \$DAEMON || true
        echo "\$NAME."
        ;;

    restart|force-reload)
        echo -n "Restarting \$DESC: "
        start-stop-daemon --stop --quiet --pidfile \
            /var/run/\$NAME.pid --exec \$DAEMON || true
        sleep 1
        test_nginx_config
        start-stop-daemon --start --quiet --pidfile \
            /var/run/\$NAME.pid --exec \$DAEMON -- \$DAEMON_OPTS || true
        echo "\$NAME."
        ;;

    reload)
        echo -n "Reloading \$DESC configuration: "
        test_nginx_config
        start-stop-daemon --stop --signal HUP --quiet --pidfile /var/run/\$NAME.pid \
            --exec \$DAEMON || true
        echo "\$NAME."
        ;;

    configtest|testconfig)
        echo -n "Testing \$DESC configuration: "
        if test_nginx_config; then
            echo "\$NAME."
        else
            exit \$?
        fi
        ;;

    status)
        status_of_proc -p /var/run/\$NAME.pid "\$DAEMON" nginx && exit 0 || exit \$?
        ;;
    *)
        echo "Usage: \$NAME {start|stop|restart|reload|force-reload|status|configtest}" >&2
        exit 1
        ;;
esac

exit 0
EOF

    cat > /usr/local/openresty/nginx/conf/sites-enabled/default <<EOF
server {
    server_name _;
    server_tokens off;
    listen 80 default_server;
    listen 8088 default_server;
    return 400;
}
EOF

}

function restart_nginx() {
    export PATH=/usr/local/openresty/bin:/usr/local/openresty/nginx/sbin:$PATH && chmod +x /etc/init.d/nginx && update-rc.d nginx defaults && pkill -9 nginx
    service nginx start
    return $?
}

function show_processlist(){
    echo "show processlist"
    echo ">>>"
    pgrep -lf nginx
    echo "<<<"
    return $?
}

for task in check_or_download_file \
    init_dir \
    install_lib \
    install_openresty_core \
    install_geoip \
    change_conf \
    restart_nginx  \
    show_processlist;do
    $task
    if [ $? -ne 0 ]; then
        echo "$task fail."
        exit 1
    fi
done

# clean
rm -rf $openresty_filename $openresty_dir

echo -e "\ndone."
