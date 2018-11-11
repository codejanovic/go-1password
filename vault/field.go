package vault

import (
	"github.com/vinc3m1/opvault"
)

// Field interface
type Field interface {
	Name() string
	Value(showPassword bool) string
	IsPassword() bool
	IsEmail() bool
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

func (f *opVaultField) IsPassword() bool {
	return f.fieldInternal.Type() == opvault.PasswordFieldType
}

func (f *opVaultField) IsEmail() bool {
	return f.fieldInternal.Type() == opvault.EmailFieldType
}

func (f *opVaultField) Value(showPassword bool) string {
	if f.IsPassword() && !showPassword {
		return "***********"
	}
	return f.fieldInternal.Value()
}
