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

type InternalFundingSource struct {
	Name             string           `json:"name,omitempty"`
	Active           bool             `json:"active,omitempty"`
	Account          string           `json:"account,omitempty"`
	Type_            string           `json:"type,omitempty"`
	Id               string           `json:"id,omitempty"`
	Token            string           `json:"token"`
	CreatedTime      time.Time        `json:"created_time,omitempty"`
	LastModifiedTime time.Time        `json:"last_modified_time,omitempty"`
	DebitAccount     *InternalAccount `json:"debit_account,omitempty"`
	IsDefaultAccount bool             `json:"is_default_account,omitempty"`
}
