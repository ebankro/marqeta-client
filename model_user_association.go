/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type UserAssociation struct {
	SingleInventoryUser      bool   `json:"single_inventory_user,omitempty"`
	SingleInventoryUserToken string `json:"single_inventory_user_token,omitempty"`
}
