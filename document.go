// Created: 2025-04-22
package did

// A full identity on the network.
type Document struct {
	AlsoKnownAs        []string             `json:"alsoKnownAs,omitempty"` // Other identities this one is known by.
	ID                 string               `json:"id"`                    // The main DID for this identity.
	Service            []Service            `json:"service,omitempty"`     // The services this identity uses.
	VerificationMethod []VerificationMethod `json:"verificationMethod,omitempty"`
}
