package qingpaysdk

import "testing"

var (
	testUserName =  "ks"
	testAccessKey = "SEUODHHPDBGOYDMRG"
	testAccessSecret = "MyHPwVxbCIRXEOJ8K2dHkJZjjbsFHqRMH"
	orderNo = "b12334542335929"
)

var TradePagePayJSON = []byte(`
{
	"content": {
		"order_no": "202012011606792939678847",
		"out_trade_no": "123345435435431232156",
		"pay_channel": "ALIPAY",
		"pay_url": "https://openapi.alipaydev.com/gateway.do?app_id=2016103100781998\u0026biz_content=%7B%22subject%22%3A%22%E4%BA%91%E6%9C%8D%E5%8A%A1%E5%99%A8%22%2C%22out_trade_no%22%3A%22123345435435431232156%22%2C%22total_amount%22%3A%220.01%22%2C%22product_code%22%3A%22FAST_INSTANT_TRADE_PAY%22%2C%22passback_params%22%3A%22username%253dks%22%2C%22time_expire%22%3A%222020-12-01+11%3A37%3A19%22%7D\u0026charset=utf-8\u0026format=JSON\u0026method=alipay.trade.page.pay\u0026notify_url=http%3A%2F%2F139.198.17.61%3A9400%2Fv1%2Fnotify%2Falipay\u0026return_url=http%3A%2F%2F139.198.17.61%3A9400%2Fv1%2Freturn%2Falipay%2Fa3M%3D%3Fout_trade_no%3D123345435435431232156\u0026sign=EL0TW%2FivcqfL%2Bu%2B8fy0lVOh2ARs3N71UHYZU%2FTQNGH5SEjCew2nFRuswzEU0haEBw7OfcfRwNTXsqofdQCcdEfQ%2BWw%2FpaV4DEC70Zv%2F75FQhvSY9HBumwD%2B1M2Mw5YmF615b0iP5hpjD0yXPC5IRWzqXFf9BfSiHNwN0JjbVjPTVOIee7mjuFHybP3O%2FCp9NmMjCDgbdTsbjAzhRuZuUM9y4pI1%2BaQvPaCoBdu9ASkR0HOkj1Lw6%2BIaXO0Oyo9Nt8e9YsRnfK%2BOjKvOlgq1wISMqmBVD2gEU5jqFsP4dqOvwRtnHA0%2F8Gy5TFk%2Fc%2FtVJgj205LhaVaMA3zWCx4sn6g%3D%3D\u0026sign_type=RSA2\u0026timestamp=2020-12-01+11%3A22%3A19\u0026version=1.0",
		"expire_time": "2020-12-01 11:37:19"
	},
	"return_code": "SUCCESS",
	"return_msg": "OK",
	"sign": ""
}
`)

func TestTradePay(t *testing.T) {
	//mock, transport := NewMockClient(200, TradePagePayJSON)
	c, err := NewClient(testAccessKey, testAccessSecret, testUserName, false, nil)
	if err != nil {
		t.Errorf("err should be nil, but %s", err)
	}

	requestData := TradePayRequest{
		Username:    "ks",
		OutTradeNo:  orderNo,
		PayChannel:  "ALIPAY",
		PayAmount:   1,
		ProductName: "apples",
		Method:      PayMethodWeb,
		NotifyUrl:   "http://129.211.58.64:8010/callback/alipay",
	}
	trade, err := c.TradePay(requestData)
	//if transport.URL != "http://127.0.0.1:9400/v1/trade/pay" {
	//	t.Errorf("URL is wrong: %s", transport.URL)
	//}
	//if transport.Method != "POST" {
	//	t.Errorf("Method should be Get, but %s", transport.Method)
	//}
	if err != nil {
		t.Errorf("err should be nil, but %v", err)
	} else if trade == nil {
		t.Error("trade should not be nil")
	} else if trade.Content.OutTradeNo != orderNo {
		t.Errorf("OutTradeNo should be %s, but %s", orderNo, trade.Content.OutTradeNo)
	} else if trade.ReturnCode != "SUCCESS" {
		t.Errorf("ResultInfo Code should be SUCCESS, but %s", trade.ReturnCode)
	}
}

func TestTradeQuery(t *testing.T) {
	c, err := NewClient(testAccessKey, testAccessSecret, testUserName, false, nil)
	if err != nil {
		t.Errorf("err should be nil, but %s", err)
	}

	requestData := TradeQueryRequest{
		Username:   "ks",
		OutTradeNo: "a123345423325196",
		TradeNo:    "",
	}

	order, err := c.TradeQuery(requestData)
	if err != nil {
		t.Errorf("err should be nil, but %v", err)
	} else if order == nil {
		t.Error("trade should not be nil")
	} else if order.ReturnCode != "SUCCESS" {
		t.Errorf("ResultInfo Code should be SUCCESS, but %s", order.ReturnCode)
	} else if order.Content.OutTradeNo != orderNo {
		t.Errorf("OutTradeNo should be %s, but %s", orderNo, order.Content.OutTradeNo)
	}
}

func TestTradeRefund(t *testing.T) {
	c, err := NewClient(testAccessKey, testAccessSecret, testUserName, false, nil)
	if err != nil {
		t.Errorf("err should be nil, but %s", err)
	}

	requestData := TradeRefundRequest{
		Username:     "ks",
		OutTradeNo:   "1231299252378",
		TradeNo:      "",
		RefundAmount: 1,
	}

	order, err := c.TradeRefund(requestData)
	if err != nil {
		t.Errorf("err should be nil, but %v", err)
	} else if order == nil {
		t.Error("trade should not be nil")
	} else if order.ReturnCode != "SUCCESS" {
		t.Errorf("ResultInfo Code should be SUCCESS, but %s", order.ReturnCode)
	} else if order.Content.OutTradeNo != orderNo {
		t.Errorf("OutTradeNo should be %s, but %s", orderNo, order.Content.OutTradeNo)
	}
}
