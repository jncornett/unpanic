package unpanic

// Handle recovers from panics that are the result of an error.
func Handle(err *error) {
	if r := recover(); r != nil {
		var isErr bool
		*err, isErr = r.(error)
		if !isErr {
			panic(r)
		}
	}
}
