package resError

type Error struct {
	Title string
	Desc  string
}

func New(title string, desc string) *Error {
	return &Error{
		Title: title,
		Desc:  desc,
	}
}
