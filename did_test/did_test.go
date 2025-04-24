// Created: 2025-04-24
package did_test

import (
	"testing"

	"github.com/lunar-parklife/did"
	"github.com/lunar-parklife/did/plc"
)

var testDID string = "did:plc:vwzwgnygau7ed7b7wt5ux7y2"

func TestDIDParses(t *testing.T) {
	did, err := did.ParseDID(testDID)
	if err != nil {
		t.Error(err)
	}
	round := did.String()
	if round != testDID {
		t.Errorf("Error round-tripping DID! Had %s, wanted %s", testDID, round)
	}
}

func TestDIDMarshall(t *testing.T) {
	resolver := plc.DefaultResolver()
	did, err := did.ParseDID(testDID)
	if err != nil {
		t.Errorf("cannot parse test did! %s", err)
	}
	document, err := resolver.ResolveDID(did)
	if err != nil {
		t.Error(err)
	}
	docs := []did.Document{document}
}
