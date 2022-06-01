package manthlyfee

package

type IManthlyfee interface {
	factory()
}

type Manthlyfee struct {
	fee int
	startDate int
	endDate int
	course string
	member Member

}

