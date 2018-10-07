package tinfer

import (
	"sort"
)

// RecordObject is a record object type.
type RecordObject struct {
	keyValues map[string]Type
	location  string
}

// NewRecordObject creates a new record object.
func NewRecordObject(m map[string]Type, l string) *RecordObject {
	return &RecordObject{m, l}
}

// Accept accepts another type.
func (o *RecordObject) Accept(t Type) error {
	oo, ok := t.(*RecordObject)

	if !ok {
		return newInferenceError("not an record object", t.Location())
	} else if !o.contain(oo) {
		return newInferenceError("not a compatible record object", oo.Location())
	}

	for k, v := range o.keyValues {
		if err := v.Accept(oo.keyValues[k]); err != nil {
			return err
		}
	}

	return nil
}

// CanAccept checks if a type is acceptable.
func (o RecordObject) CanAccept(t Type) bool {
	oo, ok := t.(*RecordObject)

	if !ok {
		return false
	} else if !o.contain(oo) {
		return false
	}

	for k, v := range o.keyValues {
		if !v.CanAccept(oo.keyValues[k]) {
			return false
		}
	}

	return true
}

// Location returns where the type is defined.
func (o RecordObject) Location() string {
	return o.location
}

func (RecordObject) isObject() {}

func (o RecordObject) contain(oo *RecordObject) bool {
	ks := oo.keys()

	for k := range o.keyValues {
		i := sort.Search(len(ks), func(i int) bool { return ks[i] >= k })

		if i < 0 || i >= len(ks) || ks[i] != k {
			return false
		}
	}

	return true
}

func (o RecordObject) keys() []string {
	ks := make([]string, 0, len(o.keyValues))

	for k := range o.keyValues {
		ks = append(ks, k)
	}

	sort.Strings(ks)

	return ks
}
