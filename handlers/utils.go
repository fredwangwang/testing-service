package handlers

func b(a string) []byte {
	return []byte(a + "\n")
}

func force(a interface{}, b error) interface{} {
	if b != nil {
		panic(b)
	}
	return a
}
