package corll

type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/* user */
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Gender   int    `json:"gender"`
}
