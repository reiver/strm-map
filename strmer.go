package verboten


import (
	"github.com/reiver/go-strm/driver"
)


const (
	// MAP is a (string) constant that this Beginner driver
	// is registered under.
	MAP = "MAP"

	defaultLimit = 5
)


func init() {
	strmDriver := newStrmer()

	strmdriver.RegisterStrmer(MAP, strmDriver)
}


type internalStrmer struct{}


func newStrmer() strmdriver.Strmer {
	strmDriver := internalStrmer{

	}

	return &strmDriver
}



func (strmDriver *internalStrmer) Strm(src <-chan []interface{}, dst chan<- []interface{}, args ...interface{}) {

	// Parse args.
	if 1 < len(args) {
//@TODO: Throw something better.
		panic("Too many parameters.")
	}
	if 1 > len(args) {
//@TODO: Throw something better.
		panic("Not enough parameters.")
	}

	arg0 := args[0]

//@TODO: Allow for other func types too.
//       Ex:
//		func([]string)[]string
//		func([]float64)[]float64
//		func([]int64)[]int64
	fn, ok := arg0.(func([]interface{})[]interface{})
	if !ok {
//@TODO: Throw something better.
		panic("Wrong type for map func.")
	}

	// Execute.
	for row := range src {
		newRow := fn(row)

		dst <- newRow
	}
	close(dst)
}
