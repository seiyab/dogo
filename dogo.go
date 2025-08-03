package dogo

type errSignal struct {
	err error
}

func Check(err error) {
	if err != nil {
		panic(errSignal{err})
	}
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(errSignal{err})
	}
	return v
}

func Do[T any](fn func() T) (result T, err error) {
	defer func() {
		if r := recover(); r != nil {
			if sig, ok := r.(errSignal); ok {
				err = sig.err
				var zero T
				result = zero
			} else {
				// Unexpected panic - propagate
				panic(r)
			}
		}
	}()

	result = fn()
	return result, nil
}
