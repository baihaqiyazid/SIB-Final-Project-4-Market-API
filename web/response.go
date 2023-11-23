package web

type Response struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Data any `json:"data"`
}

type ResponseCategoryMessage struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Message any `json:"message"`
}
 
type ResponseLogin struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Token any `json:"token"`
}

type ResponseTopup struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Message any `json:"message"`
}

type ResponseError struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Message any `json:"message"`
}