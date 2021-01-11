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

type RealTimeFeeGroup struct {
	// 36 char max
	Token            string    `json:"token"`
	CreatedTime      time.Time `json:"created_time,omitempty"`
	LastModifiedTime time.Time `json:"last_modified_time,omitempty"`
	Active           bool      `json:"active"`
	// 50 char max
	Name      string   `json:"name"`
	FeeTokens []string `json:"fee_tokens,omitempty"`
}
