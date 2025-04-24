// Created: 2025-04-22
package did

type Resolver interface {
	ResolveHandle(identity string) (*DID, error)
	ResolveDID(did *DID) (*Document, error)
}
