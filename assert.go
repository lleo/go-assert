/*
Package assert implements a very simple assertion library.

The assert package is meant to be imported directly into the client package.

    import . github.com/lleo/go-assert"

So you do not have to use prefix every call to the ASSERT constant and
Assert() function.

When you import the assert library this way you can call Assert() and this way.

    _ = ASSERT && Assert(1 == 0, "1 != 0")

or

    _ = ASSERT && Assert(true == false, "%v != %v", true, false)

So When the environment var ASSERT is set to "true" the Assert() function will
be called. And when ASSERT environment var is unset or set to anything but "true",
then the Assert() function will not be called.

Another variation is:

    _ = ASSERT && func() bool {
        result := expensive_computation()
        return result == "good"
    }
*/
package assert

import (
	"log"
	"os"
)

var ASSERT bool = os.Getenv("ASSERT") == "true"

// The Assert() function tests if the test expression evaluates to false it calls
// panic() with a formatted string and returns true; otherwise it returns false.
func Assert(test bool, fmtstr string, fmtargs ...interface{}) bool {
	if test == false {
		log.Panicf(fmtstr, fmtargs...)
	}
	return false
}
