package tinfer

func acceptMany(ts, tts []Type, m, l string) error {
	if len(ts) != len(tts) {
		return newInferenceError(m, l)
	}

	for i, t := range ts {
		if err := t.Accept(tts[i]); err != nil {
			return err
		}
	}

	return nil
}

func acceptUnion(t, tt Type, m string) error {
	u, ok := tt.(Union)

	if !ok {
		return newInferenceError(m, tt.Location())
	}

	for _, tt := range u.types {
		if err := t.Accept(tt); err != nil {
			return err
		}
	}

	return nil
}
