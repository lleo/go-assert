To test this package. You need to run `go test`. But to run it correctly
you need to set the ASSERT env var to "true" eg.

    $ ASSERT=true go test

Logging output is saved in assert_test.log. So you might run `go test` like:

    $ ASSERT=true go test; cat assert_test.log

So much docs for so unnecessary a function.
