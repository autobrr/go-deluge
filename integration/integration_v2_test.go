//go:build integration && integration_v2 && !integration_v1
// +build integration,integration_v2,!integration_v1

package main

import (
	"testing"

	"github.com/autobrr/go-deluge"
)

var (
	v2daemon = true
)

func TestKnownAccounts(t *testing.T) {
	if !v2daemon {
		t.Skip()
		return
	}

	d := deluge.(deluge.V2)
	_, err := deluge.KnownAccounts()
	if err != nil {
		t.Fatal(err)
	}
	printServerResponse(t, "KnownAccounts")
}
