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

type OfferResponseModel struct {
	Token               string    `json:"token,omitempty"`
	Active              bool      `json:"active,omitempty"`
	Name                string    `json:"name"`
	StartDate           time.Time `json:"start_date,omitempty"`
	EndDate             time.Time `json:"end_date,omitempty"`
	PurchaseAmount      float32   `json:"purchase_amount"`
	RewardAmount        float32   `json:"reward_amount"`
	RewardTriggerAmount float32   `json:"reward_trigger_amount,omitempty"`
	CampaignToken       string    `json:"campaign_token"`
	CurrencyCode        string    `json:"currency_code"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
}
