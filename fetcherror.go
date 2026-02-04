package main

type fetchError struct {
	message string
}

func (e *fetchError) Error() string {
	return e.message
}
