
/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package cpy

import (
"encoding/json"

"github.com/ugorji/go/codec"
)

// Event represents a single measurable event read from a device
type Event struct {
	ID       string    `json:"id" codec:"id,omitempty"`       // ID uniquely identifies an event, for example a UUID
	Pushed   int64     `json:"pushed" codec:"pushed,omitempty"`   // Pushed is a timestamp indicating when the event was exported. If unexported, the value is zero.
	Device   string    `json:"device" codec:"device,omitempty"`   // Device identifies the source of the event, can be a device name or id. Usually the device name.
	Created  int64     `json:"created" codec:"created,omitempty"`  // Created is a timestamp indicating when the event was created.
	Modified int64     `json:"modified" codec:"modified,omitempty"` // Modified is a timestamp indicating when the event was last modified.
	Origin   int64     `json:"origin" codec:"origin,omitempty"`   // Origin is a timestamp that can communicate the time of the original reading, prior to event creation
	Readings []Reading `json:"readings" codec:"readings,omitempty"` // Readings will contain zero to many entries for the associated readings of a given event.
}

func encodeAsCBOR(e Event) ([]byte, error) {
	var handle codec.CborHandle
	var byteBuffer = make([]byte, 0, 64)
	enc := codec.NewEncoderBytes(&byteBuffer, &handle)

	err := enc.Encode(e)
	if err != nil {
		return []byte{}, err
	}

	return byteBuffer, nil
}

// String provides a JSON representation of the Event as a string
func (e Event) String() string {
	out, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}

	return string(out)
}

// CBOR provides a byte array CBOR-encoded representation of the Event
func (e Event) CBOR() []byte {
	cbor, err := encodeAsCBOR(e)
	if err != nil {
		return []byte{}
	}

	return cbor
}

