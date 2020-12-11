package qingpaysdk

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentmethod"
)

func InitStripe(secretKey string) {
	stripe.Key = secretKey
}

func CreatePaymentMethod(cardNo, expYear, expMonth, cvc, addressLine1, addressZip string) (string, error) {
	params := &stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String(cardNo),
			ExpMonth: stripe.String(expMonth),
			ExpYear:  stripe.String(expYear),
			CVC:      stripe.String(cvc),
		},
		Type: stripe.String("card"),
	}

	if addressLine1 != "" && addressZip != "" {
		params.BillingDetails = &stripe.BillingDetailsParams{
			Address: &stripe.AddressParams{
				Line1:      stripe.String(addressLine1),
				PostalCode: stripe.String(addressZip),
			},
		}
	}

	pm, err := paymentmethod.New(params)
	if err != nil {
		return "", err
	}
	return pm.ID, nil
}
