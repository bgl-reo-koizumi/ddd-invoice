package manthlyfee

import (
	"github.com/bgl-reo-koizumi/ddd-invoice/src/domain/member"
)

type IManthlyfee interface {
	isManthlyfee()
}

type Manthlyfee struct {
	fee int
	startDate int
	endDate int
	course string
	member member.Member
}


func (m *Manthlyfee) telManthlyfee() {
	
}

