# go-point-test

Compare the cost of the Go point sum operation with the normal sum operation.

## Usage

```sh
# build
go build

# help
./go-point-test -h
Usage of ./go-point-test:
  -t int
      The target number for sum testing from 1 to (default 100000)

# test default target
./go-point-test
Sum testing from 1 to 100000

Normal direct sum result: 5000050000
Normal direct sum cost:   0.034 ms

Point direct sum result: 5000050000
Point direct sum cost:   0.2 ms

Normal function sum result: 5000050000
Normal function sum cost:   0.06 ms

Point function sum result: 5000050000
Point function sum cost:   2.237 ms

Normal struct sum result: 5000050000
Normal struct sum cost:   0.044 ms

Point struct sum result: 5000050000
Point struct sum cost:   0.188 ms

Normal struct function sum result: 5000050000
Normal struct function sum cost:   0.071 ms

Point struct function sum result: 5000050000
Point struct function sum cost:   1.925 ms

# test target 10000000
./go-point-test -t 10000000
Sum testing from 1 to 10000000

Normal direct sum result: 50000005000000
Normal direct sum cost:   2.855 ms

Point direct sum result: 50000005000000
Point direct sum cost:   20.45 ms

Normal function sum result: 50000005000000
Normal function sum cost:   3.843 ms

Point function sum result: 50000005000000
Point function sum cost:   149.851 ms

Normal struct sum result: 50000005000000
Normal struct sum cost:   2.614 ms

Point struct sum result: 50000005000000
Point struct sum cost:   20.212 ms

Normal struct function sum result: 50000005000000
Normal struct function sum cost:   3.628 ms

Point struct function sum result: 50000005000000
Point struct function sum cost:   152.088 ms

```
