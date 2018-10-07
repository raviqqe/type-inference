package tinfer

// Function is a function type.
type Function struct {
	arguments []Type
	result    Type
	location  string
}

// NewFunction creates a new function.
func NewFunction(as []Type, r Type, l string) *Function {
	return &Function{as, r, l}
}

// Accept accepts another type.
func (f *Function) Accept(t Type) error {
	ff, ok := t.(*Function)

	if !ok {
		return newInferenceError("not a function", t.Location())
	} else if err := acceptMany(ff.arguments, f.arguments, "different number of arguments", ff.location); err != nil {
		return err
	}

	return f.result.Accept(ff.result)
}

// CanAccept checks if a type is acceptable.
func (f Function) CanAccept(t Type) bool {
	ff, ok := t.(*Function)

	if !ok {
		return false
	} else if len(f.arguments) != len(ff.arguments) {
		return false
	}

	for i, t := range f.arguments {
		// contravariance check
		if !ff.arguments[i].CanAccept(t) {
			return false
		}
	}

	// covariance check
	return f.result.CanAccept(ff.result)
}

// Location returns where the type is defined.
func (f Function) Location() string {
	return f.location
}
