/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type DigitalWalletApplePayProvisionRequest struct {
	CardToken              string   `json:"card_token"`
	DeviceType             string   `json:"device_type"`
	ProvisioningAppVersion string   `json:"provisioning_app_version"`
	Certificates           []string `json:"certificates"`
	Nonce                  string   `json:"nonce"`
	NonceSignature         string   `json:"nonce_signature"`
}
