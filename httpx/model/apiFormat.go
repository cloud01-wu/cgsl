package model

type Request struct {
	Desire interface{} `json:"desire,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type Response struct {
	Meta   *Meta       `json:"meta,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Errors []Error     `json:"errors,omitempty"`
}

type Error struct {
	Status int    `json:"status"`
	Code   string `json:"code,omitempty"`
	Detail string `json:"detail"`
}

type Meta struct {
	From  int `json:"from"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
