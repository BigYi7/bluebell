package models

//定义请求的参数结构体

type ParamSignUp struct {
	Username   string `json:"username" binging:"required"`
	Password   string `json:"password" binging:"required"`
	RePassword string `json:"re_password" binging:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binging:"required"`
}
