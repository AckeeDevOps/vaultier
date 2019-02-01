package main

import (
	"testing"
)

func TestMergeResultsFunction(t *testing.T) {
	input := []map[string]interface{}{}
	input = append(input, map[string]interface{}{
		"var1": "val1",
		"var2": "val2",
	})
	input = append(input, map[string]interface{}{
		"var3": 3000,
	})

	result := mergeResults(input)

	if _, ok := result["var1"]; !ok {
		t.Errorf("result does not contain expected key")
	}

	if _, ok := result["var2"]; !ok {
		t.Errorf("result does not contain expected key")
	}

	if _, ok := result["var3"]; !ok {
		t.Errorf("result does not contain expected key")
	}

}
