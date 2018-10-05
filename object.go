package tinfer

// Object is an object type.
type Object interface {
	Type
	isObject()
}
