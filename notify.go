package qingpaysdk

import (
	"errors"
	"net/http"
)

type TradeNotification struct {
	Username   string `json:"username,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	PayChannel string `json:"pay_channel,omitempty"`
	Amount     string `json:"amount,omitempty"`
	Currency   string `json:"currency,omitempty"`
	OrderNo    string `json:"order_no,omitempty"`
	Passback   string `json:"passback,omitempty"`
	Status     string `json:"status,omitempty"`
}

func (c *Client) GetTradeNotification(req *http.Request) (noti *TradeNotification, err error) {
	return GetTradeNotification(req)
}

func GetTradeNotification(req *http.Request) (noti *TradeNotification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	noti = &TradeNotification{}
	noti.Username = req.FormValue("username")
	noti.OutTradeNo = req.FormValue("out_trade_no")
	noti.PayChannel = req.FormValue("pay_channel")
	noti.Amount = req.FormValue("amount")
	noti.Currency = req.FormValue("currency")
	noti.OrderNo = req.FormValue("order_no")
	noti.Passback = req.FormValue("passback")
	noti.Status = req.FormValue("status")

	return noti, nil
}

func ( *Client) AckNotification(w http.ResponseWriter) {
	AckNotification(w)
}

func AckNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}