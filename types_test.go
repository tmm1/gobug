package gobug

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type obj struct {
	Name      string
	Info      struct {
		Foo string
		Bar string
	}
}

func TestTypeInfo(t *testing.T) {
	e := &obj{}
	typ := reflect.ValueOf(e).Elem().Type()
	i := typeInfoFor(typ)

	finfo := i.fields["Name"]
	assert.Equal(t, []int{0}, finfo.index)

	finfo = i.fields["Info.Foo"]
	assert.Equal(t, []int{1, 0}, finfo.index)

	finfo = i.fields["Info.Bar"]
	assert.Equal(t, []int{1, 1}, finfo.index)
}
