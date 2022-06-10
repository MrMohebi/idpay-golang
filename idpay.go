package idpay

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// NewClient create new idpay client
func NewClient(apikey string, isTest bool) Client {
	return Client{APIKEY: apikey, host: "https://api.idpay.ir/v1.1", isTestMode: isTest}
}

func (c *Client) SetHost(host string) {
	c.host = host
}

// CreatePayment create payment link
func (c *Client) CreatePayment(orderId string, amount int64, callBack string, name string, phone string, mail string, desc string) CreatePaymentRes {
	url := c.host + "/payment"

	requestBody := CreatePaymentReq{
		OrderId:  orderId,
		Amount:   amount,
		Callback: callBack,
		Name:     name,
		Phone:    phone,
		Mail:     mail,
		Desc:     desc,
	}

	var data bytes.Buffer
	err := json.NewEncoder(&data).Encode(requestBody)
	if err != nil {
		log.Println(err)
		return CreatePaymentRes{ReqStatus: requestStatus{Success: false, Message: err.Error()}}
	}

	rq, err := http.NewRequest("POST", url, &data)
	if err != nil {
		log.Println(err)
		return CreatePaymentRes{ReqStatus: requestStatus{Success: false, Message: err.Error()}}
	}

	rq.Header.Set("Cache-Control", "no-cache")
	rq.Header.Set("Content-Type", "application/json")
	rq.Header.Set("X-API-KEY", c.APIKEY)
	rq.Header.Set("X-SANDBOX", strconv.FormatBool(c.isTestMode))

	rp, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		return CreatePaymentRes{ReqStatus: requestStatus{Success: false, Message: err.Error()}}
	}
	defer rp.Body.Close()

	var response CreatePaymentRes

	err = json.NewDecoder(rp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return CreatePaymentRes{ReqStatus: requestStatus{Success: true, Message: err.Error()}}
	}

	response.ReqStatus.Success = rp.StatusCode == http.StatusCreated

	return response
}

// Verify verify payment
func (c *Client) Verify(id string, orderId string) VerifyRes {
	url := c.host + "/payment/verify"

	requestBody := VerifyReq{
		OrderId: orderId,
		Id:      id,
	}
	var data bytes.Buffer
	err := json.NewEncoder(&data).Encode(requestBody)
	if err != nil {
		log.Println(err)
		return VerifyRes{ReqStatus: requestStatus{Success: false, Message: err.Error()}}
	}

	rq, err := http.NewRequest("POST", url, &data)
	if err != nil {
		log.Println(err)
		return VerifyRes{ReqStatus: requestStatus{Success: false, Message: err.Error()}}
	}

	rq.Header.Set("Cache-Control", "no-cache")
	rq.Header.Set("Content-Type", "application/json")
	rq.Header.Set("X-API-KEY", c.APIKEY)
	rq.Header.Set("X-SANDBOX", strconv.FormatBool(c.isTestMode))

	rp, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		return VerifyRes{ReqStatus: requestStatus{Success: false, Message: err.Error()}}
	}
	defer rp.Body.Close()

	var response VerifyRes

	err = json.NewDecoder(rp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return VerifyRes{ReqStatus: requestStatus{Success: true, Message: err.Error()}}
	}

	response.ReqStatus.Success = rp.StatusCode == http.StatusOK

	return response
}
