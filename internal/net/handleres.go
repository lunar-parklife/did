// Created: 2025-04-23
package net

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/lunar-parklife/did"
)

func ResolveIdentity(identity string) (*did.DID, error) {
	resolver := net.DefaultResolver
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()
	entries, err := resolver.LookupTXT(ctx, "")
	if err == nil {
		for _, entry := range entries {
			if entry[0] == '"' {
				entry = entry[1:]
			}
			if entry[len(entry)] == '"' {
				entry = entry[:len(entry)-1]
			}
			split := strings.Split(entry, "=")
			if 2 > len(split) {
				continue
			}
			if split[0] == "_atproto" {
				return did.ParseDID(split[1])
			}
		}
	} else {
		if errors.Is(err, &net.DNSError{}) {
			response, err := http.Get(fmt.Sprintf("%s/.well-known/atproto", identity))
			if err != nil {
				return nil, err
			} else if response.StatusCode != 200 {
				return nil, fmt.Errorf("cannot find did for handle %s", identity)
			}
			bytes, err := io.ReadAll(response.Body)
			if err != nil {
				response.Body.Close()
				return nil, err
			}
			response.Body.Close()
			return did.ParseDID(string(bytes))
		} else {
			return nil, err
		}
	}
	return nil, nil
}
