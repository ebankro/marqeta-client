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

type ChargebackFundingSourceModel struct {
	Token            string `json:"token"`
	Active           bool   `json:"active"`
	IsDefaultAccount bool   `json:"is_default_account"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
	Type_            string    `json:"type"`
	Name             string    `json:"name"`
	Credit           bool      `json:"credit"`
}
