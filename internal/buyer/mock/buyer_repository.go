package mock

import (
	"github.com/cpereira42/mercado-fresco-pron4/internal/buyer"
	"github.com/stretchr/testify/mock"
)

type BuyerRepository struct {
	mock.Mock
}

/*
GetAll() ([]Buyer, error)
GetId(id int) (Buyer, error)
Create(id int, card_number_ID, first_name, last_name string) (Buyer, error)
LastID() (int, error)
Update(id int, card_number_ID, first_name, last_name string) (Buyer, error)
Delete(id int) error

*/

func (b *BuyerRepository) Create(
	id int,
	card_number_ID, first_name,
	last_name string,
) (buyer.Buyer, error) {
	var (
		args     = b.Called(id, card_number_ID, first_name, last_name)
		err      error
		buyerObj buyer.Buyer
	)

	if rf, ok := args.Get(0).(func(int, string, string, string) buyer.Buyer); ok {
		buyerObj = rf(id, card_number_ID, first_name, last_name)
	} else {
		if args.Get(0) != nil {
			buyerObj = args.Get(0).(buyer.Buyer)
		}
	}

	if rf, ok := args.Get(1).(func(int, string, string, string) error); ok {
		err = rf(id, card_number_ID, first_name, last_name)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return buyerObj, err
}

func (b *BuyerRepository) GetAll() ([]buyer.Buyer, error) {
	args := b.Called()
	var (
		buyerList []buyer.Buyer
		err       error
	)

	if rf, ok := args.Get(0).(func() []buyer.Buyer); ok {
		buyerList = rf()
	} else {
		if args.Get(0) != nil {
			buyerList = args.Get(0).([]buyer.Buyer)
		}
	}

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return buyerList, err
}

func (b *BuyerRepository) GetId(id int) (buyer.Buyer, error) {
	var (
		args     = b.Called(id)
		err      error
		buyerObj buyer.Buyer
	)

	if rf, ok := args.Get(0).(func(id int) buyer.Buyer); ok {
		buyerObj = rf(id)
	} else {
		if args.Get(0) != nil {
			buyerObj = args.Get(0).(buyer.Buyer)
		}
	}

	if rf, ok := args.Get(1).(func(id int) error); ok {
		err = rf(id)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return buyerObj, err
}

func (b *BuyerRepository) LastID() (int, error) {
	var (
		err  error
		args = b.Called()
		id   int
	)

	if rf, ok := args.Get(0).(func() int); ok {
		id = rf()
	} else {
		if args.Get(0) != nil {
			id = args.Get(0).(int)
		}
	}

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}
	return id, err
}

func (b *BuyerRepository) Update(
	id int,
	card_number_ID, first_name, last_name string,
) (buyer.Buyer, error) {
	var (
		args     = b.Called(id, card_number_ID, first_name, last_name)
		buyerObj buyer.Buyer
		err      error
	)
	if rf, ok := args.Get(0).(func(
		int, string, string, string,
	) buyer.Buyer); ok {
		buyerObj = rf(id, card_number_ID, first_name, last_name)
	} else {
		if args.Get(1) != nil {
			buyerObj = args.Get(0).(buyer.Buyer)
		}
	}

	if rf, ok := args.Get(1).(func(
		id int,
		card_number_ID, first_name, last_name string,
	) error); ok {
		err = rf(id, card_number_ID, first_name, last_name)
	} else {
		err = args.Error(1)
	}

	return buyerObj, err
}

func (b *BuyerRepository) Delete(id int) error {
	var (
		args = b.Called(id)
		err  error
	)

	if rf, ok := args.Get(0).(func(id int) error); ok {
		err = rf(id)
	} else {
		err = args.Error(0)
	}

	return err
}
