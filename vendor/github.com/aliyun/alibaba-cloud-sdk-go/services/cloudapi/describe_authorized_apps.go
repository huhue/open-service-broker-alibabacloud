package cloudapi

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

// DescribeAuthorizedApps invokes the cloudapi.DescribeAuthorizedApps API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeauthorizedapps.html
func (client *Client) DescribeAuthorizedApps(request *DescribeAuthorizedAppsRequest) (response *DescribeAuthorizedAppsResponse, err error) {
	response = CreateDescribeAuthorizedAppsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAuthorizedAppsWithChan invokes the cloudapi.DescribeAuthorizedApps API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeauthorizedapps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuthorizedAppsWithChan(request *DescribeAuthorizedAppsRequest) (<-chan *DescribeAuthorizedAppsResponse, <-chan error) {
	responseChan := make(chan *DescribeAuthorizedAppsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAuthorizedApps(request)
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

// DescribeAuthorizedAppsWithCallback invokes the cloudapi.DescribeAuthorizedApps API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeauthorizedapps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAuthorizedAppsWithCallback(request *DescribeAuthorizedAppsRequest, callback func(response *DescribeAuthorizedAppsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAuthorizedAppsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAuthorizedApps(request)
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

// DescribeAuthorizedAppsRequest is the request struct for api DescribeAuthorizedApps
type DescribeAuthorizedAppsRequest struct {
	*requests.RpcRequest
	StageName     string           `position:"Query" name:"StageName"`
	GroupId       string           `position:"Query" name:"GroupId"`
	AppOwnerId    requests.Integer `position:"Query" name:"AppOwnerId"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	AppId         requests.Integer `position:"Query" name:"AppId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	ApiId         string           `position:"Query" name:"ApiId"`
}

// DescribeAuthorizedAppsResponse is the response struct for api DescribeAuthorizedApps
type DescribeAuthorizedAppsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	PageNumber     int            `json:"PageNumber" xml:"PageNumber"`
	AuthorizedApps AuthorizedApps `json:"AuthorizedApps" xml:"AuthorizedApps"`
}

// CreateDescribeAuthorizedAppsRequest creates a request to invoke DescribeAuthorizedApps API
func CreateDescribeAuthorizedAppsRequest() (request *DescribeAuthorizedAppsRequest) {
	request = &DescribeAuthorizedAppsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAuthorizedApps", "apigateway", "openAPI")
	return
}

// CreateDescribeAuthorizedAppsResponse creates a response to parse from DescribeAuthorizedApps response
func CreateDescribeAuthorizedAppsResponse() (response *DescribeAuthorizedAppsResponse) {
	response = &DescribeAuthorizedAppsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
