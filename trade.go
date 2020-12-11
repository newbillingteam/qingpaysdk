package qingpaysdk

import (
	"fmt"
	"log"

	"encoding/json"
)

type (
	TradePayRequest struct {
		Username        string `json:"username,omitempty"`
		OutTradeNo      string `json:"out_trade_no,omitempty"`
		PayChannel      string `json:"pay_channel,omitempty"`
		PayAmount       uint64 `json:"pay_amount,omitempty"`
		Currency        string `json:"currency,omitempty"`
		PaymentMethodId string `json:"payment_method_id,omitempty"`
		ReturnUrl       string `json:"return_url,omitempty"`
		NotifyUrl       string `json:"notify_url,omitempty"`
		OrderTime       string `json:"order_time,omitempty"`
		ProductName     string `json:"product_name,omitempty"`
		Passback        string `json:"passback,omitempty"`
		Method          string `json:"method,omitempty"` //WEB WAP CASHIER
		Remark          string `json:"remark,omitempty"`
		SignType        string `json:"sign_type,omitempty"`
		Sign            string `json:"sign,omitempty"`
	}

	TradePayResponse struct {
		ReturnCode string `json:"return_code"`
		ReturnMsg  string `json:"return_msg"`
		Content    struct {
			OrderNo    string `json:"order_no,omitempty"`
			OutTradeNo string `json:"out_trade_no,omitempty"`
			PayChannel string `json:"pay_channel,omitempty"`
			PayUrl     string `json:"pay_url,omitempty"`
			CashierUrl string `json:"cashier_url,omitempty"`
			PayQrcode  string `json:"pay_qrcode,omitempty"`
			ExpireTime string `json:"expire_time,omitempty"`
			SessionId  string `json:"session_id,omitempty"`
		} `json:"content"`
		Error ErrorData `json:"error,omitempty"`
		Sign  string    `json:"sign,omitempty"`
	}
)

func (c *Client) TradePay(pagePayRequest TradePayRequest) (*TradePayResponse, error) {
	resp, err := c.doFormRequest("POST", "/v1/trade/pay", nil, pagePayRequest)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	fmt.Println(string(resp))

	var pagePageResponse TradePayResponse
	if err := json.Unmarshal(resp, &pagePageResponse); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return &pagePageResponse, nil
}

type (
	TradeQueryRequest struct {
		TradeNo    string `json:"trade_no,omitempty"`
		OutTradeNo string `json:"out_trade_no,omitempty"`
		Username   string `json:"username,omitempty"`
		PayChannel string `json:"pay_channel,omitempty"`
		SignType   string `json:"sign_type,omitempty"`
		Sign       string `json:"sign,omitempty"`
	}

	TradeQueryResponse struct {
		ReturnCode string `json:"return_code"`
		ReturnMsg  string `json:"return_msg"`
		Content    struct {
			OrderNo      string `json:"order_no,omitempty"`
			Username     string `json:"username,omitempty"`
			Subject      string `json:"subject,omitempty"`
			OutTradeNo   string `json:"out_trade_no,omitempty"`
			PayChannel   string `json:"pay_channel,omitempty"`
			Status       string `json:"status,omitempty"`
			Amount       uint64 `json:"amount,omitempty"`
			Currency     string `json:"currency,omitempty"`
			ExpireTime   string `json:"expire_time,omitempty"`
			ReturnUrl    string `json:"return_url,omitempty"`
			NotifyUrl    string `json:"notify_url,omitempty"`
			Remark       string `json:"remark,omitempty"`
			Passback     string `json:"passback,omitempty"`
			IsRefund     bool   `json:"is_refund,omitempty"`
			RefundTimes  int32  `json:"refund_times,omitempty"`
			RefundAmount uint64 `json:"refund_amount,omitempty"`
			PayTime      string `json:"pay_time,omitempty"`
		} `json:"content"`
		Error *ErrorData `json:"error,omitempty"`
		Sign  string     `json:"sign,omitempty"`
	}
)

func (c *Client) TradeQuery(queryRequest TradeQueryRequest) (*TradeQueryResponse, error) {
	resp, err := c.doFormRequest("POST", "/v1/trade/query", nil, queryRequest)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	fmt.Println(string(resp))

	var queryResponse TradeQueryResponse
	if err := json.Unmarshal(resp, &queryResponse); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return &queryResponse, nil
}

type (
	TradeRefundRequest struct {
		TradeNo        string `json:"trade_no,omitempty"`
		OutTradeNo     string `json:"out_trade_no,omitempty"`
		PayChannel     string `json:"pay_channel,omitempty"`
		RefundAmount   uint64 `json:"refund_amount,omitempty"`
		RefundCurrency string `json:"refund_currency,omitempty"`
		RefundReason   string `json:"refund_reason,omitempty"`
		Username       string `json:"username,omitempty"`
		SignType       string `json:"sign_type,omitempty"`
		Sign           string `json:"sign,omitempty"`
	}

	TradeRefundResponse struct {
		ReturnCode string `json:"return_code"`
		ReturnMsg  string `json:"return_msg"`
		Content    struct {
			RefundNo       string `json:"refund_no,omitempty"`
			OutTradeNo     string `json:"out_trade_no,omitempty"`
			RefundAmount   uint64 `json:"refund_amount,omitempty"`
			RefundCurrency string `json:"refund_currency,omitempty"`
			RefundPayTime  string `json:"refund_pay_time,omitempty"`
			Username       string `json:"username,omitempty"`
		} `json:"content"`
		Error ErrorData `json:"error,omitempty"`
		Sign  string    `json:"sign,omitempty"`
	}
)

func (c *Client) TradeRefund(refundRequest TradeRefundRequest) (*TradeRefundResponse, error) {
	resp, err := c.doFormRequest("POST", "/v1/trade/refund", nil, refundRequest)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	fmt.Println(string(resp))

	var refundResponse TradeRefundResponse
	if err := json.Unmarshal(resp, &refundResponse); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return &refundResponse, nil
}

type TradeConfirmRequest struct {
	OutTradeNo      string `json:"out_trade_no,omitempty"`
	PayChannel      string `json:"pay_channel,omitempty"`
	PaymentIntentId string `json:"payment_intent_id,omitempty"`
	Username        string `json:"username,omitempty"`
	SignType        string `json:"sign_type,omitempty"`
	Sign            string `json:"sign,omitempty"`
}

func (c *Client) TradeConfirm(confirmRequest TradeConfirmRequest) (*TradePayResponse, error) {
	resp, err := c.doFormRequest("POST", "/v1/trade/confirm", nil, confirmRequest)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	fmt.Println(string(resp))

	var confirmResponse TradePayResponse
	if err := json.Unmarshal(resp, &confirmResponse); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return &confirmResponse, nil
}
