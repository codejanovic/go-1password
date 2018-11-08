package vault

import (
	"github.com/codejanovic/gordon/fatal"
	"github.com/vinc3m1/opvault"
)

// Item interface
type Item interface {
	Name() string
	FieldSize() int
	Fields() []Field
}

type opVaultItem struct {
	itemInternal   *opvault.Item
	fieldsInternal []Field
}

func newOpVaultItem(originalItem *opvault.Item) Item {
	originalDetail, err := originalItem.Detail()
	if err != nil {
		fatal.Crash(err, "Unable to fetch item details")
	}

	fields := make([]Field, 0)
	for _, originalField := range originalDetail.Fields() {
		fields = append(fields, newOpVaultField(originalField))
	}

	return &opVaultItem{
		itemInternal:   originalItem,
		fieldsInternal: fields,
	}
}

func (i *opVaultItem) Name() string {
	return i.itemInternal.Title()
}
func (i *opVaultItem) FieldSize() int {
	return len(i.fieldsInternal)
}
func (i *opVaultItem) Fields() []Field {
	return i.fieldsInternal
}
