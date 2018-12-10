#!/usr/bin/env bash

############################################################
#
# 清除记录数据
#
############################################################

server_ip=`hostname -I | grep -E -o '[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+'`
read -p "please input the server_ip: " input_ip
if echo $server_ip | grep -q $input_ip >/dev/null;then
    cat <<EOF | mysql -uroot -p mozi

DROP TABLE lotto_result_1;
DROP TABLE lotto_result_2;
DROP TABLE lotto_result_3;
DROP TABLE lotto_result_4;
DROP TABLE lotto_result_5;
DROP TABLE lotto_result_6;
DROP TABLE lotto_result_7;
DROP TABLE lotto_result_8;
DROP TABLE lotto_result_9;
DROP TABLE lotto_result_10;
DROP TABLE lotto_result_11;
DROP TABLE lotto_result_12;
DROP TABLE lotto_result_13;
DROP TABLE lotto_result_14;
DROP TABLE lotto_result_15;
DROP TABLE lotto_result_16;
DROP TABLE lotto_result_17;
DROP TABLE lotto_result_18;
DROP TABLE lotto_result_19;
DROP TABLE lotto_result_20;
DROP TABLE lotto_result_21;
DROP TABLE lotto_result_22;
DROP TABLE lotto_result_23;
DROP TABLE lotto_result_24;
DROP TABLE lotto_result_25;
DROP TABLE lotto_result_26;
DROP TABLE lotto_result_27;
DROP TABLE lotto_result_28;
DROP TABLE lotto_result_29;
DROP TABLE lotto_result_30;
DROP TABLE lotto_result_31;
DROP TABLE lotto_result_32;
DROP TABLE lotto_result_33;
DROP TABLE lotto_result_34;
DROP TABLE lotto_result_35;
DROP TABLE lotto_result_36;
DROP TABLE lotto_result_37;
DROP TABLE lotto_result_38;
DROP TABLE lotto_result_39;
DROP TABLE lotto_result_40;

DELETE FROM record_log_admin_action;
DELETE FROM record_log_user_action;
DELETE FROM record_lotto_order;
DELETE FROM record_money_change;
DELETE FROM record_user_login;

DELETE FROM report_lotto_day_count;

DELETE FROM users WHERE user_id > 1;
DELETE FROM user_wallet WHERE user_id > 1;
DELETE FROM user_score WHERE user_id > 1;
DELETE FROM user_relation WHERE user_id > 1;
DELETE FROM user_profile WHERE user_id > 1;
DELETE FROM user_power WHERE user_id > 1;
DELETE FROM user_links WHERE user_id > 1;
DELETE FROM user_bank WHERE user_id > 1;
EOF
    echo "success."
else
    echo "aborted. confirm failure."
fi
