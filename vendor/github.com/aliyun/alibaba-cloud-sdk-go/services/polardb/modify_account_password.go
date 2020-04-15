package polardb

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

// ModifyAccountPassword invokes the polardb.ModifyAccountPassword API synchronously
// api document: https://help.aliyun.com/api/polardb/modifyaccountpassword.html
func (client *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (response *ModifyAccountPasswordResponse, err error) {
	response = CreateModifyAccountPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAccountPasswordWithChan invokes the polardb.ModifyAccountPassword API asynchronously
// api document: https://help.aliyun.com/api/polardb/modifyaccountpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccountPasswordWithChan(request *ModifyAccountPasswordRequest) (<-chan *ModifyAccountPasswordResponse, <-chan error) {
	responseChan := make(chan *ModifyAccountPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAccountPassword(request)
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

// ModifyAccountPasswordWithCallback invokes the polardb.ModifyAccountPassword API asynchronously
// api document: https://help.aliyun.com/api/polardb/modifyaccountpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccountPasswordWithCallback(request *ModifyAccountPasswordRequest, callback func(response *ModifyAccountPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAccountPasswordResponse
		var err error
		defer close(result)
		response, err = client.ModifyAccountPassword(request)
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

// ModifyAccountPasswordRequest is the request struct for api ModifyAccountPassword
type ModifyAccountPasswordRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	NewAccountPassword   string           `position:"Query" name:"NewAccountPassword"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyAccountPasswordResponse is the response struct for api ModifyAccountPassword
type ModifyAccountPasswordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAccountPasswordRequest creates a request to invoke ModifyAccountPassword API
func CreateModifyAccountPasswordRequest() (request *ModifyAccountPasswordRequest) {
	request = &ModifyAccountPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyAccountPassword", "", "")
	return
}

// CreateModifyAccountPasswordResponse creates a response to parse from ModifyAccountPassword response
func CreateModifyAccountPasswordResponse() (response *ModifyAccountPasswordResponse) {
	response = &ModifyAccountPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
