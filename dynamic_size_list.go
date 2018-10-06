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

// Accept accepts another type.
func (l *DynamicSizeList) Accept(t Type) error {
	ll, ok := t.(List)

	if !ok {
		return fallback(l, t, "not a list")
	} else if ll, ok := ll.(*StaticSizeList); ok {
		for _, e := range ll.elements {
			if err := l.element.Accept(e); err != nil {
				return err
			}
		}

		return nil
	}

	return l.element.Accept(ll.(*DynamicSizeList).element)
}

// Location returns where the type is defined.
func (l *DynamicSizeList) Location() string {
	return l.location
}

func (*DynamicSizeList) isList() {}
