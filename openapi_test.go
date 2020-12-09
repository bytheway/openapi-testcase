package foo

import (
	"log"
	"testing"
)

var validPath = "valid.yaml"

func TestValidOpenAPI(t *testing.T) {
	if valid, _ := validateOpenAPI(validPath); !valid {
		t.Error("Valid OpenAPI file failed validation")
	}
}

func TestMungedValidOpenAPI(t *testing.T) {
	munged, deferFunc, err := processOpenAPI(validPath)
	defer deferFunc()

	if err != nil {
		log.Fatalf("Error proecessing openapi.yaml: %s", err)
	}
	if valid, _ := validateOpenAPI(munged); !valid {
		t.Error("Valid OpenAPI file failed validation")
	}
}
