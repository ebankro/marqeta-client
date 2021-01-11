/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type JitFundingProgramgatewayFundingSource struct {
	Enabled bool `json:"enabled,omitempty"`
	// Required if enabled
	FundingSourceToken string `json:"funding_source_token,omitempty"`
	RefundsDestination string `json:"refunds_destination,omitempty"`
	AlwaysFund         bool   `json:"always_fund,omitempty"`
}
