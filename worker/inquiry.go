package worker

type Payload struct {
	SSN string
}

type Result struct {
	Status int
}

func (P *Payload) Inquiry() (*Result, error) {
	return nil, nil
}