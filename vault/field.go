package vault

import (
	"github.com/vinc3m1/opvault"
)

// Field interface
type Field interface {
	Name() string
}

type opVaultField struct {
	fieldInternal *opvault.Field
}

func newOpVaultField(field *opvault.Field) Field {
	return &opVaultField{
		fieldInternal: field,
	}
}

func (f *opVaultField) Name() string {
	return f.fieldInternal.Name()
}
