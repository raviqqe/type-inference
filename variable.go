package tinfer

// Variable is a type variable.
type Variable struct {
	location string
}

// NewVariable creates a new variable.
func NewVariable(l string) Variable {
	return Variable{l}
}

// Unify unifies 2 types.
func (Variable) Unify(t, tt *Type) error {
	*t = *tt
	return nil
}

// Location returns where the type is defined.
func (v Variable) Location() string {
	return v.location
}
