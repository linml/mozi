package common

// 接口返回操作码

const (
	CodeOK   = 0
	CodeFail = 100

	CodeUserNotLogin     = 101
	CodeUserLoginExpired = 102
)

var codeText = map[int]string{
	CodeFail:         "操作失败",
	CodeOK:           "操作成功",
	CodeUserNotLogin: "尚未登录",
}

// 返回接口请求信息。
// 如果不存在返回空
func CodeText(code int) string {
	return codeText[code]
}

const SID = "SID"
