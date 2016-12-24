This is a simple assert library for Golang.

I will just introduce what I concider best practice.

It is a best practice to include this library directly into your namespace via:

    import . "github.com/lleo/go-assert"

That dot imports all the symbols into your namespace. All symbols is only
three symbols: ASSERT, Assert().

You can short circut the calling of Assert() by using the following code:

    _ = ASSERT && Asert(1 == 0, "1 != 0")
	
This really help when the Assert() will be expensive code you only want to calling
when you are running your code with ASSERT enabled.

    _ = ASSERT && func() bool {
		result = expensive_func()
		return Assert( result == "good", "expensive_func() didn't return \"good\"")
	}()

The ASSERT boolean can be set to true by setting the environment variable ASSERT
to "true". For example:

    $ ASSERT=true go test

Otherwise ASSERT will be set to false. Hence for the expression

    _ = ASSERT && Whatever()

The Whatever() function will never be called because Golang uses short circut
logical AND evaluation. Hence for "ASSERT && Whatever()" if ASSERT is false then
the expression MUST be false, so golang won't even bother to evaluate the second
part of the && clause.

