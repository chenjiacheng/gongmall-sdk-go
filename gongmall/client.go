package gongmall

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/chenjiacheng/gongmall-sdk-go/request"
)

const (
	defaultBaseURL = "https://openapi.gongmall.com"
	sandboxBaseURL = "https://openapi-qa.gongmall.com"

	getContractTemplateListURL       = "/api/merchant/contract/getList"
	getContractStatusURL             = "/api/merchant/employee/getContractStatus"
	getContractStatusByContractIDURL = "/api/merchant/employee/getContractStatusByContractId"
	syncInfoURL                      = "/api/merchant/employee/syncInfo"
	signContractURL                  = "/api/merchant/employee/signContract"
	deleteContractURL                = "/api/merchant/employee/deleteContract"
	updateEmployeeURL                = "/api/merchant/employee/updateEmployee"
	addBankAccountURL                = "/api/merchant/employee/addBankAccount"
	doSinglePaymentURL               = "/api/merchant/doSinglePayment"
	transQueryURL                    = "/api/merchant/transQuery"
	queryLimitURL                    = "/api/merchant/queryLimit"
	receiptQueryURL                  = "/api/merchant/receiptQuery"
	queryBalanceURL                  = "/api/merchant/queryBalance"
	queryMonthSumURL                 = "/api/merchant/tax/queryMonthSum"
	queryMonthURL                    = "/api/merchant/tax/queryMonth"
	rechargeInfoURL                  = "/api/merchant/account/rechargeInfo"
	merchantIncomeAccountInfoURL     = "/api/merchant/account/merchantIncomeAccountInfo"
)

type Client struct {
	AppKey    string
	AppSecret string
	ServiceId string
	Sandbox   bool
	Timeout   time.Duration

	common service

	Account  *AccountService
	Employee *EmployeeService
	Merchant *MerchantService
}

type service struct {
	client *Client
}

// NewClient 实例化客户端
func NewClient(appKey, appSecret, serviceId string) *Client {
	c := &Client{
		AppKey:    appKey,
		AppSecret: appSecret,
		ServiceId: serviceId,
	}

	c.common.client = c
	c.Account = (*AccountService)(&c.common)
	c.Employee = (*EmployeeService)(&c.common)
	c.Merchant = (*MerchantService)(&c.common)

	return c
}

// NewSandboxClient 实例化沙箱环境客户端
func NewSandboxClient(appKey, appSecret, serviceId string) *Client {
	c := NewClient(appKey, appSecret, serviceId)
	c.Sandbox = true
	return c
}

// 发送 POST 请求
func (c *Client) httpPostForm(urlPath string, buf []byte) ([]byte, error) {
	baseUrl := defaultBaseURL
	if c.Sandbox {
		baseUrl = sandboxBaseURL
	}

	httpClient := http.Client{
		Timeout: c.Timeout,
	}

	var resultMap map[string]interface{}
	_ = json.Unmarshal(buf, &resultMap)

	resultMap["sign"] = c.getSign(resultMap, c.AppSecret)
	values := url.Values{}
	for k, v := range resultMap {
		values.Set(k, fmt.Sprintf("%v", v))
	}

	resp, err := httpClient.PostForm(baseUrl+urlPath, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBytes))
	return respBytes, nil
}

// 获取公共请求参数
func (c *Client) getCommonReq() request.CommonReq {
	return request.CommonReq{
		AppKey:    c.AppKey,
		Nonce:     c.getNonce(),
		Timestamp: c.getTimestamp(),
	}
}

// 获取随机数
func (c *Client) getNonce() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

// 获取当前毫秒时间戳
func (c *Client) getTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

// 获取签名
func (c *Client) getSign(params map[string]interface{}, appSecret string) string {
	// 按参数名升序排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接参数名和参数值，并添加 appSecret
	var text string
	for _, k := range keys {
		text += k + "=" + fmt.Sprintf("%v", params[k]) + "&"
	}
	text += "appSecret=" + appSecret

	// 计算签名（MD5）
	hash := md5.Sum([]byte(text))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}
