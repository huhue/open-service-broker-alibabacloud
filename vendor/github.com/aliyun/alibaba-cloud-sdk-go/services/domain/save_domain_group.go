package domain

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

// SaveDomainGroup invokes the domain.SaveDomainGroup API synchronously
// api document: https://help.aliyun.com/api/domain/savedomaingroup.html
func (client *Client) SaveDomainGroup(request *SaveDomainGroupRequest) (response *SaveDomainGroupResponse, err error) {
	response = CreateSaveDomainGroupResponse()
	err = client.DoAction(request, response)
	return
}

// SaveDomainGroupWithChan invokes the domain.SaveDomainGroup API asynchronously
// api document: https://help.aliyun.com/api/domain/savedomaingroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveDomainGroupWithChan(request *SaveDomainGroupRequest) (<-chan *SaveDomainGroupResponse, <-chan error) {
	responseChan := make(chan *SaveDomainGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveDomainGroup(request)
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

// SaveDomainGroupWithCallback invokes the domain.SaveDomainGroup API asynchronously
// api document: https://help.aliyun.com/api/domain/savedomaingroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveDomainGroupWithCallback(request *SaveDomainGroupRequest, callback func(response *SaveDomainGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveDomainGroupResponse
		var err error
		defer close(result)
		response, err = client.SaveDomainGroup(request)
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

// SaveDomainGroupRequest is the request struct for api SaveDomainGroup
type SaveDomainGroupRequest struct {
	*requests.RpcRequest
	UserClientIp    string           `position:"Query" name:"UserClientIp"`
	DomainGroupName string           `position:"Query" name:"DomainGroupName"`
	Lang            string           `position:"Query" name:"Lang"`
	DomainGroupId   requests.Integer `position:"Query" name:"DomainGroupId"`
}

// SaveDomainGroupResponse is the response struct for api SaveDomainGroup
type SaveDomainGroupResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DomainGroupId     int64  `json:"DomainGroupId" xml:"DomainGroupId"`
	DomainGroupName   string `json:"DomainGroupName" xml:"DomainGroupName"`
	TotalNumber       int    `json:"TotalNumber" xml:"TotalNumber"`
	CreationDate      string `json:"CreationDate" xml:"CreationDate"`
	ModificationDate  string `json:"ModificationDate" xml:"ModificationDate"`
	DomainGroupStatus string `json:"DomainGroupStatus" xml:"DomainGroupStatus"`
	BeingDeleted      bool   `json:"BeingDeleted" xml:"BeingDeleted"`
}

// CreateSaveDomainGroupRequest creates a request to invoke SaveDomainGroup API
func CreateSaveDomainGroupRequest() (request *SaveDomainGroupRequest) {
	request = &SaveDomainGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveDomainGroup", "", "")
	return
}

// CreateSaveDomainGroupResponse creates a response to parse from SaveDomainGroup response
func CreateSaveDomainGroupResponse() (response *SaveDomainGroupResponse) {
	response = &SaveDomainGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
