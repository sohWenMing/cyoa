package htmlGenerator

import (
	"fmt"
	"testing"

	jsonDecoder "github.com/sohWenMing/cyoa/json_decoder"
)

func TestMapTemplate(t *testing.T) {
	type templateNameToString struct {
		name           string
		templateString string
	}

	type test struct {
		testName   string
		testValues templateNameToString
		want       string
	}

	tests := []test{
		{
			"basic test",
			templateNameToString{
				"test1",
				"<p>hello world{{.}}",
			},
			"{{define \"test1\"}}<p>hello world{{.}}{{end}}",
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			got := mapTemplate(test.testValues.name, test.testValues.templateString)
			if got != test.want {
				t.Errorf("\ngot %s\nwant %s\n", got, test.want)
			}
		})
	}

}

func TestGenerateHTML(t *testing.T) {
	choices, err := jsonDecoder.GetOptionsFromJSON()
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	introChoice, ok := choices["intro"]
	if !ok {
		t.Errorf("didn't find choice that was tied to key intro")
	}
	got, err := GenerateHTML(introChoice)
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	fmt.Println("got", got)
}
