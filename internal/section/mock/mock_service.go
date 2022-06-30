package mocks

import (
	"github.com/cpereira42/mercado-fresco-pron4/internal/section/entites"
	"github.com/stretchr/testify/mock"
)

/*
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(section SectionRequestCreate) (Section, error)
	UpdateSection(id int, sectionUp SectionRequestUpdate) (Section, error)
	DeleteSection(id int) error
*/

type SectionService struct {
	mock.Mock
}
func (sectionService *SectionService) ListarSectionAll() ([]entites.Section, error) {
	args := sectionService.Called()

	var sectionList []entites.Section
	if rf, ok := args.Get(0).(func() []entites.Section); ok {
		sectionList = rf()
	} else {
		if args.Get(0) != nil {
			sectionList = args.Get(0).([]entites.Section)
		}
	}

	var err error 
	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionList, err
}
func (sectionService *SectionService) ListarSectionOne(id int) (entites.Section, error) {
	args := sectionService.Called(id)

	var sectionOne entites.Section
	if rf, ok := args.Get(0).(func(int) entites.Section); ok {
		sectionOne = rf(id)
	} else {
		if args.Get(0) != nil {
			sectionOne = args.Get(0).(entites.Section)
		}
	}

	var err error 
	if rf, ok := args.Get(1).(func(int) error); ok {
		err = rf(id)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionOne, err
}
func (sectionService *SectionService) CreateSection(sectionNew entites.SectionRequestCreate) (entites.Section, error) {
	args := sectionService.Called(sectionNew)

	var sectionObj entites.Section
	if rf, ok := args.Get(0).(func(entites.SectionRequestCreate) entites.Section); ok {
		sectionObj = rf(sectionNew)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(entites.Section)
		}
	}

	var err error 
	if rf, ok := args.Get(1).(func(entites.SectionRequestCreate) error); ok {
		err = rf(sectionNew)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionObj, err
}
func (sectionService *SectionService) UpdateSection(id int, sectionUp entites.SectionRequestUpdate) (entites.Section, error) {
	args := sectionService.Called(id, sectionUp)

	var sectionObj entites.Section
	if rf, ok := args.Get(0).(func(int,entites.SectionRequestUpdate) entites.Section); ok {
		sectionObj = rf(id,sectionUp)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(entites.Section)
		}
	}

	var err error 
	if rf, ok := args.Get(1).(func(int,entites.SectionRequestUpdate) error); ok {
		err = rf(id,sectionUp)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionObj, err
}
func (sectionService *SectionService) DeleteSection(id int) error {
	args := sectionService.Called(id)

	var err error 
	if rf, ok := args.Get(0).(func(int) error); ok {
		err = rf(id)
	} else {
		if args.Get(0) != nil {
			err = args.Error(0)
		}
	}
	return err
}
