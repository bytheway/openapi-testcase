package foo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

func validateOpenAPI(path string) (bool, *validate.Result) {

	doc, err := loads.Spec(path)
	if err != nil {
		log.Fatalf("Error loading %s: %s", path, err)
	}

	validator := validate.NewSpecValidator(doc.Schema(), strfmt.Default)
	validator.SetContinueOnErrors(true)  // Set option for this validator
	result, _ := validator.Validate(doc) // Validates spec with default Swagger 2.0 format definitions

	if result.IsValid() {
		fmt.Printf("%s is valid\n", path)
		return true, result
	}

	if result.HasErrors() {
		fmt.Printf("%s has some validation errors:\n", path)
		for _, e := range result.Errors {
			fmt.Printf("\t%s\n", e)
		}
		return false, result
	}
	if result.HasWarnings() {
		fmt.Printf("%shas some validation warnings:\n", path)
		for _, w := range result.Warnings {
			fmt.Printf("\t%s\n", w)
		}
		return true, result
	}

	return false, result
}

func processOpenAPI(path string) (string, func(), error) {
	var spec []byte
	var err error

	if _, err = os.Stat(path); err != nil {
		return "", func() {}, err
	}

	if spec, err = ioutil.ReadFile(path); err != nil {
		return "", func() {}, err
	}

	munged := string(spec)

	// template variable replacment goes here

	mungedFile, err := ioutil.TempFile("", "openapi")
	if err != nil {
		return "", func() {}, err
	}

	mungedFile.WriteString(munged)

	if err := mungedFile.Close(); err != nil {
		log.Fatal(err)
	}

	return mungedFile.Name(), func() { os.Remove(mungedFile.Name()) }, nil
}
