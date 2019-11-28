package model

const (
	RegisterCode = 2001 //注册代码
)

func Response(code int, msg string) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["msg"] = msg
	return res
}

func ResponseData(code int, msg string, data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["msg"] = msg
	res["data"] = data
	return res
}
