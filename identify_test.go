package identify_test

import (
	"github.com/autom8ter/identify"
	"github.com/autom8ter/identify/options"
	"testing"
)

func TestNew(t *testing.T) {
	boss := identify.New(
		options.Empty(),
		options.WithDefaults(true, true),
		)
	if boss == nil {
		t.Fatal("failed to create new")
	}
}
