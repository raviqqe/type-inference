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

// Unify unifies 2 types.
func (o *MapObject) Unify(t Type) error {
	oo, ok := t.(Object)

	if !ok {
		return fallback(o, t, "not an object")
	} else if oo, ok := oo.(*RecordObject); ok {
		for _, kv := range oo.keyValues {
			if err := o.value.Unify(kv.value); err != nil {
				return err
			}
		}

		return nil
	}

	return o.value.Unify(oo.(*MapObject).value)
}

// Location returns where the type is defined.
func (o *MapObject) Location() string {
	return o.location
}

func (*MapObject) isObject() {}
