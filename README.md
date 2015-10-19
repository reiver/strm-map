# strm-map

A driver for the **go-strm** Go programming language library, that provides the **MAP** command.

**MAP** returns rows that result from applying a func to each row in the origin data table stream.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/strm-map

[![GoDoc](https://godoc.org/github.com/reiver/strm-map?status.svg)](https://godoc.org/github.com/reiver/strm-map)

## Example
```
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
```

(Note that in that example dot imports were used.)

## See Also

For more information about **go-strm** and for a list of other drivers, see:
https://github.com/reiver/go-strm
