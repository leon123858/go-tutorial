# go-tour

my personal notes and exercises from the [Go Tour](https://tour.golang.org/)

## install

```bash
go mod init gotour
go get golang.org/x/tour
go mod tidy
```

## run

```bash
go run cmd/basic.go
go run cmd/flow.go
go run cmd/type.go
go run cmd/interface.go
go run cmd/generic.go
go run cmd/concurrency.go
```

## ptr bench mark

```go
package main

type bigBrother struct {
	x [100000]int
	y [100000]int
	z [100000]int
}

func printBigBrother(b bigBrother) bigBrother {
	b.x[0] = 1
	return b
}

func testPtr() {
	x := bigBrother{}
	printBigBrother(x)
}

func main() {
	testPtr()
}
```

```go
package main

import "testing"

func BenchmarkPrintBigBrother(b *testing.B) {
	x := bigBrother{}
	for i := 0; i < b.N; i++ {
		printBigBrother(x)
	}
}
```