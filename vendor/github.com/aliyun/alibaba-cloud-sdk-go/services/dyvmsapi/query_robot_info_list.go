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

// QueryRobotInfoList invokes the dyvmsapi.QueryRobotInfoList API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/queryrobotinfolist.html
func (client *Client) QueryRobotInfoList(request *QueryRobotInfoListRequest) (response *QueryRobotInfoListResponse, err error) {
	response = CreateQueryRobotInfoListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRobotInfoListWithChan invokes the dyvmsapi.QueryRobotInfoList API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/queryrobotinfolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRobotInfoListWithChan(request *QueryRobotInfoListRequest) (<-chan *QueryRobotInfoListResponse, <-chan error) {
	responseChan := make(chan *QueryRobotInfoListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRobotInfoList(request)
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

// QueryRobotInfoListWithCallback invokes the dyvmsapi.QueryRobotInfoList API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/queryrobotinfolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRobotInfoListWithCallback(request *QueryRobotInfoListRequest, callback func(response *QueryRobotInfoListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRobotInfoListResponse
		var err error
		defer close(result)
		response, err = client.QueryRobotInfoList(request)
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

// QueryRobotInfoListRequest is the request struct for api QueryRobotInfoList
type QueryRobotInfoListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AuditStatus          string           `position:"Query" name:"AuditStatus"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QueryRobotInfoListResponse is the response struct for api QueryRobotInfoList
type QueryRobotInfoListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateQueryRobotInfoListRequest creates a request to invoke QueryRobotInfoList API
func CreateQueryRobotInfoListRequest() (request *QueryRobotInfoListRequest) {
	request = &QueryRobotInfoListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "QueryRobotInfoList", "", "")
	return
}

// CreateQueryRobotInfoListResponse creates a response to parse from QueryRobotInfoList response
func CreateQueryRobotInfoListResponse() (response *QueryRobotInfoListResponse) {
	response = &QueryRobotInfoListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
