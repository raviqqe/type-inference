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

// Unify unifies 2 types.
func (l *StaticSizeList) Unify(t, tt *Type) error {
	ll, ok := (*tt).(List)

	if !ok {
		return fallback(t, tt, "not a list")
	} else if ll, ok := ll.(*DynamicSizeList); ok {
		return ll.Unify(tt, t)
	}

	return unifyMany(
		l.elements,
		ll.(*StaticSizeList).elements,
		"different number of elements",
		ll.Location(),
	)
}

// Location returns where the type is defined.
func (l *StaticSizeList) Location() string {
	return l.location
}

func (*StaticSizeList) isList() {}
