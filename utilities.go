package tinfer

func fallback(t, tt *Type, m string) error {
	if _, ok := (*tt).(Variable); !ok {
		return newInferenceError(m, (*tt).Location())
	}

	*tt = *t

	return nil
}

func unifyMany(ts, tts []Type, m, l string) error {
	if len(ts) != len(tts) {
		return newInferenceError(m, l)
	}

	for i, t := range ts {
		if err := t.Unify(&ts[i], &tts[i]); err != nil {
			return err
		}
	}

	return nil
}
