
function find_code_lotto_list(params, callback) {
    $.get("/lotto/code_lotto/list", params, function (rsp) {
        callback(rsp);
    });
}

function find_code_lotto_type_list(params, callback) {
    $.get("/lotto/code_lotto/list", params, function (rsp) {
        callback(rsp);
    });
}

function get_user_detail(params, callback) {
    $.get("/member/infos", params, function (rsp) {
        callback(rsp);
    });
}

function manual_depost(params, callback) {
    $.post("/money/manual_deposit", params, function (rsp) {
        callback(rsp);
    });
}
