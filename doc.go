/*
A driver for the go-strm library, that provides the MAP command.

MAP returns rows that result from applying a func to each row in the origin data table stream.

YOU SHOULD ONLY IMPORT THIS PACKAGE USING EITHER A DOT IMPORT OR AN UNDERSCORE IMPORT.

(The package is named "verboten" on purpose to discourage people from importing it by name.)

For example:

	import (
		. "github.com/reiver/strm-map"
	)

Or:

	import (
		_ "github.com/reiver/strm-map"
	)

If you use the 1st style -- the dot import -- then the MAP constant will be made available,
which can be useful to people who use IDEs to write code (for autocomplete) and can help
reduce errors at compile time (by using a pre-defined constant rather than a string literal,
so that you get a compile time error, rather than a runtime error, for a typo).

A more complete example, using the recommended dot import, is:

	package main
	
	import (
		. "github.com/reiver/strm-csv"
		. "github.com/reiver/strm-map"
		. "github.com/reiver/strm-stdout"
	)
	
	func main() {
		Begin(CSV, "table.csv").
			Strm(MAP, func(row []interface{}) []interface{} {
				newRow := make([]interface{}, len(row))

				for i, datum := range row {
					newRow[len(row)-1-i] = datum
				}

				return newRow
			}).
		End(STDOUT, "tsv")
	}

See Also

For more information about go-strm and for a list of other drivers, see:
https://github.com/reiver/go-strm

*/
package verboten
