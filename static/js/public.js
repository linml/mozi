
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


