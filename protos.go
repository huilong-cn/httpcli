package httpcli

//Wrap niuhe wrap(Rsp)
type Wrap struct {
	Data    interface{} `json:"data"`
	Result  int         `json:"result"`
	Message string      `json:"message"`
}
