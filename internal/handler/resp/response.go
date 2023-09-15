package resp

type Response struct {
	HTTPStatusCode int         `json:"status_code"`
	Value          interface{} `json:"body"`
}

type ResponseWithId struct {
	HTTPStatusCode int `json:"status_code"`
	Id             int `json:"id"`
}
