package mocks

import (
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/stretchr/testify/mock"
)

type SectionService struct {
	mock.Mock
}
func (sectionService *SectionService) ListarSectionAll() ([]section.Section, error) {
	args := sectionService.Called()

	var sectionList []section.Section
	if rf, ok := args.Get(0).(func() []section.Section); ok {
		sectionList = rf()
	} else {
		if args.Get(0) != nil {
			sectionList = args.Get(0).([]section.Section)
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
func (sectionService *SectionService) ListarSectionOne(id int64) (section.Section, error) {
	args := sectionService.Called(id)

	var sectionOne section.Section
	if rf, ok := args.Get(0).(func(int64) section.Section); ok {
		sectionOne = rf(id)
	} else {
		if args.Get(0) != nil {
			sectionOne = args.Get(0).(section.Section)
		}
	}

	var err error 
	if rf, ok := args.Get(1).(func(int64) error); ok {
		err = rf(id)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionOne, err
}
func (sectionService *SectionService) CreateSection(sectionNew section.SectionRequestCreate) (section.Section, error) {
	args := sectionService.Called(sectionNew)

	var sectionObj section.Section
	if rf, ok := args.Get(0).(func(section.SectionRequestCreate) section.Section); ok {
		sectionObj = rf(sectionNew)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(section.Section)
		}
	}

	var err error 
	if rf, ok := args.Get(1).(func(section.SectionRequestCreate) error); ok {
		err = rf(sectionNew)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionObj, err
}
func (sectionService *SectionService) UpdateSection(id int64, sectionUp section.SectionRequestUpdate) (section.Section, error) {
	args := sectionService.Called(id, sectionUp)

	var sectionObj section.Section
	if rf, ok := args.Get(0).(func(int64,section.SectionRequestUpdate) section.Section); ok {
		sectionObj = rf(id, sectionUp)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(section.Section)
		}
	}

	var err error 
	if rf, ok := args.Get(1).(func(int64, section.SectionRequestUpdate) error); ok {
		err = rf(id, sectionUp)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionObj, err
}
func (sectionService *SectionService) DeleteSection(id int64) error {
	args := sectionService.Called(id)

	var err error 
	if rf, ok := args.Get(0).(func(int64) error); ok {
		err = rf(id)
	} else {
		if args.Get(0) != nil {
			err = args.Error(0)
		}
	}
	return err
}
