package main

import (
	"shortener/src"
	"testing"
)

func TestParseYaml(t *testing.T) {

	text := []byte(
		`- path: name
  url: value`)

	computed, err := src.ParseYaml(text)
	if err != nil {
		t.Error("Got an error", err)
	}

	if computed["name"] != "value" {
		t.Errorf("Expected %v got %v", "value", computed["name"])
	}

}
