package common

// 接口返回操作码

const (
	CodeFail = 0
	CodeOK   = 1

	CodeUserNotLogin = 2
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
