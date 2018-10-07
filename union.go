package tinfer

// Union is a union type.
type Union struct {
	types    []Type
	location string
}

// NewUnion creates a new union type.
func NewUnion(ts []Type, l string) Union {
	if len(ts) < 2 {
		panic("types must be more than 1.")
	}

	return Union{ts, l}
}

// Accept unifies 2 types.
func (u Union) Accept(t Type) error {
	if uu, ok := t.(Union); ok {
		return u.unifyUnion(uu)
	}

	for _, tt := range u.types {
		if tt.CanAccept(t) {
			return tt.Accept(t)
		}
	}

	return newInferenceError("not compatible with union", t.Location())
}

// CanAccept checks if a type is acceptable.
func (Union) CanAccept(t Type) bool {
	panic("unimplemented")
}

// Location returns where the type is defined.
func (u Union) Location() string {
	return u.location
}

func (u Union) unifyUnion(uu Union) error {
top:
	for _, tt := range uu.types {
		for _, t := range u.types {
			if t.CanAccept(tt) {
				if err := t.Accept(tt); err != nil {
					panic("unreachable")
				}

				continue top
			}
		}

		return newInferenceError("union not compatible", uu.Location())
	}

	return nil
}
