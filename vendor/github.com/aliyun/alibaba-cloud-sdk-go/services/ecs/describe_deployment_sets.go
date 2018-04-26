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

// DescribeDeploymentSets invokes the ecs.DescribeDeploymentSets API synchronously
// api document: https://help.aliyun.com/api/ecs/describedeploymentsets.html
func (client *Client) DescribeDeploymentSets(request *DescribeDeploymentSetsRequest) (response *DescribeDeploymentSetsResponse, err error) {
	response = CreateDescribeDeploymentSetsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDeploymentSetsWithChan invokes the ecs.DescribeDeploymentSets API asynchronously
// api document: https://help.aliyun.com/api/ecs/describedeploymentsets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDeploymentSetsWithChan(request *DescribeDeploymentSetsRequest) (<-chan *DescribeDeploymentSetsResponse, <-chan error) {
	responseChan := make(chan *DescribeDeploymentSetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDeploymentSets(request)
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

// DescribeDeploymentSetsWithCallback invokes the ecs.DescribeDeploymentSets API asynchronously
// api document: https://help.aliyun.com/api/ecs/describedeploymentsets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDeploymentSetsWithCallback(request *DescribeDeploymentSetsRequest, callback func(response *DescribeDeploymentSetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeploymentSetsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDeploymentSets(request)
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

// DescribeDeploymentSetsRequest is the request struct for api DescribeDeploymentSets
type DescribeDeploymentSetsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DeploymentSetIds     string           `position:"Query" name:"DeploymentSetIds"`
	NetworkType          string           `position:"Query" name:"NetworkType"`
	Strategy             string           `position:"Query" name:"Strategy"`
	DeploymentSetName    string           `position:"Query" name:"DeploymentSetName"`
	Granularity          string           `position:"Query" name:"Granularity"`
	Domain               string           `position:"Query" name:"Domain"`
}

// DescribeDeploymentSetsResponse is the response struct for api DescribeDeploymentSets
type DescribeDeploymentSetsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	RegionId       string         `json:"RegionId" xml:"RegionId"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	PageNumber     int            `json:"PageNumber" xml:"PageNumber"`
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	DeploymentSets DeploymentSets `json:"DeploymentSets" xml:"DeploymentSets"`
}

// CreateDescribeDeploymentSetsRequest creates a request to invoke DescribeDeploymentSets API
func CreateDescribeDeploymentSetsRequest() (request *DescribeDeploymentSetsRequest) {
	request = &DescribeDeploymentSetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDeploymentSets", "ecs", "openAPI")
	return
}

// CreateDescribeDeploymentSetsResponse creates a response to parse from DescribeDeploymentSets response
func CreateDescribeDeploymentSetsResponse() (response *DescribeDeploymentSetsResponse) {
	response = &DescribeDeploymentSetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
