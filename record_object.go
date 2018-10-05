package tinfer

import (
	"sort"
)

type keyValue struct {
	key   string
	value Type
}

// RecordObject is a record object type.
type RecordObject struct {
	keyValues []keyValue // sorted by keys
	location  string
}

// NewRecordObject creates a new record object.
func NewRecordObject(m map[string]Type, l string) *RecordObject {
	ks := make([]string, 0, len(m))

	for k := range m {
		ks = append(ks, k)
	}

	sort.Strings(ks)
	kvs := make([]keyValue, 0, len(ks))

	for _, k := range ks {
		kvs = append(kvs, keyValue{k, m[k]})
	}

	return &RecordObject{kvs, l}
}

// Unify unifies 2 types.
func (o *RecordObject) Unify(t, tt *Type) error {
	oo, ok := (*tt).(Object)

	if !ok {
		return fallback(t, tt, "not an object")
	} else if oo, ok := oo.(*MapObject); ok {
		return oo.Unify(tt, t)
	}

	return unifyRecordObjects(o, oo.(*RecordObject), t, tt)
}

// Location returns where the type is defined.
func (o RecordObject) Location() string {
	return o.location
}

func (RecordObject) isObject() {}

func (o RecordObject) contain(oo *RecordObject) bool {
	ks := oo.keys()

	for _, k := range o.keys() {
		i := sort.Search(len(ks), func(i int) bool { return ks[i] >= k })

		if i < 0 || i >= len(ks) || ks[i] != k {
			return false
		}
	}

	return true
}

func (o RecordObject) keys() []string {
	ks := make([]string, 0, len(o.keyValues))

	for _, kv := range o.keyValues {
		ks = append(ks, kv.key)
	}

	return ks
}

func unifyRecordObjects(o, oo *RecordObject, t, tt *Type) error {
	if !o.contain(oo) && !oo.contain(o) {
		return newInferenceError("not a compatible object", oo.Location())
	} else if !o.contain(oo) && oo.contain(o) {
		return unifyRecordObjects(oo, o, tt, t)
	}

	for _, kv1 := range o.keyValues {
		for _, kv2 := range oo.keyValues {
			if kv1.key == kv2.key {
				if err := kv1.value.Unify(&kv1.value, &kv2.value); err != nil {
					return err
				}
			}
		}
	}

	*t = *tt

	return nil
}
