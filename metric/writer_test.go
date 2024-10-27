package metric

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/ToySin/finance/portfolio"
)

//go:embed testdata/portfolio.md
var testfile string

func TestWritePortfolio(t *testing.T) {
	p := portfolio.GetTestPortfolio()

	buffer := new(bytes.Buffer)
	writePortfolio(buffer, *p)

	if buffer.String() != testfile {
		t.Errorf("writePortfolio() = %v, want %v", buffer.String(), testfile)
	}
}
