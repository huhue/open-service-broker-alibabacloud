package cdn

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

// SetPathCacheExpiredConfig invokes the cdn.SetPathCacheExpiredConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setpathcacheexpiredconfig.html
func (client *Client) SetPathCacheExpiredConfig(request *SetPathCacheExpiredConfigRequest) (response *SetPathCacheExpiredConfigResponse, err error) {
	response = CreateSetPathCacheExpiredConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetPathCacheExpiredConfigWithChan invokes the cdn.SetPathCacheExpiredConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setpathcacheexpiredconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetPathCacheExpiredConfigWithChan(request *SetPathCacheExpiredConfigRequest) (<-chan *SetPathCacheExpiredConfigResponse, <-chan error) {
	responseChan := make(chan *SetPathCacheExpiredConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetPathCacheExpiredConfig(request)
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

// SetPathCacheExpiredConfigWithCallback invokes the cdn.SetPathCacheExpiredConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setpathcacheexpiredconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetPathCacheExpiredConfigWithCallback(request *SetPathCacheExpiredConfigRequest, callback func(response *SetPathCacheExpiredConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetPathCacheExpiredConfigResponse
		var err error
		defer close(result)
		response, err = client.SetPathCacheExpiredConfig(request)
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

// SetPathCacheExpiredConfigRequest is the request struct for api SetPathCacheExpiredConfig
type SetPathCacheExpiredConfigRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	Weight        string           `position:"Query" name:"Weight"`
	CacheContent  string           `position:"Query" name:"CacheContent"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	TTL           string           `position:"Query" name:"TTL"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// SetPathCacheExpiredConfigResponse is the response struct for api SetPathCacheExpiredConfig
type SetPathCacheExpiredConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetPathCacheExpiredConfigRequest creates a request to invoke SetPathCacheExpiredConfig API
func CreateSetPathCacheExpiredConfigRequest() (request *SetPathCacheExpiredConfigRequest) {
	request = &SetPathCacheExpiredConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetPathCacheExpiredConfig", "", "")
	return
}

// CreateSetPathCacheExpiredConfigResponse creates a response to parse from SetPathCacheExpiredConfig response
func CreateSetPathCacheExpiredConfigResponse() (response *SetPathCacheExpiredConfigResponse) {
	response = &SetPathCacheExpiredConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
