package gongmall

import (
	"encoding/json"

	"github.com/chenjiacheng/gongmall-sdk-go/request"
	"github.com/chenjiacheng/gongmall-sdk-go/response"
)

type MerchantService service

type DoSinglePaymentReq struct {
	ServiceID           string  `json:"serviceId,omitempty"`
	RequestID           string  `json:"requestId,omitempty"`
	Mobile              string  `json:"mobile,omitempty"`
	Name                string  `json:"name,omitempty"`
	Amount              float64 `json:"amount,omitempty"`
	Identity            string  `json:"identity,omitempty"`
	BankAccount         string  `json:"bankAccount,omitempty"`
	WithdrawCallbackUrl string  `json:"withdrawCallbackUrl,omitempty"`
	DateTime            string  `json:"dateTime,omitempty"`
	Remark              string  `json:"remark,omitempty"`
	ExtRemark           string  `json:"extRemark,omitempty"`
	SalaryType          string  `json:"salaryType,omitempty"`
	AlipayAccountType   string  `json:"alipayAccountType,omitempty"`
	WxAppid             string  `json:"wxAppid,omitempty"`
}

type DoSinglePaymentResp struct {
	response.CommonResp
	Data struct {
		RequestID   string `json:"requestId"`
		AppmentTime string `json:"AppmentTime"`
	}
}

// DoSinglePayment 提现
// https://opendoc.gongmall.com/merchant/shi-shi-ti-xian/ti-xian-merchant.html
func (s *MerchantService) DoSinglePayment(req DoSinglePaymentReq) (*DoSinglePaymentResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		DoSinglePaymentReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(doSinglePaymentURL, buf)
	if err != nil {
		return nil, err
	}

	resp := DoSinglePaymentResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type transQueryReq struct {
	RequestID string `json:"requestId,omitempty"`
}

type transQueryResp struct {
	response.CommonResp
	Data struct {
		RequestID          string  `json:"requestId"`
		InnerTradeNo       string  `json:"innerTradeNo"`
		Status             int     `json:"status"`
		FailReason         string  `json:"failReason"`
		Mobile             string  `json:"mobile"`
		Name               string  `json:"name"`
		Amount             float64 `json:"amount"`
		CurrentRealWage    float64 `json:"currentRealWage"`
		CurrentTax         float64 `json:"currentTax"`
		CurrentManageFee   float64 `json:"currentManageFee"`
		CurrentAddTax      float64 `json:"currentAddTax"`
		CurrentAddValueTax float64 `json:"currentAddValueTax"`
		Identity           string  `json:"identity"`
		BankName           string  `json:"bankName"`
		BankAccount        string  `json:"bankAccount"`
		DateTime           string  `json:"dateTime"`
		CreateTime         string  `json:"createTime"`
		PayTime            string  `json:"payTime"`
		Remark             string  `json:"remark"`
	}
}

// transQuery 查询单笔提现结果
// https://opendoc.gongmall.com/merchant/shi-shi-ti-xian/cha-xun-ti-xian-jie-guo-merchant.html
func (s *MerchantService) transQuery(req transQueryReq) (*transQueryResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		transQueryReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(transQueryURL, buf)
	if err != nil {
		return nil, err
	}

	resp := transQueryResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type QueryLimitReq struct {
	ServiceID string `json:"serviceId,omitempty"`
	Identity  string `json:"identity,omitempty"`
}

type QueryLimitResp struct {
	response.CommonResp
	Data struct {
		PayLimitAmount float64 `json:"payLimitAmount"`
		IssuedAmount   float64 `json:"issuedAmount"`
		AllowAmount    float64 `json:"allowAmount"`
	}
}

// queryLimit 查询月度可发放限额
// https://opendoc.gongmall.com/merchant/shi-shi-ti-xian/cha-xun-ti-xian-jie-guo-merchant.html
func (s *MerchantService) queryLimit(req QueryLimitReq) (*QueryLimitResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		QueryLimitReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(queryLimitURL, buf)
	if err != nil {
		return nil, err
	}

	resp := QueryLimitResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type ReceiptQueryReq struct {
	RequestID string `json:"requestId,omitempty"`
}

type ReceiptQueryResp struct {
	response.CommonResp
	Data struct {
		RequestID   string `json:"requestId,omitempty"`
		Status      int    `json:"status"`
		DownloadURL string `json:"downloadUrl"`
		PayTime     string `json:"payTime"`
	}
}

// ReceiptQuery 查询单笔提现回单
// https://opendoc.gongmall.com/merchant/shi-shi-ti-xian/order-receipt.html
func (s *MerchantService) ReceiptQuery(req ReceiptQueryReq) (*ReceiptQueryResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		ReceiptQueryReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(receiptQueryURL, buf)
	if err != nil {
		return nil, err
	}

	resp := ReceiptQueryResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}
