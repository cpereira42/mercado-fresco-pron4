package mocks

import (
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/stretchr/testify/mock"
)

type SectionRepository struct {
	mock.Mock
}

/*
	ListarSectionAll() ([]Section, error)
	ListarSectionOne(id int) (Section, error)
	CreateSection(section Section) (Section, error)
	UpdateSection(id int, section Section) (Section, error)
	DeleteSection(id int) error
	LastID() (int, error)
*/
func (sectionRepository *SectionRepository) ListarSectionAll() ([]section.Section, error) {
	args := sectionRepository.Called()

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

func (sectionRepository *SectionRepository) ListarSectionOne(id int64) (section.Section, error) {
	args := sectionRepository.Called(id)

	var sectionObj section.Section
	if rf, ok := args.Get(0).(func(int64) section.Section); ok {
		sectionObj = rf(id)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(section.Section)
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
	return sectionObj, err
}

func (sectionRepository *SectionRepository) CreateSection(newSection section.Section) (section.Section, error) {
	args := sectionRepository.Called(newSection)

	var sectionObj section.Section
	if rf, ok := args.Get(0).(func(section.Section) section.Section); ok {
		sectionObj = rf(newSection)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(section.Section)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func(section.Section) error); ok {
		err = rf(newSection)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionObj, err
}

func (sectionRepository *SectionRepository) UpdateSection(id int64, sectionUp section.Section) (section.Section, error) {
	args := sectionRepository.Called(id, sectionUp)

	var sectionObj section.Section
	if rf, ok := args.Get(0).(func(int64, section.Section) section.Section); ok {
		sectionObj = rf(id, sectionObj)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(section.Section)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func(int64, section.Section) error); ok {
		err = rf(id, sectionObj)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionObj, err
}

func (sectionRepository *SectionRepository) DeleteSection(id int64) error {
	args := sectionRepository.Called(id)

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
 