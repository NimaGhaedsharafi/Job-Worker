package worker

// Job represents the job to be run
type Job struct {
	SSN string
	Result Result
}

type Result struct {
	Status int
}

func (j *Job) Inquiry() error {
	return nil
}