package drivers

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/dgraph-io/badger"
	"github.com/gogo/protobuf/proto"
	"github.com/oklog/ulid"
)

type BadgerBag struct {
	db *badger.DB
}

func NewBagerBag(dir, valueDir string, opts ...badger.Options) (*BadgerBag, error) {
	// db, err := badger.Open(opts)
	opt := badger.DefaultOptions
	if len(opts) > 0 {
		opt = opts[1]
	}
	opt.Dir = dir
	opt.ValueDir = valueDir
	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	return &BadgerBag{
		db: db,
	}, nil
}

// GetAllErrors implements Ergo bag
func (b *BadgerBag) GetAllErrors() ([]*schema.Error, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()
	iterator := txn.NewIterator(badger.IteratorOptions{})

	errors := make([]*schema.Error, 0)

	for i := iterator.Item(); iterator.Valid(); iterator.Next() {
		data, err := i.ValueCopy(nil)
		if err != nil {
			return nil, err
		}

		e := new(schema.Error)

		err = proto.Unmarshal(data, e)
		if err != nil {
			return nil, err
		}

		errors = append(errors, e)
	}

	return errors, nil
}

// GetAllErrorsFromCode implements Ergo bag
func (b *BadgerBag) GetAllErrorsFromCode(code uint64) ([]*schema.Error, error) {
	panic("unimplemented")
	return []*schema.Error{}, nil
}

// GetErrorByID implements Ergo bag
func (b *BadgerBag) GetErrorByID(id ulid.ULID) (*schema.Error, error) {
	txn := b.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(id[:])
	if err != nil {
		return nil, err
	}

	data, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	ergoError := new(schema.Error)

	err = proto.Unmarshal(data, ergoError)
	if err != nil {
		return nil, err
	}

	return ergoError, nil
}

// GetErrorByNativeError implements Ergo bag
func (b *BadgerBag) GetErrorByNativeError(errN error) (*schema.Error, error) {
	panic("unimplemented")
	return &schema.Error{}, nil
}

// RegisterNewErrorFromNative implements Ergo bag
func (b *BadgerBag) RegisterNewErrorFromNative(errN error, message ...string) (*schema.Error, error) {
	panic("unimplemented")
	return &schema.Error{}, nil
}

// RegisterNewError implements Ergo bag
func (b *BadgerBag) nRegisterNewError(ergoError *schema.Error) (*schema.Error, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	data, err := proto.Marshal(ergoError)
	if err != nil {
		return nil, err
	}

	err = txn.Set(ergoError.Id, data)

	var id ulid.ULID
	copy(id[:], ergoError.Id)

	return b.GetErrorByID(id)
}

// UpdateErrorByID implements Ergo bag
func (b *BadgerBag) UpdateErrorByID(id ulid.ULID, update *schema.Error) (*schema.Error, error) {
	panic("unimplemented")
}

// UpdateErrorByNative implements Ergo bag
func (b *BadgerBag) UpdateErrorByNative(errN error, update *schema.Error) (*schema.Error, error) {
	panic("unimplemented")
}

// RemoveErrorByID implements Ergo bag
func (b *BadgerBag) RemoveErrorByID(id ulid.ULID) (*schema.Error, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	sch, err := b.GetErrorByID(id)
	if err != nil {
		return nil, err
	}

	err = txn.Delete(id[:])
	if err != nil {
		return nil, err
	}

	return sch, nil
}

// RemoveErrorByNative implements Ergo bag
func (b *BadgerBag) RemoveErrorByNative(errN error) (*schema.Error, error) {
	panic("unimplemented")
	return &schema.Error{}, nil
}
