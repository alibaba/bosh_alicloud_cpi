package slb

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

// CreateLoadBalancerUDPListener invokes the slb.CreateLoadBalancerUDPListener API synchronously
func (client *Client) CreateLoadBalancerUDPListener(request *CreateLoadBalancerUDPListenerRequest) (response *CreateLoadBalancerUDPListenerResponse, err error) {
	response = CreateCreateLoadBalancerUDPListenerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLoadBalancerUDPListenerWithChan invokes the slb.CreateLoadBalancerUDPListener API asynchronously
func (client *Client) CreateLoadBalancerUDPListenerWithChan(request *CreateLoadBalancerUDPListenerRequest) (<-chan *CreateLoadBalancerUDPListenerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerUDPListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancerUDPListener(request)
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

// CreateLoadBalancerUDPListenerWithCallback invokes the slb.CreateLoadBalancerUDPListener API asynchronously
func (client *Client) CreateLoadBalancerUDPListenerWithCallback(request *CreateLoadBalancerUDPListenerRequest, callback func(response *CreateLoadBalancerUDPListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerUDPListenerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancerUDPListener(request)
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

// CreateLoadBalancerUDPListenerRequest is the request struct for api CreateLoadBalancerUDPListener
type CreateLoadBalancerUDPListenerRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer                          `position:"Query" name:"ResourceOwnerId"`
	HealthCheckURI            string                                    `position:"Query" name:"HealthCheckURI"`
	AclStatus                 string                                    `position:"Query" name:"AclStatus"`
	FullNatEnabled            requests.Boolean                          `position:"Query" name:"FullNatEnabled"`
	AclType                   string                                    `position:"Query" name:"AclType"`
	FailoverStrategy          string                                    `position:"Query" name:"FailoverStrategy"`
	PersistenceTimeout        requests.Integer                          `position:"Query" name:"PersistenceTimeout"`
	VpcIds                    string                                    `position:"Query" name:"VpcIds"`
	Tag                       *[]CreateLoadBalancerUDPListenerTag       `position:"Query" name:"Tag"  type:"Repeated"`
	MasterSlaveModeEnabled    requests.Boolean                          `position:"Query" name:"MasterSlaveModeEnabled"`
	VServerGroupId            string                                    `position:"Query" name:"VServerGroupId"`
	AclId                     string                                    `position:"Query" name:"AclId"`
	PortRange                 *[]CreateLoadBalancerUDPListenerPortRange `position:"Query" name:"PortRange"  type:"Repeated"`
	HealthCheckMethod         string                                    `position:"Query" name:"HealthCheckMethod"`
	HealthCheckDomain         string                                    `position:"Query" name:"HealthCheckDomain"`
	OwnerId                   requests.Integer                          `position:"Query" name:"OwnerId"`
	Tags                      string                                    `position:"Query" name:"Tags"`
	LoadBalancerId            string                                    `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroupId  string                                    `position:"Query" name:"MasterSlaveServerGroupId"`
	HealthCheckReq            string                                    `position:"Query" name:"healthCheckReq"`
	BackendServerPort         requests.Integer                          `position:"Query" name:"BackendServerPort"`
	HealthCheckInterval       requests.Integer                          `position:"Query" name:"healthCheckInterval"`
	HealthCheckExp            string                                    `position:"Query" name:"healthCheckExp"`
	FailoverThreshold         requests.Integer                          `position:"Query" name:"FailoverThreshold"`
	ProxyProtocolV2Enabled    requests.Boolean                          `position:"Query" name:"ProxyProtocolV2Enabled"`
	ConnectionDrain           string                                    `position:"Query" name:"ConnectionDrain"`
	HealthCheckSwitch         string                                    `position:"Query" name:"HealthCheckSwitch"`
	AccessKeyId               string                                    `position:"Query" name:"access_key_id"`
	HealthCheckConnectTimeout requests.Integer                          `position:"Query" name:"HealthCheckConnectTimeout"`
	SlaveServerGroupId        string                                    `position:"Query" name:"SlaveServerGroupId"`
	QuicVersion               string                                    `position:"Query" name:"QuicVersion"`
	Description               string                                    `position:"Query" name:"Description"`
	UnhealthyThreshold        requests.Integer                          `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          requests.Integer                          `position:"Query" name:"HealthyThreshold"`
	Scheduler                 string                                    `position:"Query" name:"Scheduler"`
	MaxConnection             requests.Integer                          `position:"Query" name:"MaxConnection"`
	MasterServerGroupId       string                                    `position:"Query" name:"MasterServerGroupId"`
	ListenerPort              requests.Integer                          `position:"Query" name:"ListenerPort"`
	HealthCheckType           string                                    `position:"Query" name:"HealthCheckType"`
	ResourceOwnerAccount      string                                    `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                 requests.Integer                          `position:"Query" name:"Bandwidth"`
	OwnerAccount              string                                    `position:"Query" name:"OwnerAccount"`
	ConnectionDrainTimeout    requests.Integer                          `position:"Query" name:"ConnectionDrainTimeout"`
	HealthCheckConnectPort    requests.Integer                          `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode       string                                    `position:"Query" name:"HealthCheckHttpCode"`
}

// CreateLoadBalancerUDPListenerTag is a repeated param struct in CreateLoadBalancerUDPListenerRequest
type CreateLoadBalancerUDPListenerTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateLoadBalancerUDPListenerPortRange is a repeated param struct in CreateLoadBalancerUDPListenerRequest
type CreateLoadBalancerUDPListenerPortRange struct {
	StartPort string `name:"StartPort"`
	EndPort   string `name:"EndPort"`
}

// CreateLoadBalancerUDPListenerResponse is the response struct for api CreateLoadBalancerUDPListener
type CreateLoadBalancerUDPListenerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateLoadBalancerUDPListenerRequest creates a request to invoke CreateLoadBalancerUDPListener API
func CreateCreateLoadBalancerUDPListenerRequest() (request *CreateLoadBalancerUDPListenerRequest) {
	request = &CreateLoadBalancerUDPListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancerUDPListener", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLoadBalancerUDPListenerResponse creates a response to parse from CreateLoadBalancerUDPListener response
func CreateCreateLoadBalancerUDPListenerResponse() (response *CreateLoadBalancerUDPListenerResponse) {
	response = &CreateLoadBalancerUDPListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
