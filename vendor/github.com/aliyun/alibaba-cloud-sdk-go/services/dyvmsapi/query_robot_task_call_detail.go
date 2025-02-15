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

// QueryRobotTaskCallDetail invokes the dyvmsapi.QueryRobotTaskCallDetail API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/queryrobottaskcalldetail.html
func (client *Client) QueryRobotTaskCallDetail(request *QueryRobotTaskCallDetailRequest) (response *QueryRobotTaskCallDetailResponse, err error) {
	response = CreateQueryRobotTaskCallDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRobotTaskCallDetailWithChan invokes the dyvmsapi.QueryRobotTaskCallDetail API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/queryrobottaskcalldetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRobotTaskCallDetailWithChan(request *QueryRobotTaskCallDetailRequest) (<-chan *QueryRobotTaskCallDetailResponse, <-chan error) {
	responseChan := make(chan *QueryRobotTaskCallDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRobotTaskCallDetail(request)
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

// QueryRobotTaskCallDetailWithCallback invokes the dyvmsapi.QueryRobotTaskCallDetail API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/queryrobottaskcalldetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRobotTaskCallDetailWithCallback(request *QueryRobotTaskCallDetailRequest, callback func(response *QueryRobotTaskCallDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRobotTaskCallDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryRobotTaskCallDetail(request)
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

// QueryRobotTaskCallDetailRequest is the request struct for api QueryRobotTaskCallDetail
type QueryRobotTaskCallDetailRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Callee               string           `position:"Query" name:"Callee"`
	TaskId               requests.Integer `position:"Query" name:"TaskId"`
	QueryDate            requests.Integer `position:"Query" name:"QueryDate"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryRobotTaskCallDetailResponse is the response struct for api QueryRobotTaskCallDetail
type QueryRobotTaskCallDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateQueryRobotTaskCallDetailRequest creates a request to invoke QueryRobotTaskCallDetail API
func CreateQueryRobotTaskCallDetailRequest() (request *QueryRobotTaskCallDetailRequest) {
	request = &QueryRobotTaskCallDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "QueryRobotTaskCallDetail", "", "")
	return
}

// CreateQueryRobotTaskCallDetailResponse creates a response to parse from QueryRobotTaskCallDetail response
func CreateQueryRobotTaskCallDetailResponse() (response *QueryRobotTaskCallDetailResponse) {
	response = &QueryRobotTaskCallDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
