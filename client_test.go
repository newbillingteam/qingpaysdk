package qingpay

import "testing"

func TestNewClient(t *testing.T) {
	c, err := NewClient("SEUODHHPDBGOYDMRG", "MyHPwVxbCIRXEOJ8K2dHkJZjjbsFHqRMH", "ks", false, nil)
	if err != nil {
		t.Errorf("%v", err)
	}
	if c == nil {
		t.Error("service should be valid")
	}
	if c.APIBase() != SandBoxBaseURL {
		t.Errorf(`APIBase should be "%s", but "%s"`, SandBoxBaseURL, c.APIBase())
	}
}
