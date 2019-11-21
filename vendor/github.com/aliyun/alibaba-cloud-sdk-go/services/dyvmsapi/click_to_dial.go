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

// ClickToDial invokes the dyvmsapi.ClickToDial API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/clicktodial.html
func (client *Client) ClickToDial(request *ClickToDialRequest) (response *ClickToDialResponse, err error) {
	response = CreateClickToDialResponse()
	err = client.DoAction(request, response)
	return
}

// ClickToDialWithChan invokes the dyvmsapi.ClickToDial API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/clicktodial.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ClickToDialWithChan(request *ClickToDialRequest) (<-chan *ClickToDialResponse, <-chan error) {
	responseChan := make(chan *ClickToDialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ClickToDial(request)
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

// ClickToDialWithCallback invokes the dyvmsapi.ClickToDial API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/clicktodial.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ClickToDialWithCallback(request *ClickToDialRequest, callback func(response *ClickToDialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ClickToDialResponse
		var err error
		defer close(result)
		response, err = client.ClickToDial(request)
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

// ClickToDialRequest is the request struct for api ClickToDial
type ClickToDialRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CallerShowNumber     string           `position:"Query" name:"CallerShowNumber"`
	SessionTimeout       requests.Integer `position:"Query" name:"SessionTimeout"`
	CalledNumber         string           `position:"Query" name:"CalledNumber"`
	CalledShowNumber     string           `position:"Query" name:"CalledShowNumber"`
	AsrFlag              requests.Boolean `position:"Query" name:"AsrFlag"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RecordFlag           requests.Boolean `position:"Query" name:"RecordFlag"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OutId                string           `position:"Query" name:"OutId"`
	AsrModelId           string           `position:"Query" name:"AsrModelId"`
	CallerNumber         string           `position:"Query" name:"CallerNumber"`
}

// ClickToDialResponse is the response struct for api ClickToDial
type ClickToDialResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CallId    string `json:"CallId" xml:"CallId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateClickToDialRequest creates a request to invoke ClickToDial API
func CreateClickToDialRequest() (request *ClickToDialRequest) {
	request = &ClickToDialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "ClickToDial", "", "")
	return
}

// CreateClickToDialResponse creates a response to parse from ClickToDial response
func CreateClickToDialResponse() (response *ClickToDialResponse) {
	response = &ClickToDialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
