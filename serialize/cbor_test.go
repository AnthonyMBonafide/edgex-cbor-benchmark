package serialize

import (
	"strconv"
	"testing"
)

var SmallFile = "small.txt"
var MediumFile = "medium.txt"
var LargeFile = "large.txt"

func BenchmarkDecodeSmallEvent(b *testing.B) {
	bytes, _ := NewBinaryEvent(SmallFile)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Decode(bytes)
	}
}
func BenchmarkDecodeMediumEvent(b *testing.B) {
	bytes, _ := NewBinaryEvent(MediumFile)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Decode(bytes)
	}
}
func BenchmarkDecodeLargeEvent(b *testing.B) {
	bytes, _ := NewBinaryEvent(LargeFile)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Decode(bytes)
	}
}

func BenchmarkEncodeSmallEvent(b *testing.B) {
	bytes, _ := NewBinaryEvent(SmallFile)
	e := Decode(bytes)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Encode(e)
	}
}
func BenchmarkEncodeMediumEvent(b *testing.B) {
	bytes, _ := NewBinaryEvent(MediumFile)
	e := Decode(bytes)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Encode(e)
	}
}
func BenchmarkEncodeLargeEvent(b *testing.B) {
	bytes, _ := NewBinaryEvent(LargeFile)
	e := Decode(bytes)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Encode(e)
	}
}

func BenchmarkReEncodeSmallEvent(b *testing.B) {
	bytes, _ := NewBinaryEvent(SmallFile)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		event := Decode(bytes)
		event.ID = "Something New" + strconv.Itoa(i)
		Encode(event)
	}
}

func BenchmarkReEncodeMediumEvent(b *testing.B) {
	bytes, _ := NewBinaryEvent(MediumFile)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		event := Decode(bytes)
		event.ID = "Something New" + strconv.Itoa(i)
		Encode(event)
	}
}

func BenchmarkReEncodeLargeEvent(b *testing.B) {
	bytes, _ := NewBinaryEvent(LargeFile)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		event := Decode(bytes)
		event.ID = "Something New" + strconv.Itoa(i)
		Encode(event)
	}
}
