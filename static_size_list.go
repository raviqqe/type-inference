package tinfer

// StaticSizeList is a tuple type represented as a list.
type StaticSizeList struct {
	elements []Type
	location string
}

// NewStaticSizeList creates a new list.
func NewStaticSizeList(ts []Type, l string) *StaticSizeList {
	return &StaticSizeList{ts, l}
}

// Accept accepts another type.
func (l *StaticSizeList) Accept(t Type) error {
	ll, ok := t.(List)

	if !ok {
		return newInferenceError("not a list", t.Location())
	} else if ll, ok := ll.(*DynamicSizeList); ok {
		for _, e := range l.elements {
			if err := e.Accept(ll.element); err != nil {
				return err
			}
		}

		return nil
	}

	return acceptMany(
		l.elements,
		ll.(*StaticSizeList).elements,
		"different number of elements",
		ll.Location(),
	)
}

// Location returns where the type is defined.
func (l StaticSizeList) Location() string {
	return l.location
}

func (*StaticSizeList) isList() {}
