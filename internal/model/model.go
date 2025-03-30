package model

type ValidatorRequest struct {
	NumValidators int    `json:"num_validators"`
	FeeRecipient  string `json:"fee_recipient"`
}

type ValidatorResponse struct {
	RequestID string   `db:"request_id" json:"request_id"`
	Status    string   `db:"status" json:"status"`
	Keys      []string `json:"keys,omitempty"`
}
