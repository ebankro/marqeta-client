/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

import (
	"time"
)

type GpaResponse struct {
	Token  string      `json:"token"`
	Amount float32     `json:"amount"`
	Tags   string      `json:"tags,omitempty"`
	Memo   string      `json:"memo,omitempty"`
	Fees   []FeeDetail `json:"fees,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime *time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime          *time.Time     `json:"last_modified_time"`
	TransactionToken          string         `json:"transaction_token"`
	State                     string         `json:"state"`
	Response                  *Response      `json:"response"`
	Funding                   *Funding       `json:"funding"`
	FundingSourceToken        string         `json:"funding_source_token"`
	FundingSourceAddressToken string         `json:"funding_source_address_token,omitempty"`
	JitFunding                *JitFundingApi `json:"jit_funding,omitempty"`
	UserToken                 string         `json:"user_token,omitempty"`
	BusinessToken             string         `json:"business_token,omitempty"`
	CurrencyCode              string         `json:"currency_code"`
	GatewayToken              int64          `json:"gateway_token,omitempty"`
	GatewayMessage            string         `json:"gateway_message,omitempty"`
}
