package d

func ReturnJson(code int,msg string,data ...interface{})(jsonData map[string]interface{}){
	jsonData = make(map[string]interface{}, 3)
	jsonData["code"] = code
	jsonData["msg"] = msg
	if len(data) > 0 && data[0] !=nil{
		jsonData["data"] = data[0]
		jsonData["count"] = data[1]
	}
	return
}

func TableJson(data interface{},col... interface{})(jsonData map[string]interface{}){
	jsonData = make(map[string]interface{}, 3)
	jsonData["rows"] = data
	if len(col)>0{
		jsonData["limit"] = col[0]
		jsonData["offset"] = col[1]
		jsonData["total"] = col[2]
	}
	return
}