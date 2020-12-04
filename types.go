package qingpaysdk

import (
	"fmt"
	"net/http"
)

const (
	SandBoxBaseURL    = "http://139.198.121.68:19400"
	ProductionBaseURL = "http://139.198.121.68:19400"
)

const (
	PayMethodWeb     = "WEB"
	PayMethodWap     = "WAP"
	PayMethodCashier = "CASHIER"
)

type (
	ErrorData struct {
		Code    int32               `json:"code"`
		Message string              `json:"message"`
		Status  string              `json:"status"`
		Details []map[string]string `json:"details"`
	}

	Metadata struct{}

	Link struct {
		Href        string `json:"href"`
		Rel         string `json:"rel,omitempty"`
		Method      string `json:"method,omitempty"`
		Description string `json:"description,omitempty"`
		Enctype     string `json:"enctype,omitempty"`
	}

	ErrorResponseDetail struct {
		Field string `json:"field"`
		Issue string `json:"issue"`
		Links []Link `json:"link"`
	}

	ErrorResponse struct {
		Response        *http.Response        `json:"-"`
		Name            string                `json:"name"`
		DebugID         string                `json:"debug_id"`
		Message         string                `json:"message"`
		InformationLink string                `json:"information_link"`
		Details         []ErrorResponseDetail `json:"details"`
	}
)

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s, %+v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message, r.Details)
}
