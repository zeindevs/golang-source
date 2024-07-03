package util

func ErrorPanic(err interface{}) {
	if err != nil {
		panic(err)
	}
}
