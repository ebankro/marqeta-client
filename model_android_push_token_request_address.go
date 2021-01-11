/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type AndroidPushTokenRequestAddress struct {
	Name       string `json:"name,omitempty"`
	Address1   string `json:"address1,omitempty"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Zip        string `json:"zip,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country,omitempty"`
	Phone      string `json:"phone,omitempty"`
}
