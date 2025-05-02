package testTemplate

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	template, err := initTemplate("testing")
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	buf := &bytes.Buffer{}
	err = template.ExecuteTemplate(buf, "T", "Test value")
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	returnedString := buf.String()
	fmt.Println("returnedString: ", returnedString)
}
