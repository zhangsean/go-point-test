# go-point-test

Go point sum cost compare with normal sum.

## Usage

```sh
# build
go build

# help
./go-point-test -h
Usage of ./go-point-test:
  -c int
        Target number for sum testing from 1 to count (default 100)

# test default
./go-point-test
Target number count: 100

Normal sum start:  1615196170516531000
Normal sum result: 5050
Normal sum stop:   1615196170516538000
Normal sum cost:   0.007 ms

Point sum start:  1615196170516549000
Point sum result: 5050
Point sum stop:   1615196170516552000
Point sum cost:   0.003 ms

# test 1000
./go-point-test -c 1000
Target number count: 1000

Normal sum start:  1615196200935763000
Normal sum result: 500500
Normal sum stop:   1615196200935766000
Normal sum cost:   0.003 ms

Point sum start:  1615196200935778000
Point sum result: 500500
Point sum stop:   1615196200935802000
Point sum cost:   0.024 ms

# test 10000
Target number count: 10000

Normal sum start:  1615196226715660000
Normal sum result: 50005000
Normal sum stop:   1615196226715678000
Normal sum cost:   0.018 ms

Point sum start:  1615196226715691000
Point sum result: 50005000
Point sum stop:   1615196226715872000
Point sum cost:   0.181 ms

# test 100000
./go-point-test -c 100000
Target number count: 100000

Normal sum start:  1615196302694484000
Normal sum result: 5000050000
Normal sum stop:   1615196302694540000
Normal sum cost:   0.056 ms

Point sum start:  1615196302694552000
Point sum result: 5000050000
Point sum stop:   1615196302696605000
Point sum cost:   2.053 ms

# test 1000000
./go-point-test -c 1000000
Target number count: 1000000

Normal sum start:  1615196606899236000
Normal sum result: 500000500000
Normal sum stop:   1615196606899748000
Normal sum cost:   0.512 ms

Point sum start:  1615196606899761000
Point sum result: 500000500000
Point sum stop:   1615196606917723000
Point sum cost:   17.962 ms

# test 10000000
./go-point-test -c 10000000
Target number count: 10000000

Normal sum start:  1615196671623129000
Normal sum result: 50000005000000
Normal sum stop:   1615196671627125000
Normal sum cost:   3.996 ms

Point sum start:  1615196671627140000
Point sum result: 50000005000000
Point sum stop:   1615196671785845000
Point sum cost:   158.705 ms

```
