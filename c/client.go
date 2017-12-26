/*
 * Copyright GoIIoT (https://github.com/goiiot)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// +build cgo lib

package main

// #cgo CFLAGS: -I include
/*
#include "LibMQTTConst.h"
#include "LibMQTTCallback.h"
*/
import "C"

// Publish (topic *C.char, qos C.goiiot_lib_mqtt_qos_t, payload *C.char, payloadSize C.int)
//export Publish
func Publish(topic *C.char, qos C.goiiot_lib_mqtt_qos_t, payload *C.char, payloadSize C.int) {
	// goTopic := C.GoString(topic)
}

func main() {}
