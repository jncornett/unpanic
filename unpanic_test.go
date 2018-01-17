package unpanic

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	err := func() (err error) {
		defer Handle(&err)
		return
	}()
	assert.NoError(t, err)
}

func TestHandle_error(t *testing.T) {
	myErr := errors.New("test")
	err := func() (err error) {
		defer Handle(&err)
		panic(myErr)
	}()
	assert.EqualError(t, err, myErr.Error())
}

func TestHandle_nonError(t *testing.T) {
	myVal := "test"
	assert.PanicsWithValue(t, myVal, func() {
		func() (err error) {
			defer Handle(&err)
			panic(myVal)
		}()
	})
}

func ExampleHandle() {
	someFuncThatMightPanic := func() (err error) {
		defer Handle(&err)
		panic(errors.New("this panic should be recovered"))
	}
	err := someFuncThatMightPanic()
	fmt.Println("returned error:", err)
}
