package web

type WebResponse struct {//berguna sebagai response 
	Code int	`json:"code"`
	Status string `json:"string"`
	Data interface{} `json:"data"`
}

type WebResponseLogout struct {
	Data interface{} `json:"data"`
}