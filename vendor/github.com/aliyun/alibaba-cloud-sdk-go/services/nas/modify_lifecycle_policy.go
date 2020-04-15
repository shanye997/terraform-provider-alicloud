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

package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyLifecyclePolicy invokes the nas.ModifyLifecyclePolicy API synchronously
// api document: https://help.aliyun.com/api/nas/modifylifecyclepolicy.html
func (client *Client) ModifyLifecyclePolicy(request *ModifyLifecyclePolicyRequest) (response *ModifyLifecyclePolicyResponse, err error) {
	response = CreateModifyLifecyclePolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLifecyclePolicyWithChan invokes the nas.ModifyLifecyclePolicy API asynchronously
// api document: https://help.aliyun.com/api/nas/modifylifecyclepolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLifecyclePolicyWithChan(request *ModifyLifecyclePolicyRequest) (<-chan *ModifyLifecyclePolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyLifecyclePolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLifecyclePolicy(request)
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

// ModifyLifecyclePolicyWithCallback invokes the nas.ModifyLifecyclePolicy API asynchronously
// api document: https://help.aliyun.com/api/nas/modifylifecyclepolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLifecyclePolicyWithCallback(request *ModifyLifecyclePolicyRequest, callback func(response *ModifyLifecyclePolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLifecyclePolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyLifecyclePolicy(request)
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

// ModifyLifecyclePolicyRequest is the request struct for api ModifyLifecyclePolicy
type ModifyLifecyclePolicyRequest struct {
	*requests.RpcRequest
	FileSystemId        string           `position:"Query" name:"FileSystemId"`
	LifecyclePolicyName string           `position:"Query" name:"LifecyclePolicyName"`
	Path                string           `position:"Query" name:"Path"`
	Recursive           requests.Boolean `position:"Query" name:"Recursive"`
	LifecycleRuleName   string           `position:"Query" name:"LifecycleRuleName"`
	StorageType         string           `position:"Query" name:"StorageType"`
}

// ModifyLifecyclePolicyResponse is the response struct for api ModifyLifecyclePolicy
type ModifyLifecyclePolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyLifecyclePolicyRequest creates a request to invoke ModifyLifecyclePolicy API
func CreateModifyLifecyclePolicyRequest() (request *ModifyLifecyclePolicyRequest) {
	request = &ModifyLifecyclePolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "ModifyLifecyclePolicy", "nas", "openAPI")
	return
}

// CreateModifyLifecyclePolicyResponse creates a response to parse from ModifyLifecyclePolicy response
func CreateModifyLifecyclePolicyResponse() (response *ModifyLifecyclePolicyResponse) {
	response = &ModifyLifecyclePolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
