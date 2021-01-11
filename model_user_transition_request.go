/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type UserTransitionRequest struct {
	IdempotentHash string `json:"idempotentHash,omitempty"`
	Token          string `json:"token,omitempty"`
	Status         string `json:"status"`
	ReasonCode     string `json:"reason_code"`
	Reason         string `json:"reason,omitempty"`
	Channel        string `json:"channel"`
	UserToken      string `json:"user_token"`
}
