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
func (l *DynamicSizeList) Unify(t, tt *Type) error {
	ll, ok := (*tt).(List)

	if !ok {
		return fallback(t, tt, "not a list")
	} else if ll, ok := ll.(*StaticSizeList); ok {
		for _, e := range ll.elements {
			if err := l.element.Unify(&l.element, &e); err != nil {
				return err
			}
		}

		*tt = l

		return nil
	}

	return l.element.Unify(&l.element, &ll.(*DynamicSizeList).element)
}

// Location returns where the type is defined.
func (l *DynamicSizeList) Location() string {
	return l.location
}

func (*DynamicSizeList) isList() {}
