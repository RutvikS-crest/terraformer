package vpc

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

// MetricData is a nested struct in vpc response
type MetricData struct {
	NatGatewayId     string `json:"NatGatewayId" xml:"NatGatewayId"`
	PrivateIpAddress string `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
	Timestamp        int64  `json:"Timestamp" xml:"Timestamp"`
	ActiveSessionNum int64  `json:"ActiveSessionNum" xml:"ActiveSessionNum"`
	NewSessionRate   int64  `json:"NewSessionRate" xml:"NewSessionRate"`
	RxBps            int64  `json:"RxBps" xml:"RxBps"`
	TxBps            int64  `json:"TxBps" xml:"TxBps"`
	RxPps            int64  `json:"RxPps" xml:"RxPps"`
	TxPps            int64  `json:"TxPps" xml:"TxPps"`
}
