/*
 * Marqeta Core API
 *
 * Simplified management of your payment programs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package marqeta

type Images struct {
	Card                *ImagesCard                `json:"card,omitempty"`
	Carrier             *ImagesCarrier             `json:"carrier,omitempty"`
	Signature           *ImagesSignature           `json:"signature,omitempty"`
	CarrierReturnWindow *ImagesCarrierReturnWindow `json:"carrier_return_window,omitempty"`
}
