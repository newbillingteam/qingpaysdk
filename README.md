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

```go
package main

import (
	"fmt"
	"github.com/newbillingteam/qingpaysdk"
)

func main()  {
	// 第一个参数表示 加密key
	// 第二个参数表示 加密私钥
	// 第三个参数表示 接入方 名称
	// 第四个参数表示 接口环境 true表生产环境 false 表沙箱环境
	c, err := qingpaysdk.NewClient("access_key", "secret_access_key", "name", false, nil)
	if err != nil {
		fmt.Printf("%v", err)
	}
	requestData := qingpaysdk.TradePayRequest{
		Username:    "NAME",
		OutTradeNo:  "123456789",
		PayChannel:  "ALIPAY",
		PayAmount:   1,
		ProductName: "apple",
		Method:      "WEB",
		NotifyUrl:   "http://www.xxxx.com/callback",
	}
	trade, err := c.TradePay(requestData)
	if err == nil && trade.ReturnCode == "SUCCESS"{
		fmt.Printf("%v", trade.Content.PayUrl)
	}

}

```



## LICENSE

The Apache License (Version 2.0, 2020).