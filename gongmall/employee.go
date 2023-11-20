package gongmall

import (
	"encoding/json"

	"github.com/chenjiacheng/gongmall-sdk-go/request"
	"github.com/chenjiacheng/gongmall-sdk-go/response"
)

type EmployeeService service

type GetListReq struct {
	ServiceID string `json:"serviceId,omitempty"`
}

type GetListResp struct {
	response.CommonResp
	Data []struct {
		TemplateID          string `json:"templateId"`
		TemplateName        string `json:"templateName"`
		PlatformCompanyName string `json:"platformCompanyName"`
		ContractAddr        string `json:"contractAddr"`
	} `json:"data"`
}

// GetContractTemplateList 查询合同模板
// https://opendoc.gongmall.com/merchant/dian-qian-he-tong/he-tong-mu-ban-merchant.html
func (s *EmployeeService) GetContractTemplateList(req GetListReq) (*GetListResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		GetListReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(getContractTemplateListURL, buf)
	if err != nil {
		return nil, err
	}

	resp := GetListResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type SyncInfoReq struct {
	ServiceID                string `json:"serviceId,omitempty"`
	Name                     string `json:"name,omitempty"`
	Mobile                   string `json:"mobile,omitempty"`
	CertificateType          int    `json:"certificateType,omitempty"`
	Identity                 string `json:"identity,omitempty"`
	TemplateID               string `json:"templateId,omitempty"`
	BankAccountNo            string `json:"bankAccountNo,omitempty"`
	AlipayAccountNo          string `json:"alipayAccountNo,omitempty"`
	IdentityFrontBase64      string `json:"identityFrontBase64,omitempty"`
	IdentityBackgroundBase64 string `json:"identityBackgroundBase64,omitempty"`
}

type SyncInfoResp struct {
	response.CommonResp
	Data struct {
		ContractID    string `json:"contractId"`
		ProcessStatus int    `json:"processStatus"`
		BankName      string `json:"bankName"`
	} `json:"data"`
}

// SyncInfo 发起签署
// https://opendoc.gongmall.com/merchant/dian-qian-he-tong/yuan-gong-tong-bu-merchant.html
func (s *EmployeeService) SyncInfo(req SyncInfoReq) (*SyncInfoResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		SyncInfoReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(syncInfoURL, buf)
	if err != nil {
		return nil, err
	}

	resp := SyncInfoResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type SignContractReq struct {
	ServiceID  string `json:"serviceId,omitempty"`
	Mobile     string `json:"mobile,omitempty"`
	ContractID string `json:"contractId,omitempty"`
	Captcha    string `json:"captcha,omitempty"`
}

type SignContractResp struct {
	response.CommonResp
}

// SignContract 确认签署
// https://opendoc.gongmall.com/merchant/dian-qian-he-tong/dian-qian-ye-wu-shou-li-merchant.html
func (s *EmployeeService) SignContract(req SignContractReq) (*SignContractResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		SignContractReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(signContractURL, buf)
	if err != nil {
		return nil, err
	}

	resp := SignContractResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type GetContractStatusReq struct {
	ServiceID  string `json:"serviceId,omitempty"`
	Identity   string `json:"identity,omitempty"`
	TemplateID string `json:"templateId,omitempty"`
}

type GetContractStatusResp struct {
	response.CommonResp
	Data struct {
		ProcessStatus int    `json:"processStatus"`
		FileURL       string `json:"fileUrl"`
	} `json:"data"`
}

// GetContractStatus 查询合同状态（使用模板ID）
// https://opendoc.gongmall.com/merchant/dian-qian-he-tong/dian-qian-jie-guo-cha-xun-merchant.html
func (s *EmployeeService) GetContractStatus(req GetContractStatusReq) (*GetContractStatusResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		GetContractStatusReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(getContractStatusURL, buf)
	if err != nil {
		return nil, err
	}

	resp := GetContractStatusResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type GetContractStatusByContractIDReq struct {
	ServiceID  string `json:"serviceId,omitempty"`
	Identity   string `json:"identity,omitempty"`
	ContractID string `json:"contractId,omitempty"`
}

type GetContractStatusByContractIDResp struct {
	response.CommonResp
	Data struct {
		ProcessStatus int    `json:"processStatus"`
		FileURL       string `json:"fileUrl"`
	} `json:"data"`
}

// GetContractStatusByContractID 查询合同状态（使用合同ID）
// https://opendoc.gongmall.com/merchant/dian-qian-he-tong/get-contract-status-merchant.html
func (s *EmployeeService) GetContractStatusByContractID(req GetContractStatusByContractIDReq) (*GetContractStatusByContractIDResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		GetContractStatusByContractIDReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(getContractStatusByContractIDURL, buf)
	if err != nil {
		return nil, err
	}

	resp := GetContractStatusByContractIDResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type DeleteContractReq struct {
	Identity string `json:"identity,omitempty"`
}

type DeleteContractResp struct {
	response.CommonResp
}

// DeleteContract 解除签署
// https://opendoc.gongmall.com/merchant/dian-qian-he-tong/delete-employee-merchant.html
func (s *EmployeeService) DeleteContract(req DeleteContractReq) (*DeleteContractResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		DeleteContractReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(deleteContractURL, buf)
	if err != nil {
		return nil, err
	}

	resp := DeleteContractResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type UpdateEmployeeReq struct {
	Name            string `json:"name,omitempty"`
	Mobile          string `json:"mobile,omitempty"`
	Identity        string `json:"identity,omitempty"`
	BankAccountNo   string `json:"bankAccountNo,omitempty"`
	AlipayAccountNo string `json:"alipayAccountNo,omitempty"`
}

type UpdateEmployeeResp struct {
	response.CommonResp
}

// UpdateEmployee 更新员工默认手机号或者账号
// https://opendoc.gongmall.com/merchant/dian-qian-he-tong/gen-xin-account-merchant.html
func (s *EmployeeService) UpdateEmployee(req UpdateEmployeeReq) (*UpdateEmployeeResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		UpdateEmployeeReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(updateEmployeeURL, buf)
	if err != nil {
		return nil, err
	}

	resp := UpdateEmployeeResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

type AddBankAccountReq struct {
	Name            string `json:"name,omitempty"`
	Mobile          string `json:"mobile,omitempty"`
	Identity        string `json:"identity,omitempty"`
	BankAccountNo   string `json:"bankAccountNo,omitempty"`
	AlipayAccountNo string `json:"alipayAccountNo,omitempty"`
}

type AddBankAccountResp struct {
	response.CommonResp
	Data struct {
		BankAccountNo string `json:"bankAccountNo"`
		BankName      string `json:"bankName"`
	}
}

// AddBankAccount 员工添加银行卡账号
// https://opendoc.gongmall.com/merchant/dian-qian-he-tong/add-employee-bank-account.html
func (s *EmployeeService) AddBankAccount(req AddBankAccountReq) (*AddBankAccountResp, error) {
	buf, _ := json.Marshal(struct {
		request.CommonReq
		AddBankAccountReq
	}{
		s.client.getCommonReq(),
		req,
	})

	respBytes, err := s.client.httpPostForm(addBankAccountURL, buf)
	if err != nil {
		return nil, err
	}

	resp := AddBankAccountResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}
