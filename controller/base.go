package controller

func Success(data interface{}) map[string]interface{} {
	var response map[string]interface{}
	response = make(map[string]interface{})
	response["code"] = 200
	response["msg"] = "操作成功"
	response["data"] = data
	return response
}
func Error(data interface{}) map[string]interface{} {
	var response map[string]interface{}
	response = make(map[string]interface{})
	response["code"] = 400
	response["msg"] = "操作失败"
	response["data"] = data
	return response
}
