package tinfer

// Variable is a type variable.
type Variable struct {
	inferredType Type
	location     string
}

// NewVariable creates a new variable.
func NewVariable(l string) *Variable {
	return &Variable{nil, l}
}

// Unify unifies 2 types.
func (v *Variable) Unify(t Type) error {
	if v.inferredType != nil {
		return v.inferredType.Unify(t)
	}

	v.inferredType = t

	return nil
}

// Location returns where the type is defined.
func (v Variable) Location() string {
	return v.location
}
