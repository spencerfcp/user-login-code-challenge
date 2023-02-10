package ptr

func WithDefault[T any](s *T, defaultVal T) T {
	if s == nil {
		return defaultVal
	}

	return *s
}
