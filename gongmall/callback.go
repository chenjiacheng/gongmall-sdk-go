package gongmall

type CallbackService service

type CallbackResp struct {
	Success bool `json:"success"`
}

// Success 回调处理成功
func (s *CallbackService) Success() (*CallbackResp, error) {
	resp := CallbackResp{
		Success: true,
	}

	return &resp, nil
}

// Failure 回调处理失败
func (s *CallbackService) Failure() (*CallbackResp, error) {
	resp := CallbackResp{
		Success: false,
	}

	return &resp, nil
}

// VerifySign 回调验证签名
func (s *CallbackService) VerifySign(params map[string]interface{}) bool {
	sign, ok := params["sign"]
	if !ok {
		return false
	}

	delete(params, "sign")

	return sign == s.client.getSign(params, s.client.AppSecret)
}
