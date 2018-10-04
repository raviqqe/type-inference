package tinfer

func fallback(t, tt *Type, m string) error {
	if _, ok := (*tt).(Variable); !ok {
		return newInferenceError(m, (*tt).Location())
	}

	*tt = *t

	return nil
}
