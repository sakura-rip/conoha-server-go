package conoha_server_go

import (
	"github.com/imroc/req"
	"golang.org/x/xerrors"
)

type GetPaymentSummaryResponse struct {
	PaymentSummary PaymentSummary `json:"payment_summary"`
}
type PaymentSummary struct {
	TotalDepositAmount int `json:"total_deposit_amount"`
}

func (c *Conoha) GetPaymentSummary() (int, error) {
	r, err := req.Get(c.endPoint.ToUrl(AccountService, "payment-summary"), c.header)
	if err != nil {
		return -1, err
	}
	if r.Response().StatusCode != 200 {
		return 0, xerrors.Errorf("wrong status code: %v, message: %v", r.Response().StatusCode, r.String())
	}
	res := GetPaymentSummaryResponse{}
	err = r.ToJSON(&res)
	if err != nil {
		return -1, err
	}
	return res.PaymentSummary.TotalDepositAmount, nil
}
