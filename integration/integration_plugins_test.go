//go:build integration
// +build integration

package main

import (
	"testing"

	"github.com/autobrr/go-deluge"
)

func enablePlugin(t *testing.T, name string) {
	err := deluge.EnablePlugin(name)
	if err != nil {
		t.Fatal(err)
	}

	printServerResponse(t, "EnablePlugin")
}

func disablePlugin(t *testing.T, name string) {
	err := deluge.DisablePlugin(name)
	if err != nil {
		t.Fatal(err)
	}

	printServerResponse(t, "DisablePlugin")
}

func TestLabelPluginGetLabels(t *testing.T) {
	enablePlugin(t, "Label")
	defer disablePlugin(t, "Label")

	var labelPlugin = &deluge.LabelPlugin{Client: c}

	_, err := labelPlugin.GetLabels()
	if err != nil {
		t.Fatal(err)
	}

	printServerResponse(t, "LabelPlugin.GetLabels")
}
