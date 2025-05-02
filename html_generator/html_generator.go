package htmlGenerator

import (
	"bytes"
	"fmt"
	"html/template"

	jsonDecoder "github.com/sohWenMing/cyoa/json_decoder"
)

var headerTemplate string = `{{define "header"}} `
var titleTemplate *template.Template = template.New("title")
var storylineTemplate *template.Template = template.New("storyline")

type TitleStruct struct {
	Title string
}

type StoryLineStruct struct {
	StoryLine string
}

func GenerateHTML(storyChoice jsonDecoder.StoryChoice) (html string, err error) {

	s, err := storylineTemplate.Parse("{{.StoryLine}}<br>")
	titleStruct := TitleStruct{
		storyChoice.Title,
	}

	buf := &bytes.Buffer{}
	buf.WriteString("<html>")
	buf.WriteString("<head>")
	buf.WriteString("</head>")
	buf.WriteString("<body>")
	err = writeTitleString(buf, titleStruct, titleTemplate)
	if err != nil {
		return "", err
	}
	err = writeStoryLines(buf, storyChoice, s)
	if err != nil {
		return "", err
	}

	buf.WriteString("</body>")
	buf.WriteString("</html>")
	return buf.String(), nil
}

func writeStoryLines(buf *bytes.Buffer, storyChoice jsonDecoder.StoryChoice, storyLineTemplate *template.Template) (err error) {
	for _, storyLine := range storyChoice.Story {
		storyLineStruct := StoryLineStruct{storyLine}
		err = writeStoryLineString(buf, storyLineStruct, storyLineTemplate)
		if err != nil {
			return err
		}
	}
	return nil

}

func writeStoryLineString(buf *bytes.Buffer, storyLineStruct StoryLineStruct, storyLineTemplate *template.Template) (err error) {
	err = storyLineTemplate.Execute(buf, storyLineStruct)
	if err != nil {
		return err
	}
	return nil
}

func writeTitleString(buf *bytes.Buffer, titleStruct TitleStruct, titleTemplate *template.Template) (err error) {
	t, err := titleTemplate.Parse("<h1>{{.Title}}</h1>")
	if err != nil {
		return err
	}
	err = t.Execute(buf, titleStruct)
	if err != nil {
		return err
	}
	return nil
}
func mapTemplate(templateName, templateString string) (returnedTemplateString string) {
	return fmt.Sprintf("{{define \"%s\"}}%s{{end}}", templateName, templateString)
}

// type Inventory struct {
// 	Material string
// 	Count    uint
// }
// sweaters := Inventory{"wool", 17}
// tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
// if err != nil { panic(err) }
// err = tmpl.Execute(os.Stdout, sweaters)
// if err != nil { panic(err) }

// func initTemplate(templateName string) (returnedTemplate *template.Template, err error) {
// 	newTemplate := template.New(templateName)
// 	t, err := newTemplate.Parse(`{{define "T"}} <b>This is a template to hold {{.}}!</b>{{end}}`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return t, nil
// }

//   "intro": {
//     "title": "The Little Blue Gopher",
//     "story": [
//       "Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?",
//       "One of his friends once recommended going to New York to make friends at this mysterious thing called \"GothamGo\". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.",
//       "On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit."
//     ],
//     "options": [
//       {
//         "text": "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.",
//         "arc": "new-york"
//       },
//       {
//         "text": "Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.",
//         "arc": "denver"
//       }
//     ]
//   },
