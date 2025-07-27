package category

type CategoryError struct {
	msg string
}

func (c CategoryError) Error() string {
	return c.msg
}
