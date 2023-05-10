package define

type Request struct {
	Module string         `json:"module"`
	Method string         `json:"method"`
	Data   map[string]any `json:"data"`
}

type Response struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}

type RetResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  string `json:"data"`
}
type OriRequest struct {
	Data  string `json:"data"`
	Debug int    `json:"debug"`
}

type OriJsonRequest struct {
	Data map[string]any `json:"data"`
}
