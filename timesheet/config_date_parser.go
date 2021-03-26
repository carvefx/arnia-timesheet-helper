package timesheet

// configDateParser is an interface that ensures special date structs allow input from CLI
// and can self validate against a single day
type configDateParser interface {
	Parse(cfg Config, ldom int)
	HasDate(day int) bool
	GetDates() []int
}
