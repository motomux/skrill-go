package target

// Target is target value for return_url_target and cancel_url_target
type Target int

// Target lists
const (
	Top    Target = 1
	Parent Target = 2
	Self   Target = 3
	Blank  Target = 4
)
