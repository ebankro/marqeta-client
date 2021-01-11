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

type AutoReloadResponseModel struct {
	Token  string `json:"token,omitempty"`
	Active bool   `json:"active,omitempty"`
	// Required when order scope is GPA
	FundingSourceToken        string                 `json:"funding_source_token,omitempty"`
	FundingSourceAddressToken string                 `json:"funding_source_address_token,omitempty"`
	Association               *AutoReloadAssociation `json:"association,omitempty"`
	// either GPA or MSA is required
	OrderScope   *OrderScope `json:"order_scope"`
	CurrencyCode string      `json:"currency_code"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
}
