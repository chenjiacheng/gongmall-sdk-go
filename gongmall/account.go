package gongmall

import (
	"encoding/json"

	"github.com/chenjiacheng/gongmall-sdk-go/request"
	"github.com/chenjiacheng/gongmall-sdk-go/response"
)

type AccountService service

type QueryBalanceReq struct {
	ServiceID string `json:"serviceId,omitempty"`
}

type QueryBalanceResp struct {
	response.CommonResp
	Data struct {
		TotalAvailableBalance float64 `json:"totalAvailableBalance"`
		TotalFrozenAmount     float64 `json:"totalFrozenAmount"`
		AccountInfoList       []struct {
			AccountNo                   string  `json:"accountNo"`
			AccountChannelGroupType     string  `json:"accountChannelGroupType"`
			AccountChannelGroupTypeName string  `json:"accountChannelGroupTypeName"`
			AvailableBalance            float64 `json:"availableBalance"`
			FrozenAmount                float64 `json:"frozenAmount"`
			Opened                      bool    `json:"opened"`
			Status                      int     `json:"status"`
		} `json:"accountInfoList"`
	} `json:"data"`
}

// QueryBalance 查询商户当前余额
// https://opendoc.gongmall.com/merchant/merchant-account/cha-xun-qi-ye-yu-e-merchant.html
func (s *AccountService) QueryBalance(req QueryBalanceReq) (*QueryBalanceResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		QueryBalanceReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(queryBalanceURL, buf)
	if err != nil {
		return nil, err
	}

	resp := QueryBalanceResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type QueryMonthSumReq struct {
	ServiceID        string `json:"serviceId,omitempty"`
	TransactionMonth string `json:"transactionMonth,omitempty"`
	Offset           int    `json:"offset,omitempty"`
	Limit            int    `json:"limit,omitempty"`
}

type QueryMonthSumResp struct {
	response.CommonResp
	Data struct {
		TotalAmount    float64 `json:"totalAmount"`
		TotalFeeAmount float64 `json:"totalFeeAmount"`
	} `json:"data"`
}

// QueryMonthSum 查询月税费汇总-按进单时间统计
// https://opendoc.gongmall.com/merchant/merchant-account/get-month-tax-fee.html
func (s *AccountService) QueryMonthSum(req QueryMonthSumReq) (*QueryMonthSumResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		QueryMonthSumReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(queryMonthSumURL, buf)
	if err != nil {
		return nil, err
	}

	resp := QueryMonthSumResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type QueryMonthReq struct {
	ServiceID        string `json:"serviceId,omitempty"`
	TransactionMonth string `json:"transactionMonth,omitempty"`
}

type QueryMonthResp struct {
	response.CommonResp
	Data struct {
		Total int `json:"total"`
		List  []struct {
			Identity         string  `json:"identity"`
			TransactionMonth string  `json:"transactionMonth"`
			TotalCount       int     `json:"totalCount"`
			TotalAmount      float64 `json:"totalAmount"`
			TotalFeeAmount   float64 `json:"totalFeeAmount"`
		} `json:"list"`
	} `json:"data"`
}

// QueryMonth 查询月税费明细-按进单时间
// https://opendoc.gongmall.com/merchant/merchant-account/get-month-detail-tax-fee.html
func (s *AccountService) QueryMonth(req QueryMonthReq) (*QueryMonthResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		QueryMonthReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(queryMonthURL, buf)
	if err != nil {
		return nil, err
	}

	resp := QueryMonthResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type RechargeInfoResp struct {
	response.CommonResp
	Data struct {
		RechargeInfoList []struct {
			Province       string `json:"province"`
			AccountName    string `json:"accountName"`
			City           string `json:"city"`
			AccountNo      string `json:"accountNo"`
			BankNo         string `json:"bankNo"`
			ChannelType    string `json:"channelType"`
			BankBranchName string `json:"bankBranchName,omitempty"`
		} `json:"rechargeInfoList"`
		ServiceID string `json:"serviceId"`
	} `json:"data"`
}

// RechargeInfo 查询商户充值信息
// https://opendoc.gongmall.com/merchant/merchant-account/account-merchant-recharge-info.html
func (s *AccountService) RechargeInfo() (*RechargeInfoResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
	}{
		s.client.getCommonReq(),
	})

	respBytes, err := s.client.httpPostForm(rechargeInfoURL, buf)
	if err != nil {
		return nil, err
	}

	resp := RechargeInfoResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type MerchantIncomeAccountInfoResp struct {
	response.CommonResp
	Data struct {
		SpName       string  `json:"spName"`
		Balance      float64 `json:"balance"`
		AccountNo    string  `json:"accountNo"`
		MerchantName string  `json:"merchantName"`
		FrozenAmount float64 `json:"frozenAmount"`
		Status       int     `json:"status"`
	} `json:"data"`
}

// MerchantIncomeAccountInfo 查询其他户账户信息
// https://opendoc.gongmall.com/merchant/merchant-account/merchant-income-account-info.html
func (s *AccountService) MerchantIncomeAccountInfo() (*MerchantIncomeAccountInfoResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
	}{
		s.client.getCommonReq(),
	})

	respBytes, err := s.client.httpPostForm(merchantIncomeAccountInfoURL, buf)
	if err != nil {
		return nil, err
	}

	resp := MerchantIncomeAccountInfoResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}
