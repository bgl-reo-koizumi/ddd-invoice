package manthlyfee

import (
	"github.com/bgl-reo-koizumi/ddd-invoice/src/domain/member"
)

type Manthlyfee interface {
	canManthlyfee()
}

type telManthlyfee struct {
	fee int
	startDate int
	endDate int
	course string
	member member.Member
}

type tvManthlyfee struct {
	fee int
	startDate int
	endDate int
	course string
	member member.Member
}

type lineManthlyfee struct {
	fee int
	startDate int
	endDate int
	course string
	member member.Member
}


func (m *tvManthlyfee) canManthlyfee() {
	
}

func (m *telManthlyfee) canManthlyfee() {
	
}

func (m *lineManthlyfee) canManthlyfee() {
	
}