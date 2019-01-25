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
	return &schema.Error{}, nil
}

// RegisterNewErrorFromNative implements Ergo bag
func (b *MongoBag) RegisterNewErrorFromNative(errN error, message ...string) (*schema.Error, error) {
	panic("unimplemented")
	return &schema.Error{}, nil
}

// RegisterNewError implements Ergo bag
func (b *MongoBag) RegisterNewError(ergoError *schema.Error) (*schema.Error, error) {
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
