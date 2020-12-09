package foo

import "testing"

var validPath = "valid.yaml"

func TestValidOpenAPI(t *testing.T) {
	if valid, _ := validateOpenAPI(validPath); !valid {
		t.Error("Valid OpenAPI file failed validation")
	}
}
