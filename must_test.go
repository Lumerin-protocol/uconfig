package uconfig_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/omeid/uconfig"
	"github.com/omeid/uconfig/internal/f"
	"github.com/omeid/uconfig/plugins/file"
)

func TestMust(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Must should not panic, but did: %v", r)
		}
	}()

	value := f.Config{}
	uconfig.Must(&value)

}

func TestMustPanic(t *testing.T) {

	defer func() {
		r := recover()

		if r == nil {
			t.Error("Was expecting panic but got nil")
		}

		expectErr := "read testdata/classic.json: file already closed"

		if err, ok := r.(error); !ok || err.Error() != expectErr {
			t.Errorf("unexpected panic: %v", r)
		}
	}()

	open, err := os.Open("testdata/classic.json")
	if err != nil {
		t.Fatal(err)
	}

	open.Close() // close it so we get an error!
	reader := file.NewReader(open, json.Unmarshal)

	value := f.Config{}
	uconfig.Must(&value, reader)
}
