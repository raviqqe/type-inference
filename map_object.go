package tinfer

// MapObject is a map object type.
type MapObject struct {
	value    Type
	location string
}

// NewMapObject creates a new map object.
func NewMapObject(t Type, l string) *MapObject {
	return &MapObject{t, l}
}

// Accept accepts another type.
func (o *MapObject) Accept(t Type) error {
	oo, ok := t.(Object)

	if !ok {
		return acceptUnion(o, t, "not an object")
	} else if oo, ok := oo.(*RecordObject); ok {
		for _, v := range oo.keyValues {
			if err := o.value.Accept(v); err != nil {
				return err
			}
		}

		return nil
	}

	return o.value.Accept(oo.(*MapObject).value)
}

// CanAccept checks if a type is acceptable.
func (o MapObject) CanAccept(t Type) bool {
	switch oo := t.(type) {
	case *MapObject:
		return o.value.CanAccept(oo.value)
	case *RecordObject:
		for _, v := range oo.keyValues {
			if !o.value.CanAccept(v) {
				return false
			}
		}

		return true
	}

	return false
}

// Location returns where the type is defined.
func (o MapObject) Location() string {
	return o.location
}

func (MapObject) isObject() {}
