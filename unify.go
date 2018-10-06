package tinfer

// Unify unifies 2 types into one.
func Unify(t, tt Type) error {
	return t.Unify(tt)
}
