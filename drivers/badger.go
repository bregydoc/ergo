package drivers

import (
	"github.com/bregydoc/ergo"
	"github.com/dgraph-io/badger"
	"github.com/oklog/ulid"
)

type BadgerBag struct {
	db *badger.DB
}

func NewBagerBag() {
	// db, err := badger.Open(opts)

}

// GetAllErrors implements Ergo bag
func (b *BadgerBag) GetAllErrors() ([]*ergo.Error, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	return []*ergo.Error{}, nil
}

// GetAllErrorsFromCode implements Ergo bag
func (b *BadgerBag) GetAllErrorsFromCode(code uint64) ([]*ergo.Error, error) {

	return []*ergo.Error{}, nil
}

// GetErrorByID implements Ergo bag
func (b *BadgerBag) GetErrorByID(id ulid.ULID) (*ergo.Error, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	item, err := txn.Get(id[:])
	if err != nil {
		return nil, err
	}

	_, err = item.ValueCopy(nil)

	// TODO: Encode data to ergo.Error struct

	if err != nil {
		return nil, err
	}
	return &ergo.Error{}, nil
}

// GetErrorByNativeError implements Ergo bag
func (b *BadgerBag) GetErrorByNativeError(err error) (*ergo.Error, error) {
	return &ergo.Error{}, nil
}

// RegisterNewErrorFromNative implements Ergo bag
func (b *BadgerBag) RegisterNewErrorFromNative(err error, message ...string) (*ergo.Error, error) {
	return &ergo.Error{}, nil
}

// RegisterNewError implements Ergo bag
func (b *BadgerBag) RegisterNewError(err *ergo.Error) (*ergo.Error, error) {

	return &ergo.Error{}, nil
}

// RemoveErrorByID implements Ergo bag
func (b *BadgerBag) RemoveErrorByID(id ulid.ULID) (*ergo.Error, error) {

	return &ergo.Error{}, nil
}

// RemoveErrorByNative implements Ergo bag
func (b *BadgerBag) RemoveErrorByNative(err error) (*ergo.Error, error) {

	return &ergo.Error{}, nil
}
