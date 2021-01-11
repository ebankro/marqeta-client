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

type CardholderAddressResponse struct {
	// Required if 'business_token' is not specified
	UserToken string `json:"user_token,omitempty"`
	// Required if 'user_token' is not specified
	BusinessToken    string `json:"business_token,omitempty"`
	Token            string `json:"token"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Address1         string `json:"address_1"`
	Address2         string `json:"address_2,omitempty"`
	City             string `json:"city"`
	State            string `json:"state"`
	Zip              string `json:"zip"`
	PostalCode       string `json:"postal_code"`
	Country          string `json:"country"`
	Phone            string `json:"phone,omitempty"`
	IsDefaultAddress bool   `json:"is_default_address,omitempty"`
	Active           bool   `json:"active,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
}
