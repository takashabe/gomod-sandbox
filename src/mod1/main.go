package mod1

import (
	"fmt"

	"github.com/takashabe/gomod-sandbox/src/bar"
	"github.com/takashabe/gomod-sandbox/src/foo"
)

func Hello() {
	fmt.Println("foo: ", foo.Foo{})
	fmt.Println("bar: ", bar.Foo{})
}
