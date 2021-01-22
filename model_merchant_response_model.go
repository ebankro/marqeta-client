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

type MerchantResponseModel struct {
	Name         string  `json:"name"`
	Active       bool    `json:"active,omitempty"`
	Contact      string  `json:"contact,omitempty"`
	ContactEmail string  `json:"contact_email,omitempty"`
	Longitude    float32 `json:"longitude,omitempty"`
	Latitude     float32 `json:"latitude,omitempty"`
	Address1     string  `json:"address1,omitempty"`
	Address2     string  `json:"address2,omitempty"`
	City         string  `json:"city,omitempty"`
	State        string  `json:"state,omitempty"`
	Province     string  `json:"province,omitempty"`
	Zip          string  `json:"zip,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	Country      string  `json:"country,omitempty"`
	// The unique identifier of the merchant
	Token           string `json:"token,omitempty"`
	PartialAuthFlag bool   `json:"partial_auth_flag,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime *time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime *time.Time `json:"last_modified_time"`
}
