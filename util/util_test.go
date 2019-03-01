package util_test

import (
	"fmt"
	"github.com/autom8ter/identify/util"
	"testing"
)

func TestRandomToken(t *testing.T) {
	state, err := util.RandomToken()
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(state)
	}
}
