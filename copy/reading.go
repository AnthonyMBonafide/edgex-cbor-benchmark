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

import "encoding/json"

/*
 * This file is for the Reading model in EdgeX
 * Holds data that was gathered from a device
 *
 *
 * Struct for the Reading object in EdgeX
 */
type Reading struct {
	Id          string `json:"id"`
	Pushed      int64  `json:"pushed"`  // When the data was pushed out of EdgeX (0 - not pushed yet)
	Created     int64  `json:"created"` // When the reading was created
	Origin      int64  `json:"origin"`
	Modified    int64  `json:"modified"`
	Device      string `json:"device"`
	Name        string `json:"name"`
	Value       string `json:"value"`       // Device sensor data value
	BinaryValue []byte `json:"binaryValue"` // Binary data payload
}

/*
 * To String function for Reading Struct
 */
func (r Reading) String() string {
	out, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}

	return string(out)
}
