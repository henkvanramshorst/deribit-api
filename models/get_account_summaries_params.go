package models

type GetAccountSummariesParams struct {
	SubAccountId int  `json:"subaccount_id,omitempty"`
	Extended     bool `json:"extended,omitempty"`
}
