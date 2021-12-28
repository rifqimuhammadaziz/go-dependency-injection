package simple

type Foo struct {
}

// provider foo
func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

// provider bar
func NewBar() *Bar {
	return &Bar{}
}

// create dependency injection without provider
type FooBar struct {
	*Foo
	*Bar
}
