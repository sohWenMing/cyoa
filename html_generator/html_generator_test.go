package htmlGenerator

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"testing"

	jsonDecoder "github.com/sohWenMing/cyoa/json_decoder"
)

var choices jsonDecoder.StoryMap

func TestMain(m *testing.M) {
	returnedChoices, err := jsonDecoder.GetOptionsFromJSON()
	if err != nil {
		log.Fatal("error when decoding json")
	}
	choices = returnedChoices
	code := m.Run()
	os.Exit(code)
}
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
	introChoice, ok := choices["intro"]
	if !ok {
		t.Errorf("didn't find choice that was tied to key intro")
	}
	got, err := GenerateHTML(introChoice)
	want := "<html><head></head><body><h1>The Little Blue Gopher</h1>Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn&#39;t sure where to go. Will you go on an adventure with him?<br>One of his friends once recommended going to New York to make friends at this mysterious thing called &#34;GothamGo&#34;. It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.<br>On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit.<br></body></html>"
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	if got != want {
		t.Errorf("\ngot%s\nwant%s", got, want)
	}
}

func TestWriteMultipleStoryLineStrings(t *testing.T) {
	newStoryLineTemplate := template.New("storyline")
	template, err := newStoryLineTemplate.Parse("{{.StoryLine}}<br>")
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	buf := &bytes.Buffer{}
	introChoice, ok := choices["intro"]
	if !ok {
		t.Errorf("didn't find choice that was tied to key intro")
	}
	for _, storyLine := range introChoice.Story {
		curStruct := StoryLineStruct{storyLine}
		err := writeStoryLineString(buf, curStruct, template)
		if err != nil {
			t.Errorf("didn't expect error, got %v", err)
		}
	}
}
