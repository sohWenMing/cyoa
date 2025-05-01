package jsonDecoder

import "testing"

func TestJsonDecoder(t *testing.T) {
	options, err := GetOptionsFromJSON()
	if err != nil {
		t.Errorf("Didn't expect error, got %v", err)
	}
	expectedOptions := []string{
		"intro",
		"new-york",
		"debate",
		"sean-kelly",
		"mark-bates",
		"denver",
		"home",
	}
	for _, expectedOption := range expectedOptions {
		_, ok := options[expectedOption]
		if !ok {
			t.Errorf("could find record for option %s", expectedOption)
		}
	}
}
