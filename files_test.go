package swaggerFiles

import (
	"testing"
)

func TestOpenFiles(t *testing.T) {
	for _, filename := range []string{
		"services.json",
		"doc.html",
		"robots.txt",
	} {

		_, err := FS.Open(filename)
		if err != nil {
			t.Fatalf("unable to open %s: %s", filename, err)
		}

	}
}
