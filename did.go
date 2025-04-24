// Created: 2025-04-23
package did

import (
	"errors"
	"fmt"
	"strings"
)

// Models a DID.
type DID struct {
	Method string
	Value  string
}

func ParseDID(did string) (*DID, error) {
	split := strings.Split(did, ":")
	if 2 > len(split) {
		return nil, errors.New("")
	}
	return &DID{
		Method: split[1],
		Value:  split[2],
	}, nil
}

func (did *DID) String() string {
	return fmt.Sprintf("did:%s:%s", did.Method, did.Value)
}
