package stripe

import "encoding/json"

const (
	RefundReasonDuplicate           string = "duplicate"
	RefundReasonFraudulent          string = "fraudulent"
	RefundReasonRequestedByCustomer string = "requested_by_customer"

	RefundStatusCanceled  string = "canceled"
	RefundStatusFailed    string = "failed"
	RefundStatusPending   string = "pending"
	RefundStatusSucceeded string = "succeeded"
)

// RefundParams is the set of parameters that can be used when refunding a charge.
// For more details see https://stripe.com/docs/api#refund.
type RefundParams struct {
	Params               `form:"*"`
	Amount               *int64  `form:"amount"`
	Charge               *string `form:"charge"`
	Reason               *string `form:"reason"`
	RefundApplicationFee *bool   `form:"refund_application_fee"`
	ReverseTransfer      *bool   `form:"reverse_transfer"`
}

// RefundListParams is the set of parameters that can be used when listing refunds.
// For more details see https://stripe.com/docs/api#list_refunds.
type RefundListParams struct {
	ListParams `form:"*"`
}

// Refund is the resource representing a Stripe refund.
// For more details see https://stripe.com/docs/api#refunds.
type Refund struct {
	Amount             int64               `json:"amount"`
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	Charge             *Charge             `json:"charge"`
	Created            int64               `json:"created"`
	Currency           Currency            `json:"currency"`
	ID                 string              `json:"id"`
	Metadata           map[string]string   `json:"metadata"`
	Reason             string              `json:"reason"`
	ReceiptNumber      string              `json:"receipt_number"`
	Status             string              `json:"status"`
}

// RefundList is a list object for refunds.
type RefundList struct {
	ListMeta
	Data []*Refund `json:"data"`
}

// UnmarshalJSON handles deserialization of a Refund.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *Refund) UnmarshalJSON(data []byte) error {
	type refund Refund
	var rr refund
	err := json.Unmarshal(data, &rr)
	if err == nil {
		*r = Refund(rr)
	} else {
		// the id is surrounded by "\" characters, so strip them
		r.ID = string(data[1 : len(data)-1])
	}

	return nil
}
