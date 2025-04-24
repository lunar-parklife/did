// Created: 2025-04-23
package plc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/lunar-parklife/did"
	netutil "github.com/lunar-parklife/did/internal/net"
)

type Resolver struct {
	PLCHost *url.URL
}

func (resolver *Resolver) ResolveHandle(handle string) (*did.DID, error) {
	return netutil.ResolveIdentity(handle)
}

func (resolver *Resolver) ResolveDID(targetDid *did.DID) (*did.Document, error) {
	target := fmt.Sprintf("%s/%s", resolver.PLCHost.String(), targetDid.String())
	response, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	switch response.StatusCode {
	case 200:
		bytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		document := &did.Document{}
		err = json.Unmarshal(bytes, document)
		if err != nil {
			return nil, err
		}
		return document, nil
	default:
		return nil, fmt.Errorf("unknown plc status code %d", response.StatusCode)
	}
}
