package types_test

import (
	"github.com/stafiprotocol/scale.go/types"
	"testing"
)

func TestRegCustomTypes(t *testing.T) {
	types.RuntimeType{}.Reg()
}
