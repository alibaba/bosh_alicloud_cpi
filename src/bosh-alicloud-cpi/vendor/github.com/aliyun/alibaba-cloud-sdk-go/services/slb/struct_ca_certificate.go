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

// CACertificate is a nested struct in slb response
type CACertificate struct {
	CreateTimeStamp     int64                        `json:"CreateTimeStamp" xml:"CreateTimeStamp"`
	StandardType        string                       `json:"StandardType" xml:"StandardType"`
	ExpireTime          string                       `json:"ExpireTime" xml:"ExpireTime"`
	CreateTime          string                       `json:"CreateTime" xml:"CreateTime"`
	EncryptionKeyLength int                          `json:"EncryptionKeyLength" xml:"EncryptionKeyLength"`
	ExpireTimeStamp     int64                        `json:"ExpireTimeStamp" xml:"ExpireTimeStamp"`
	CACertificateId     string                       `json:"CACertificateId" xml:"CACertificateId"`
	RegionId            string                       `json:"RegionId" xml:"RegionId"`
	EncryptionAlgorithm string                       `json:"EncryptionAlgorithm" xml:"EncryptionAlgorithm"`
	Fingerprint         string                       `json:"Fingerprint" xml:"Fingerprint"`
	ResourceGroupId     string                       `json:"ResourceGroupId" xml:"ResourceGroupId"`
	CommonName          string                       `json:"CommonName" xml:"CommonName"`
	CACertificateName   string                       `json:"CACertificateName" xml:"CACertificateName"`
	Tags                TagsInDescribeCACertificates `json:"Tags" xml:"Tags"`
}
