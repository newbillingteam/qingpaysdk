# QingPay SDK for Go

The official QingPay SDK for the Go programming language.

## Getting Started

### Installation

```
import "github.com/newbillingteam/qingpaysdk"

```

### Preparation

Before your start, please go to get a pair of QingPay API AccessKey.

___API AccessKey Example:___

``` yaml
access_key: 'ACCESS_KEY_ID_EXAMPLE'
secret_access_key: 'SECRET_ACCESS_KEY_EXAMPLE'
```

### Usage

Now you are ready to code. You can read the test method in the code to have a clear understanding.
```
package main

import "github.com/newbillingteam/qingpaysdk"

func main()  {
    c, err := NewClient("access_key", "secret_access_key", "name", false, nil)
    if err != nil {
		t.Errorf("err should be nil, but %s", err)
	}

	requestData := TradePayRequest{
		Username:    "NAME",
		OutTradeNo:  "123456789",
		PayChannel:  "ALIPAY",
		PayAmount:   1,
		ProductName: "apple",
		Method:      "WEB",
	}
	trade, err := c.TradePay(requestData)
    if err != nil {
		t.Errorf("err should be nil, but %v", err)
	} 
    ......
    
}

```

## LICENSE

The Apache License (Version 2.0, 2020).