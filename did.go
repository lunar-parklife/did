// Created: 2025-04-23
package did

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Models a DID.
type DID struct {
	Method string
	Value  string
}

func (did *DID) MarshalJSON() ([]byte, error) {
	return json.Marshal(did.String())
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

func (did *DID) UnmarshalJSON(buffer []byte) error {
	var str string
	err := json.Unmarshal(buffer, &str)
	if err != nil {
		return err
	}
	sourceDID, err := ParseDID(str)
	if err != nil {
		return err
	}
	did.Method = sourceDID.Method
	did.Value = sourceDID.Value
	return nil
}

func (did *DID) String() string {
	return fmt.Sprintf("did:%s:%s", did.Method, did.Value)
}
