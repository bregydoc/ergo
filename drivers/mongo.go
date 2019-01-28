package drivers

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
)

type MongoBag struct {
}

func NewMongoBag() (*MongoBag, error) {
	// db, err := badger.Open(opts)
	panic("unimplemented")
}

// GetAllErrors implements Ergo bag
func (b *MongoBag) GetAllErrors() ([]*schema.Error, error) {
	panic("unimplemented")
}

// GetAllErrorsFromCode implements Ergo bag
func (b *MongoBag) GetAllErrorsFromCode(code uint64) ([]*schema.Error, error) {
	panic("unimplemented")
}

// GetErrorByID implements Ergo bag
func (b *MongoBag) GetErrorByID(id ulid.ULID) (*schema.Error, error) {
	panic("unimplemented")
}

// GetErrorByNativeError implements Ergo bag
func (b *MongoBag) GetErrorByNativeError(errN error) (*schema.Error, error) {
	panic("unimplemented")
}

// RegisterNewErrorFromNative implements Ergo bag
func (b *MongoBag) RegisterNewErrorFromNative(errN error, message ...string) (*schema.Error, error) {
	panic("unimplemented")

}

// RegisterNewError implements Ergo bag
func (b *MongoBag) RegisterNewError(ergoError *schema.Error) (*schema.Error, error) {
	panic("unimplemented")
}

// UpdateErrorByID implements Ergo bag
func (b *MongoBag) UpdateErrorByID(id ulid.ULID, update *schema.Error) (*schema.Error, error) {
	panic("unimplemented")
}

// UpdateErrorByNative implements Ergo bag
func (b *MongoBag) UpdateErrorByNative(errN error, update *schema.Error) (*schema.Error, error) {
	panic("unimplemented")
}

// RemoveErrorByID implements Ergo bag
func (b *MongoBag) RemoveErrorByID(id ulid.ULID) (*schema.Error, error) {
	panic("unimplemented")
}

// RemoveErrorByNative implements Ergo bag
func (b *MongoBag) RemoveErrorByNative(errN error) (*schema.Error, error) {
	panic("unimplemented")
}

// RegisterNewLang implements Ergo bag
func (b *MongoBag) RegisterNewLang(language *schema.Language) (*schema.Language, error) {
	panic("unimplemented")
}

// AddNewLangToError implements Ergo bag
func (b *MongoBag) AddNewLangToError(errorID ulid.ULID, language *schema.Language, message string) (*schema.Language, error) {
	panic("unimplemented")
}

// GetLanguageByID implements Ergo bag
func (b *MongoBag) GetLanguageByID(id ulid.ULID) (*schema.Language, error) {
	panic("unimplemented")
}

// UpdateLanguageByID implements Ergo bag
func (b *MongoBag) UpdateLanguageByID(id ulid.ULID, update *schema.Language) (*schema.Language, error) {
	panic("unimplemented")
}

// RemoveLanguageByID implements Ergo bag
func (b *MongoBag) RemoveLanguageByID(id ulid.ULID) (*schema.Language, error) {
	panic("unimplemented")
}
