package mocks

import ( 
	"github.com/cpereira42/mercado-fresco-pron4/internal/section/entites"
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
func (sectionRepository *SectionRepository) ListarSectionAll()([]entites.Section, error) {

	args := sectionRepository.Called()

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
func (sectionRepository *SectionRepository) ListarSectionOne(id int) (entites.Section, error) {
	var args = sectionRepository.Called(id)

	var sectionObj entites.Section
	if rf, ok := args.Get(0).(func(int) entites.Section); ok {
		sectionObj = rf(id)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(entites.Section)
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
	return sectionObj, err
}
func (sectionRepository *SectionRepository) CreateSection(newSection entites.Section) (entites.Section, error) {
	var args = sectionRepository.Called(newSection)

	var sectionObj entites.Section
	if rf, ok := args.Get(0).(func(entites.Section) entites.Section); ok {
		sectionObj = rf(newSection)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(entites.Section)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func(entites.Section) error); ok {
		err = rf(newSection)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionObj, err 
}
func (sectionRepository *SectionRepository) UpdateSection(id int, sectionUp entites.Section) (entites.Section, error) {
	var args = sectionRepository.Called(id,sectionUp)

	var sectionObj entites.Section
	if rf, ok := args.Get(0).(func(int,entites.Section) entites.Section); ok {
		sectionObj = rf(id,sectionObj)
	} else {
		if args.Get(0) != nil {
			sectionObj = args.Get(0).(entites.Section)
		}
	}

	var err error
	if rf, ok := args.Get(1).(func(int,entites.Section) error); ok {
		err = rf(id,sectionObj)
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}
	return sectionObj, err 
}
func (sectionRepository *SectionRepository) DeleteSection(id int) error {
	var args = sectionRepository.Called(id)

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
func (sectionRepository *SectionRepository) LastID() (int, error) {
	var args = sectionRepository.Called()

	var lastID int
	if rf, ok := args.Get(0).(func() int); ok {
		lastID = rf()
	} else {
		if args.Get(0) != nil {
			lastID = args.Get(0).(int)
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
	return lastID, err 
}
