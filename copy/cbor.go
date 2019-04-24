package cpy

import (
	"bufio"
	"github.com/ugorji/go/codec"
	"os"
)

func NewBinaryEvent(fileName string) (data []byte, err error) {
	fileName = "/tmp/" + fileName
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	bytes := make([]byte, fileInfo.Size())

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	timestamp := int64(123)
	deviceName := "RandomDevice-2"
	evt := Event{Created: timestamp, Modified: timestamp, Device: deviceName}
	var readings []Reading
	readings = append(readings, Reading{Created: timestamp, Modified: timestamp, Device: deviceName, Name: "Reading2", Value: "789"})
	readings = append(readings, Reading{Created: timestamp, Modified: timestamp, Device: deviceName, Name: "Reading1", Value: "XYZ"})
	readings = append(readings, Reading{Created: timestamp, Modified: timestamp, Device: deviceName, Name: "Reading1", BinaryValue: bytes})
	evt.Readings = readings

	var handle codec.CborHandle
	data = make([]byte, 0, 64)
	enc := codec.NewEncoderBytes(&data, &handle)

	err = enc.Encode(&evt)
	return
}

func Decode(cborBytes []byte) Event {
	var event Event
	h := codec.CborHandle{}
	dec := codec.NewDecoderBytes(cborBytes, &h)
	err := dec.Decode(&event)
	if err != nil {
		panic("Failed to decode event: " + err.Error())

	}

	return event
}

func Encode(event Event) []byte{
	return event.CBOR()
}