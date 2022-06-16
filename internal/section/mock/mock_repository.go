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