package main

import "log"

type Foo interface {
	GetFoo() Foo
}

type FooImpl struct {
	Abc string
	Def string
}

func (fooImpl *FooImpl) GetFoo() Foo {
	return &FooImpl{
		Abc: "ABC",
		Def: "DEF",
	}
}

func main () {
	fooImpl := FooImpl{}

	foo := fooImpl.GetFoo()

	if (foo == nil) {
		log.Fatal("foo is nill")
	} else {
		log.Fatal("Foo is not nil")
	}
}