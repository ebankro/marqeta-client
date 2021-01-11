/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type FeeTransferRequest struct {
	Tags  string     `json:"tags,omitempty"`
	Fees  []FeeModel `json:"fees,omitempty"`
	Token string     `json:"token,omitempty"`
	// Required if 'business_token' is null
	UserToken string `json:"user_token"`
	// Required if 'user_token' is null
	BusinessToken string `json:"business_token"`
}
