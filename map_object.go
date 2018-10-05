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
func (o *MapObject) Unify(t, tt *Type) error {
	oo, ok := (*tt).(Object)

	if !ok {
		return fallback(t, tt, "not an object")
	}

	return o.value.Unify(&o.value, &oo.(*MapObject).value)
}

// Location returns where the type is defined.
func (o *MapObject) Location() string {
	return o.location
}

func (*MapObject) isObject() {}
