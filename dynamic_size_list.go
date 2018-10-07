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
		return acceptUnion(l, t, "not a list")
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

// CanAccept checks if a type is acceptable.
func (l DynamicSizeList) CanAccept(t Type) bool {
	switch ll := t.(type) {
	case *DynamicSizeList:
		return l.element.CanAccept(ll.element)
	case *StaticSizeList:
		for _, t := range ll.elements {
			if !l.element.CanAccept(t) {
				return false
			}
		}

		return true
	}

	return false
}

// Location returns where the type is defined.
func (l DynamicSizeList) Location() string {
	return l.location
}

func (DynamicSizeList) isList() {}
