package serialize

import (
	"bufio"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/ugorji/go/codec"
	"os"
)

func Decode(cborBytes []byte) models.Event {
	var event models.Event
	h := codec.CborHandle{}
	dec := codec.NewDecoderBytes(cborBytes, &h)
	err := dec.Decode(&event)
	if err != nil {
		panic("Failed to decode event: " + err.Error())

	}

	return event
}

func Encode(event models.Event) []byte{
	return event.CBOR()
}

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
	evt := models.Event{Created: timestamp, Modified: timestamp, Device: deviceName}
	var readings []models.Reading
	readings = append(readings, models.Reading{Created: timestamp, Modified: timestamp, Device: deviceName, Name: "Reading2", Value: "789"})
	readings = append(readings, models.Reading{Created: timestamp, Modified: timestamp, Device: deviceName, Name: "Reading1", Value: "XYZ"})
	readings = append(readings, models.Reading{Created: timestamp, Modified: timestamp, Device: deviceName, Name: "Reading1", BinaryValue: bytes})
	evt.Readings = readings

	var handle codec.CborHandle
	data = make([]byte, 0, 64)
	enc := codec.NewEncoderBytes(&data, &handle)

	err = enc.Encode(evt)
	return
}
