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
