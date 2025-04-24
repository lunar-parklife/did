// Created: 2025-04-24
package did

import (
	"testing"
)

func TestDIDParses(t *testing.T) {
	raw := "did:plc:vwzwgnygau7ed7b7wt5ux7y2"
	did, err := ParseDID(raw)
	if err != nil {
		t.Error(err)
	}
	round := did.String()
	if round != raw {
		t.Errorf("Error round-tripping DID! Had %s, wanted %s", raw, round)
	}
}
