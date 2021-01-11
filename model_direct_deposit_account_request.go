/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type DirectDepositAccountRequest struct {
	// Required if 'business_token' is null
	UserToken string `json:"user_token,omitempty"`
	// Required if 'user_token' is null
	BusinessToken        string `json:"business_token,omitempty"`
	Type_                string `json:"type,omitempty"`
	AllowImmediateCredit bool   `json:"allow_immediate_credit,omitempty"`
	Token                string `json:"token,omitempty"`
	// Required if account type = Checking
	CustomerDueDiligence []CustomerDueDiligenceRequest `json:"customer_due_diligence,omitempty"`
}
