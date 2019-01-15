package main

import (
	"io/ioutil"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestValidManifest(t *testing.T) {
	specsFile, _ := ioutil.ReadFile("examples/specs.yaml")
	var specs Specs
	yaml.Unmarshal(specsFile, &specs)
	e := specs.validate()
	if e != nil {
		t.Errorf("Valid manifest should not produce errors")
	}
}
