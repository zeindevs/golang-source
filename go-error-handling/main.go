package main

func Foo() (int, error) {
	return 666, nil
}

func Bar() (int, error) {
	return 666, nil
}

func Baz() (int, error) {
	return 666, nil
}

func userJustFared() (int, error) {
	return 10, nil
}

func userJustFaredTheirPants() error {
	return nil
}

func changeDaipers() {}

func bar() error {
	return foo()
}

func foo() error {
	return nil
}

func main() {
	x, err := Foo()
	if err != nil {
		return
	}
	y, err := Bar()
	if err != nil {
		return
	}

	_ = y
	_ = x

	smell, _ := userJustFared()
	if err := userJustFaredTheirPants(); err != nil {
		changeDaipers()
	}
	if err := bar(); err != nil {
		// handler
	}
	_ = smell
}
