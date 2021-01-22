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

type BankTransferTransitionResponseModel struct {
	Token               string     `json:"token,omitempty"`
	BankTransferToken   string     `json:"bank_transfer_token"`
	Status              string     `json:"status"`
	Reason              string     `json:"reason,omitempty"`
	Channel             string     `json:"channel"`
	BatchNumber         string     `json:"batch_number,omitempty"`
	ProgramReserveToken string     `json:"program_reserve_token,omitempty"`
	TransactionToken    string     `json:"transaction_token,omitempty"`
	CreatedTime         *time.Time `json:"created_time,omitempty"`
	LastModifiedTime    *time.Time `json:"last_modified_time,omitempty"`
}
