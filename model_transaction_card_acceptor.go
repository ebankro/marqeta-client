/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type TransactionCardAcceptor struct {
	Mid         string         `json:"mid,omitempty"`
	Mcc         string         `json:"mcc,omitempty"`
	NetworkMid  string         `json:"network_mid,omitempty"`
	MccGroups   []string       `json:"mcc_groups,omitempty"`
	Name        string         `json:"name,omitempty"`
	Address     string         `json:"address,omitempty"`
	City        string         `json:"city,omitempty"`
	State       string         `json:"state,omitempty"`
	PostalCode  string         `json:"postal_code,omitempty"`
	CountryCode string         `json:"country_code,omitempty"`
	Poi         *TerminalModel `json:"poi,omitempty"`
}
