package dyvmsapi

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// VoipGetToken invokes the dyvmsapi.VoipGetToken API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/voipgettoken.html
func (client *Client) VoipGetToken(request *VoipGetTokenRequest) (response *VoipGetTokenResponse, err error) {
	response = CreateVoipGetTokenResponse()
	err = client.DoAction(request, response)
	return
}

// VoipGetTokenWithChan invokes the dyvmsapi.VoipGetToken API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/voipgettoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VoipGetTokenWithChan(request *VoipGetTokenRequest) (<-chan *VoipGetTokenResponse, <-chan error) {
	responseChan := make(chan *VoipGetTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VoipGetToken(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// VoipGetTokenWithCallback invokes the dyvmsapi.VoipGetToken API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/voipgettoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VoipGetTokenWithCallback(request *VoipGetTokenRequest, callback func(response *VoipGetTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VoipGetTokenResponse
		var err error
		defer close(result)
		response, err = client.VoipGetToken(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// VoipGetTokenRequest is the request struct for api VoipGetToken
type VoipGetTokenRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VoipId               string           `position:"Query" name:"VoipId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DeviceId             string           `position:"Query" name:"DeviceId"`
	IsCustomAccount      requests.Boolean `position:"Query" name:"IsCustomAccount"`
}

// VoipGetTokenResponse is the response struct for api VoipGetToken
type VoipGetTokenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateVoipGetTokenRequest creates a request to invoke VoipGetToken API
func CreateVoipGetTokenRequest() (request *VoipGetTokenRequest) {
	request = &VoipGetTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "VoipGetToken", "", "")
	return
}

// CreateVoipGetTokenResponse creates a response to parse from VoipGetToken response
func CreateVoipGetTokenResponse() (response *VoipGetTokenResponse) {
	response = &VoipGetTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
