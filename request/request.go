package request

type CommonReq struct {
	AppKey    string `json:"appKey"`
	Nonce     string `json:"nonce"`
	Timestamp string `json:"timestamp"`
}
