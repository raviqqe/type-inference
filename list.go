package tinfer

// List is a list type.
type List interface {
	Type
	isList()
}
