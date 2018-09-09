package routes

import "github.com/gin-gonic/gin"

type ShowResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ParamHelper map[string]string

func (g ParamHelper) GetQuery(c *gin.Context, key string) {
	val, b := c.GetQuery(key)
	if true == b {
		g[key] = val
	}
}

func (g ParamHelper) GetQueryNotEmpty(c *gin.Context, key string) {
	val, b := c.GetQuery(key)
	if true == b && "" != val {
		g[key] = val
	}
}

func (g ParamHelper) GetPostForm(c *gin.Context, key string) {
	val, b := c.GetPostForm(key)
	if true == b {
		g[key] = val
	}
}

func (g ParamHelper) GetPostFormNotEmpty(c *gin.Context, key string) {
	val, b := c.GetPostForm(key)
	if true == b && "" != val {
		g[key] = val
	}
}

func (g ParamHelper) Exist(key string) bool {
	if _, ok := g[key]; ok {
		return true
	}
	return false
}

func (g ParamHelper) Get(key string) string {
	if val, ok := g[key]; ok {
		return val
	}
	return ""
}

func (g ParamHelper) Set(key string, val string) {
	g[key] = val
}

func (g ParamHelper) Del(key string) {
	delete(g, key)
}

func ApiShowResult(code int, message string) ShowResult {
	return ShowResult{Code: code, Msg: message}
}

func ApiResult(code int, message string, data interface{}) Result {
	return Result{Code: code, Msg: message, Data: data}
}
