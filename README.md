# edgex-cbor-benchmark
Benchmark tests for determining the impact of decoding and re-encoding data in edgex-go.

## Running benchmark tests

There are different ways which the benchmark tests can be executed:

1. Command line
1. Go benchmark tests

### Setup

First you will need to clone this repo locally with `git clone`. Then execute the `create-data.sh` script to create the necessary data for the benchmark tests to execute, this will create 3 files in the `/tmp` directory which will act as reading data for the test events. Finally install the CLI by executing `go install` at the root of the repo.

### Command line

*NOTE*

You will first need to either clone the repo or install the CLI tool with the `go get` tool.

When executing the benchmark test via the command line you need to specify the number of iterations you would like the tests to execute and the size of the Event which to test(`small`,`medium`,`large`). For example, to run the tests with `1,000,000` iterations with a large event:
 
```bash
 $ edgex-cbor-benchmark 1000000 large
``` 

Example result:
```bash
System statistics before executing tests:
------------------------------------------------------------
Allocated Memory: 25455296
Total Memory: 25572616
System Memory: 72022264
Memory Allocations: 1654
Memory Frees: 368
Heap Objects: 1286
GC runs: 2
GC Stop the world time: 0ms
------------------------------------------------------------
Starting test....
Execution took: 4326 ns on average to process 1000000 iterations with an Event containing a reading of 12583310 bytes
System statistics after executing tests:
------------------------------------------------------------
Allocated Memory: 16075816
Total Memory: 2233577792
System Memory: 72743160
Memory Allocations: 30001699
Memory Frees: 29956842
Heap Objects: 44857
GC runs: 186
GC Stop the world time: 19ms
------------------------------------------------------------

```

*NOTE*

Each iteration will decode an Event struct form CBOR, edit information within the decoded struct, and re-encode the event into CBOR. After the test ahas concluded you will be presented with information regarding elapsed time, and other system information such as memory, CPU, etc.


### Go Benchmark Test

There are 3 benchmark tests in this repo, the most useful is `BenchmarkReEncodeEvent` which will:

1. Decode an Event struct to CBOR
1. Update some basic information within the decoded struct(ID)
1. Encode the updated struct into CBOR

All the tests can be run with the following command at the root directory of the repo:

```bash
$ go test -bench . ./...
```

The results for the previous command will show general information regarding time to execute each operation and allocations on the heap.

For example:

```bash
goos: linux
goarch: amd64
pkg: github.com/AnthonyMBonafide/edgex-cbor-benchmark/serialize
BenchmarkDecodeSmallEvent-8      	 1000000	      1563 ns/op	    1024 B/op	      12 allocs/op
BenchmarkDecodeMediumEvent-8     	 1000000	      1818 ns/op	    1024 B/op	      12 allocs/op
BenchmarkDecodeLargeEvent-8      	 1000000	      1460 ns/op	    1024 B/op	      12 allocs/op
BenchmarkEncodeSmallEvent-8      	 1000000	      1707 ns/op	    1424 B/op	      17 allocs/op
BenchmarkEncodeMediumEvent-8     	 1000000	      1681 ns/op	    1424 B/op	      17 allocs/op
BenchmarkEncodeLargeEvent-8      	 1000000	      1741 ns/op	    1424 B/op	      17 allocs/op
BenchmarkReEncodeSmallEvent-8    	  500000	      3337 ns/op	    2183 B/op	      29 allocs/op
BenchmarkReEncodeMediumEvent-8   	  500000	      4210 ns/op	    2183 B/op	      29 allocs/op
BenchmarkReEncodeLargeEvent-8    	  500000	      3437 ns/op	    2183 B/op	      29 allocs/op
PASS
ok  	github.com/AnthonyMBonafide/edgex-cbor-benchmark/serialize	17.644s
```
For more detailed information regarding the execution flow through the stack you can execute the following command:
```bash
$ go test ./serialize -bench=BenchmarkReEncodeEvent -benchmem -cpuprofile reencode.out
$ go tool pprof -pdf reencode.out > reencode.pdf
```
*NOTE*

GraphViz will need to be installed prior to executing the commands listed below in-order to get a graph visualization of the execution path.

See the [sample output](./docs/sample_reencode.pdf)