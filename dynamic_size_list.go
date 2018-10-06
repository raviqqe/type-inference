package tinfer

// DynamicSizeList is a uniform list type.
type DynamicSizeList struct {
	element  Type
	location string
}

// NewDynamicSizeList creates a new uniform list.
func NewDynamicSizeList(t Type, l string) *DynamicSizeList {
	return &DynamicSizeList{t, l}
}

// Unify unifies 2 types.
func (l *DynamicSizeList) Unify(t Type) error {
	ll, ok := t.(List)

	if !ok {
		return fallback(l, t, "not a list")
	} else if ll, ok := ll.(*StaticSizeList); ok {
		for _, e := range ll.elements {
			if err := l.element.Unify(e); err != nil {
				return err
			}
		}

		return nil
	}

	return l.element.Unify(ll.(*DynamicSizeList).element)
}

// Location returns where the type is defined.
func (l *DynamicSizeList) Location() string {
	return l.location
}

func (*DynamicSizeList) isList() {}
