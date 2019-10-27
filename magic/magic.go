package magic

const (
	// BentSpoonTrick consists in bending a spoon using the power of the magician's mind
	BentSpoonTrick = iota + 1
	// WringerTrick consists of getting the magician's stage assistant
	// through a wringer turning it into paper
	WringerTrick
)

// BentSpoonTrickError base type for the error that occurs when an attempt to
// perform the bent spoon strick fails.
type BentSpoonTrickError struct{}

func (e *BentSpoonTrickError) Error() string { return "Failed to bend the spoon" }

func newBentSpoonTrickError() error { return &BentSpoonTrickError{} }

// WringerTrickError base type for the error that occurs when an attempt to
// perform the wringer trick fails.
type WringerTrickError struct{}

func (e *WringerTrickError) Error() string { return "Failed to perform the wringer" }

func newWringerTrickError() error { return &WringerTrickError{} }

var (
	// ErrBentSpoonTrick occurs when an attempt to perform the bent spoon trick fails
	ErrBentSpoonTrick = newBentSpoonTrickError()
	// ErrWringerTrick occurs when an attempt to perform the wringer trick fails
	ErrWringerTrick = newWringerTrickError()
)

// DoMagic is an attempt to do magic in code, but it ended up just
// throwning errors based on the given value.
func DoMagic(value int) error {
	if value == 1 {
		return ErrBentSpoonTrick
	} else if value == 2 {
		return ErrWringerTrick
	}

	return nil
}
