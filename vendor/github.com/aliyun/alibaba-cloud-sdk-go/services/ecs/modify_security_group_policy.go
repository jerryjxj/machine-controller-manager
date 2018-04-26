package ecs

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

// ModifySecurityGroupPolicy invokes the ecs.ModifySecurityGroupPolicy API synchronously
// api document: https://help.aliyun.com/api/ecs/modifysecuritygrouppolicy.html
func (client *Client) ModifySecurityGroupPolicy(request *ModifySecurityGroupPolicyRequest) (response *ModifySecurityGroupPolicyResponse, err error) {
	response = CreateModifySecurityGroupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySecurityGroupPolicyWithChan invokes the ecs.ModifySecurityGroupPolicy API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifysecuritygrouppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySecurityGroupPolicyWithChan(request *ModifySecurityGroupPolicyRequest) (<-chan *ModifySecurityGroupPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifySecurityGroupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySecurityGroupPolicy(request)
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

// ModifySecurityGroupPolicyWithCallback invokes the ecs.ModifySecurityGroupPolicy API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifysecuritygrouppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySecurityGroupPolicyWithCallback(request *ModifySecurityGroupPolicyRequest, callback func(response *ModifySecurityGroupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySecurityGroupPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifySecurityGroupPolicy(request)
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

// ModifySecurityGroupPolicyRequest is the request struct for api ModifySecurityGroupPolicy
type ModifySecurityGroupPolicyRequest struct {
	*requests.RpcRequest
	SecurityGroupId      string           `position:"Query" name:"SecurityGroupId"`
	InnerAccessPolicy    string           `position:"Query" name:"InnerAccessPolicy"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifySecurityGroupPolicyResponse is the response struct for api ModifySecurityGroupPolicy
type ModifySecurityGroupPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySecurityGroupPolicyRequest creates a request to invoke ModifySecurityGroupPolicy API
func CreateModifySecurityGroupPolicyRequest() (request *ModifySecurityGroupPolicyRequest) {
	request = &ModifySecurityGroupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifySecurityGroupPolicy", "ecs", "openAPI")
	return
}

// CreateModifySecurityGroupPolicyResponse creates a response to parse from ModifySecurityGroupPolicy response
func CreateModifySecurityGroupPolicyResponse() (response *ModifySecurityGroupPolicyResponse) {
	response = &ModifySecurityGroupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
