package jobs

// Job represents a job that can be executed.
type Job interface {
	Do() error
}
