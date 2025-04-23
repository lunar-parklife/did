// Created: 2025-04-22
package did

import "net/url"

// A service used by an identity.
type Service struct {
	ID              string   `json:"id"`              // I'm unsure what this is.
	Type            string   `json:"type"`            // The type of endpoint this is.
	ServiceEndpoint *url.URL `json:"serviceEndpoint"` // The URL of the endpoint.
}
