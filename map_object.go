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
		return fallback(o, t, "not an object")
	} else if oo, ok := oo.(*RecordObject); ok {
		for _, kv := range oo.keyValues {
			if err := o.value.Accept(kv.value); err != nil {
				return err
			}
		}

		return nil
	}

	return o.value.Accept(oo.(*MapObject).value)
}

// Location returns where the type is defined.
func (o *MapObject) Location() string {
	return o.location
}

func (*MapObject) isObject() {}
