package main

import (
	"fmt"
	"github.com/AnthonyMBonafide/edgex-cbor-benchmark/serialize"
	"os"
	"runtime"
	"strconv"
	"time"
)

var SmallFile = "/home/anthony/GoLandProjects/edgex-cbor-benchmark/small.txt"
var MediumFile = "/home/anthony/GoLandProjects/edgex-cbor-benchmark/medium.txt"
var LargeFile = "/home/anthony/GoLandProjects/edgex-cbor-benchmark/large.txt"

func main() {
	argsWithoutProg := os.Args[1:]
	numberOfIterations, err := strconv.ParseInt(argsWithoutProg[0], 10, 64)
	fileArg := argsWithoutProg[1]

	var file string
	switch fileArg {
	case "small":
		file = "small.txt"
	case "medium":
		file = "medium.txt"
	case "large":
		file = "large.txt"
	default:
		panic("Must specify the event size")
	}

	if err != nil {
		panic("Failed to parse in number of iterations, Please enter a valid integer value")
	}

	//e := codecgen.GenerateTestEvent()
	cborBytes, err := serialize.NewBinaryEvent(file)
	if err != nil{
		panic("Error creating binary event: " + err.Error())
	}

	// Metrics starts from here
	startTime := time.Now()

	for i := int64(0); i < numberOfIterations; i++ {
		// 1. simulates getting a CBOR request and serializing into domain object
		se := serialize.Decode(cborBytes)

		// 2. Update the domain object with information only the backed service has
		se.ID = "Something completely different"
		se.Pushed = 9876543

		// 3. Re-encode the data to CBOR
		serialize.Encode(se)
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Process took: %d ns on average to process %d iterations with an Event containing a reading of %d bytes\n", elapsedTime.Nanoseconds()/numberOfIterations, numberOfIterations,len(cborBytes))

	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	fmt.Printf("Allocated Memory: %d\n", rtm.Alloc)
	fmt.Printf("Total Memory: %d\n", rtm.TotalAlloc)
	fmt.Printf("System Memory: %d\n", rtm.Sys)
	fmt.Printf("Memory Allocations: %d\n", rtm.Mallocs)
	fmt.Printf("Memory Frees: %d\n", rtm.Frees)
	fmt.Printf("Heap Objects: %d\n", rtm.HeapObjects)
	fmt.Printf("GC runs: %d\n", rtm.NumGC)
	fmt.Printf("GC Stop the world time: %dms\n", rtm.PauseTotalNs/1000000)

}
