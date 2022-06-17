package mocks

import (
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/stretchr/testify/mock"
)
 
type SectionRepository struct {
	mock.Mock
}

func (sectionRepository *SectionRepository) ListarSectionAll()([]section.Section, error) {

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

func (SectionRepository *SectionRepository) ListarSectionOne(id int) ([]section.Section, error) {
	args := SectionRepository.Called()

	var sectionList []section.Section

	if rf, ok := args.Get(0).(func(id int) []section.Section); ok {
		sectionList = rf(id)
	} else {
		if args.Get(0) != nil {
			sectionList = rf(id)
		}
	}

	var err error 

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return sectionList, err
}

func (sectionRepository *SectionRepository) CreateSection(newSection section.Section) (section.Section, error){

	var (
		sectionObjOfReturn section.Section
		err error
		args = sectionRepository.Called()
	)

	if rf, ok := args.Get(0).(func(newSection section.Section) section.Section); ok {
		sectionObjOfReturn = rf(newSection)
	} else {
		if args.Get(0) != nil {
			sectionObjOfReturn = rf(newSection)
		}
	}

	if rf, ok := args.Get(1).(func(newSection section.Section) error); ok {
		err = rf(newSection)
	} else {
		if args.Get(1) != nil{
			err = args.Error(1)
		}		
	}

	return sectionObjOfReturn, err
}

func (sectionRepository *SectionRepository) DeleteSection(id int) error {
	var (
		err error 
		args = sectionRepository.Called()
	)

	if rf, ok := args.Get(0).(func(id int) error); ok {
		err = rf(id)
	} else {
		err = args.Error(0)
	}

	return err
}

func (sectionRepository *SectionRepository) LastID() (int, error) {
	var (
		args = sectionRepository.Called()
		id int 
		err error 
	)

	if rf, ok := args.Get(0).(func() int); ok {
		id = rf()
	} else {
		if args.Get(0) != nil {
			id = rf()
		}
	}

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}

	return id, err
}