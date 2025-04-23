// Created: 2025-04-22
package did

type Resolver interface {
	ResolveHandle(identity string) (string, error)
	ResolveDID(did string) (*Document, error)
}
