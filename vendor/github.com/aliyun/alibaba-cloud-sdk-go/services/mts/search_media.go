package mts

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

// SearchMedia invokes the mts.SearchMedia API synchronously
// api document: https://help.aliyun.com/api/mts/searchmedia.html
func (client *Client) SearchMedia(request *SearchMediaRequest) (response *SearchMediaResponse, err error) {
	response = CreateSearchMediaResponse()
	err = client.DoAction(request, response)
	return
}

// SearchMediaWithChan invokes the mts.SearchMedia API asynchronously
// api document: https://help.aliyun.com/api/mts/searchmedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchMediaWithChan(request *SearchMediaRequest) (<-chan *SearchMediaResponse, <-chan error) {
	responseChan := make(chan *SearchMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchMedia(request)
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

// SearchMediaWithCallback invokes the mts.SearchMedia API asynchronously
// api document: https://help.aliyun.com/api/mts/searchmedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchMediaWithCallback(request *SearchMediaRequest, callback func(response *SearchMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchMediaResponse
		var err error
		defer close(result)
		response, err = client.SearchMedia(request)
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

// SearchMediaRequest is the request struct for api SearchMedia
type SearchMediaRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Title                string           `position:"Query" name:"Title"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	CateId               string           `position:"Query" name:"CateId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	From                 string           `position:"Query" name:"From"`
	SortBy               string           `position:"Query" name:"SortBy"`
	To                   string           `position:"Query" name:"To"`
	Tag                  string           `position:"Query" name:"Tag"`
	KeyWord              string           `position:"Query" name:"KeyWord"`
}

// SearchMediaResponse is the response struct for api SearchMedia
type SearchMediaResponse struct {
	*responses.BaseResponse
	RequestId  string                 `json:"RequestId" xml:"RequestId"`
	TotalNum   int64                  `json:"TotalNum" xml:"TotalNum"`
	PageNumber int64                  `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64                  `json:"PageSize" xml:"PageSize"`
	MediaList  MediaListInSearchMedia `json:"MediaList" xml:"MediaList"`
}

// CreateSearchMediaRequest creates a request to invoke SearchMedia API
func CreateSearchMediaRequest() (request *SearchMediaRequest) {
	request = &SearchMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SearchMedia", "mts", "openAPI")
	return
}

// CreateSearchMediaResponse creates a response to parse from SearchMedia response
func CreateSearchMediaResponse() (response *SearchMediaResponse) {
	response = &SearchMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
