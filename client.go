package qingpaysdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/yunify/qingcloud-sdk-go/request"
	"github.com/yunify/qingcloud-sdk-go/utils"
)

type Client struct {
	signer   request.Signer
	apiBase  string
	openName string
	Client   *http.Client
}

func NewClient(key, secret, name string, isProduction bool, client *http.Client) (*Client, error) {
	if key == "" || secret == "" {
		return nil, errors.New("APIKey, APISecret and APIBase are required to create a Client")
	}
	if client == nil {
		client = http.DefaultClient
	}
	signer := request.Signer{
		AccessKeyID:     key,
		SecretAccessKey: secret,
	}

	c := &Client{
		signer:   signer,
		openName: name,
		Client:   client,
	}
	if isProduction {
		c.apiBase = ProductionBaseURL
	} else {
		c.apiBase = SandBoxBaseURL
	}
	return c, nil
}

func (c *Client) APIBase() string {
	return c.apiBase
}

func (c *Client) header(isJson bool) map[string]string {
	contentType := "application/x-www-form-urlencoded;charset=utf-8"
	if isJson {
		contentType = "application/json;charset=utf-8"
	}
	return map[string]string{
		"Content-Type": contentType,
		"open-name":    c.openName,
		"expires":      time.Now().In(time.UTC).Format("2006-01-02T15:04:05Z"),
		"Date":         utils.TimeToString(time.Now(), "RFC 822"),
	}
}

func (c *Client) sign(method, path string, urlValues url.Values) (url.Values, error) {
	req, err := http.NewRequest(method, path, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return nil, err
	}
	if req.Form == nil {
		req.Form = urlValues
	}
	sign, err := c.signer.BuildSignature(req)
	if err != nil {
		fmt.Printf("open create access signature error %+v", err)
		return nil, err
	}
	urlValues.Set("signature", sign)
	return urlValues, nil
}

func (c *Client) doRequest(method, path string, query map[string]string, param interface{}) (body []byte, err error) {
	data, err := json.Marshal(param)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}

	baseURL, err := url.Parse(c.apiBase)
	if err != nil {
		return nil, err
	}
	apiURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	endpoint := baseURL.ResolveReference(apiURL).String()

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range c.header(true) {
		req.Header.Add(key, value)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) doFormRequest(method, path string, query map[string]string, param interface{}) (body []byte, err error) {
	baseURL, err := url.Parse(c.apiBase)
	if err != nil {
		return nil, err
	}
	apiURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	urlValues := ToUrlValues(structToMap(param))

	urlValues, err = c.sign(method, endpoint, urlValues)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, endpoint, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range c.header(false) {
		req.Header.Add(key, value)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}