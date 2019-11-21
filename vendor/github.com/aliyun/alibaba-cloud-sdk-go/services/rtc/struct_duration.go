package rtc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Duration is a nested struct in rtc response
type Duration struct {
	Sd          float64     `json:"Sd" xml:"Sd"`
	All         float64     `json:"All" xml:"All"`
	Fhd         float64     `json:"Fhd" xml:"Fhd"`
	Audio       float64     `json:"Audio" xml:"Audio"`
	Hd          float64     `json:"Hd" xml:"Hd"`
	PubDuration PubDuration `json:"PubDuration" xml:"PubDuration"`
	SubDuration SubDuration `json:"SubDuration" xml:"SubDuration"`
}
