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

// Unify unifies 2 types.
func (f *Function) Unify(t, tt *Type) error {
	ff, ok := (*tt).(*Function)

	if !ok {
		return fallback(t, tt, "not a function")
	} else if err := unifyMany(f.arguments, ff.arguments, "different number of arguments", ff.location); err != nil {
		return err
	} else if err := f.result.Unify(&f.result, &ff.result); err != nil {
		return err
	}

	return nil
}

// Location returns where the type is defined.
func (f *Function) Location() string {
	return f.location
}
