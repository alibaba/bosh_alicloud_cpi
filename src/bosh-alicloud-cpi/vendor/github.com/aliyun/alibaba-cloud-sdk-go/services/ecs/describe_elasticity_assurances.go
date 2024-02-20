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

// DescribeElasticityAssurances invokes the ecs.DescribeElasticityAssurances API synchronously
func (client *Client) DescribeElasticityAssurances(request *DescribeElasticityAssurancesRequest) (response *DescribeElasticityAssurancesResponse, err error) {
	response = CreateDescribeElasticityAssurancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeElasticityAssurancesWithChan invokes the ecs.DescribeElasticityAssurances API asynchronously
func (client *Client) DescribeElasticityAssurancesWithChan(request *DescribeElasticityAssurancesRequest) (<-chan *DescribeElasticityAssurancesResponse, <-chan error) {
	responseChan := make(chan *DescribeElasticityAssurancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeElasticityAssurances(request)
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

// DescribeElasticityAssurancesWithCallback invokes the ecs.DescribeElasticityAssurances API asynchronously
func (client *Client) DescribeElasticityAssurancesWithCallback(request *DescribeElasticityAssurancesRequest, callback func(response *DescribeElasticityAssurancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeElasticityAssurancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeElasticityAssurances(request)
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

// DescribeElasticityAssurancesRequest is the request struct for api DescribeElasticityAssurances
type DescribeElasticityAssurancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer                   `position:"Query" name:"ResourceOwnerId"`
	Platform              string                             `position:"Query" name:"Platform"`
	ResourceGroupId       string                             `position:"Query" name:"ResourceGroupId"`
	NextToken             string                             `position:"Query" name:"NextToken"`
	InstanceType          string                             `position:"Query" name:"InstanceType"`
	Tag                   *[]DescribeElasticityAssurancesTag `position:"Query" name:"Tag"  type:"Repeated"`
	InstanceChargeType    string                             `position:"Query" name:"InstanceChargeType"`
	ResourceOwnerAccount  string                             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string                             `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer                   `position:"Query" name:"OwnerId"`
	PrivatePoolOptionsIds string                             `position:"Query" name:"PrivatePoolOptions.Ids"`
	MaxResults            requests.Integer                   `position:"Query" name:"MaxResults"`
	ZoneId                string                             `position:"Query" name:"ZoneId"`
	PackageType           string                             `position:"Query" name:"PackageType"`
	Status                string                             `position:"Query" name:"Status"`
}

// DescribeElasticityAssurancesTag is a repeated param struct in DescribeElasticityAssurancesRequest
type DescribeElasticityAssurancesTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeElasticityAssurancesResponse is the response struct for api DescribeElasticityAssurances
type DescribeElasticityAssurancesResponse struct {
	*responses.BaseResponse
	NextToken              string                 `json:"NextToken" xml:"NextToken"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	TotalCount             int                    `json:"TotalCount" xml:"TotalCount"`
	MaxResults             int                    `json:"MaxResults" xml:"MaxResults"`
	ElasticityAssuranceSet ElasticityAssuranceSet `json:"ElasticityAssuranceSet" xml:"ElasticityAssuranceSet"`
}

// CreateDescribeElasticityAssurancesRequest creates a request to invoke DescribeElasticityAssurances API
func CreateDescribeElasticityAssurancesRequest() (request *DescribeElasticityAssurancesRequest) {
	request = &DescribeElasticityAssurancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeElasticityAssurances", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeElasticityAssurancesResponse creates a response to parse from DescribeElasticityAssurances response
func CreateDescribeElasticityAssurancesResponse() (response *DescribeElasticityAssurancesResponse) {
	response = &DescribeElasticityAssurancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
