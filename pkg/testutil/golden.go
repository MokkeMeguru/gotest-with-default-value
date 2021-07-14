package testutil

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func GetGoldenJSON(t *testing.T, golden string, dst interface{}) {
	g, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Fatalf("failed reading .golden: %s", err)
	}

	if err := json.Unmarshal(g, dst); err != nil {
		t.Errorf("failed to parse golden file: %s, err: %v", golden, err)
		return
	}
}
