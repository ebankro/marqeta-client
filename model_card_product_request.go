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

type CardProductRequest struct {
	Token     string             `json:"token,omitempty"`
	Name      string             `json:"name"`
	Active    bool               `json:"active,omitempty"`
	StartDate time.Time          `json:"start_date"`
	EndDate   time.Time          `json:"end_date,omitempty"`
	Config    *CardProductConfig `json:"config,omitempty"`
}
