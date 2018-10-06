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

// Accept accepts another type.
func (v *Variable) Accept(t Type) error {
	if v.inferredType != nil {
		return v.inferredType.Accept(t)
	}

	v.inferredType = t

	return nil
}

// Location returns where the type is defined.
func (v Variable) Location() string {
	return v.location
}
