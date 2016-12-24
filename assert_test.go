package assert_test

import (
	"log"
	"math"
	"os"
	"testing"

	. "github.com/lleo/go-assert"
	"github.com/pkg/errors"
)

func TestMain(m *testing.M) {
	log.SetFlags(log.Lshortfile)

	var logfile, err = os.Create("assert_test.log")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to os.Create(\"assert_test.log\")"))
	}

	log.SetOutput(logfile)

	var xit = m.Run()
	os.Exit(xit)
}

func TestASSERT_IsFalse(t *testing.T) {
	defer func() {
		if ret := recover(); ret != nil {
			log.Printf("TestASSERT_IsFalse: recover() return from panic call: %#v. BAD!", ret)
			t.Fatal("TestASSERT_IsFalse: recover() called") //this retrows a panic
		} else {
			log.Println("TestASSERT_IsFalse: recover() not called. GOOD!")
		}
	}()

	ASSERT = false //make sure Assert() not called due to short circut of &&

	_ = ASSERT && Assert(1 == 0, "%d != %d", 1, 0)
}

func TestASSERT_IsTrue(t *testing.T) {
	defer func() {
		if ret := recover(); ret != nil {
			log.Printf("TestASSERT_IsTrue: recover() return from panic call: %#v. GOOD!", ret)
		} else {
			log.Printf("TestASSERT_IsTrue: recover() not called. GOOD!")
			t.Fatal("TestASSERT_IsTrue: recover() not called.") //this throws a panic
		}
	}()

	ASSERT = true //make sure Assert() is called

	_ = ASSERT && Assert(1 == 0, "%d != %d", 1, 0)
}

func TestASSERT_SetFromEnviroment(t *testing.T) {
	defer func() {
		if ret := recover(); ret != nil {
			log.Printf("TestASSERT_SetFromEnvironment: recover() return from panic call: %#v. GOOD!", ret)
		} else {
			log.Printf("TestASSERT_SetFromEnvironment: Oops! you forgot to set ASSERT=true in the environment before calling `go test`. BAD!")
			t.Fatal("TestASSERT_SetFromEnvironment: Oops! you forgot to set ASSERT=true in the environment before calling `go test`")
		}
	}()

	ASSERT = os.Getenv("ASSERT") == "true"

	_ = ASSERT && Assert(1 == 0, "1 != 0")
}

func isPrime(num uint64) bool {
	hi := uint64(math.Ceil(math.Sqrt(float64(num))))

	for i := uint64(2); i < hi; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func TestASSERT_ExpensiveFunc(t *testing.T) {
	defer func() {
		if ret := recover(); ret != nil {
			log.Printf("TestASSERT_ExpensiveFunc: recover() return from panic call: %#v. BAD!", ret)
			t.Fatal("TestASSERT_ExpensiveFunc: recover() called") //this retrows a panic
		} else {
			log.Println("TestASSERT_ExpensiveFunc: recover() not called. GOOD!")
		}
	}()

	ASSERT = false //make sure Assert() not called due to short circut of &&

	_ = ASSERT && func() bool {
		var bigNum uint64 = math.MaxUint64
		return Assert(isPrime(bigNum), "%d fails isPrime() test", bigNum)
	}()
}
