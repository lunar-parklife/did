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

var defaultHost *url.URL

func init() {
	url, err := url.Parse("https://plc.directory")
	if err != nil {
		panic("wtf")
	}
	defaultHost = url
}

type resolver struct {
	cache   netutil.Cache
	plcHost *url.URL
}

func DefaultResolver() did.Resolver {
	return &resolver{
		plcHost: defaultHost,
	}
}

func (resolver *resolver) ResolveHandle(handle string) (*did.DID, error) {
	return netutil.ResolveIdentity(handle)
}

func (resolver *resolver) ResolveDID(targetDid *did.DID) (*did.Document, error) {
	target := fmt.Sprintf("%s/%s", resolver.plcHost.String(), targetDid.String())
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
