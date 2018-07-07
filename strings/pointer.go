package strings

func Pointer(v string) *string {
	if v == "" {
		return nil
	}

	return &v
}
