package dnsknocker

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

// BatchDeleteRr invokes the dnsknocker.BatchDeleteRr API synchronously
// api document: https://help.aliyun.com/api/dnsknocker/batchdeleterr.html
func (client *Client) BatchDeleteRr(request *BatchDeleteRrRequest) (response *BatchDeleteRrResponse, err error) {
	response = CreateBatchDeleteRrResponse()
	err = client.DoAction(request, response)
	return
}

// BatchDeleteRrWithChan invokes the dnsknocker.BatchDeleteRr API asynchronously
// api document: https://help.aliyun.com/api/dnsknocker/batchdeleterr.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteRrWithChan(request *BatchDeleteRrRequest) (<-chan *BatchDeleteRrResponse, <-chan error) {
	responseChan := make(chan *BatchDeleteRrResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchDeleteRr(request)
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

// BatchDeleteRrWithCallback invokes the dnsknocker.BatchDeleteRr API asynchronously
// api document: https://help.aliyun.com/api/dnsknocker/batchdeleterr.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteRrWithCallback(request *BatchDeleteRrRequest, callback func(response *BatchDeleteRrResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchDeleteRrResponse
		var err error
		defer close(result)
		response, err = client.BatchDeleteRr(request)
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

// BatchDeleteRrRequest is the request struct for api BatchDeleteRr
type BatchDeleteRrRequest struct {
	*requests.RpcRequest
	AccessID        string `position:"Body" name:"AccessID"`
	AccessSecret    string `position:"Body" name:"AccessSecret"`
	ResourceRecords string `position:"Body" name:"ResourceRecords"`
	Line            string `position:"Body" name:"Line"`
	ZoneName        string `position:"Body" name:"ZoneName"`
	TransactionId   string `position:"Body" name:"TransactionId"`
	Group           string `position:"Body" name:"Group"`
}

// BatchDeleteRrResponse is the response struct for api BatchDeleteRr
type BatchDeleteRrResponse struct {
	*responses.BaseResponse
	ResultCode    string `json:"ResultCode" xml:"ResultCode"`
	ResultMessage string `json:"ResultMessage" xml:"ResultMessage"`
	Success       bool   `json:"Success" xml:"Success"`
	TransactionId string `json:"TransactionId" xml:"TransactionId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchDeleteRrRequest creates a request to invoke BatchDeleteRr API
func CreateBatchDeleteRrRequest() (request *BatchDeleteRrRequest) {
	request = &BatchDeleteRrRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DnsKnocker", "2019-09-10", "BatchDeleteRr", "dns_knocker", "openAPI")
	return
}

// CreateBatchDeleteRrResponse creates a response to parse from BatchDeleteRr response
func CreateBatchDeleteRrResponse() (response *BatchDeleteRrResponse) {
	response = &BatchDeleteRrResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
