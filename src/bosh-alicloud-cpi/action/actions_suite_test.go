/*
 * Copyright (C) 2017-2018 Alibaba Group Holding Limited
 */
package action

import (
	"bosh-alicloud-cpi/alicloud"
	"bosh-alicloud-cpi/mock"
	"os"
	"testing"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestActions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Actions Suite")
}

var caller Caller
var mockContext mock.TestContext

var configForTest = []byte(`{
    "cloud": {
        "plugin": "alicloud",
        "properties": {
            "alicloud": {
                "region": "cn-beijing",
				"availability_zone": "cn-beijing-a",
                "access_key_id": "---",
                "access_key_secret": "---"
            },
            "registry": {
                "user": "registry",
                "password": "registry",
                "protocol": "http",
                "address": "127.0.0.1",
                "port": 25777
            },
            "agent": {
                "ntp": ["0.pool.ntp.org", "1.pool.ntp.org"],
                "mbus": "http://mbus:mbus@0.0.0.0:6868",
                "blobstore": {
                    "provider": "dav",
                    "options": {
                        "endpoint": "http://10.0.0.2:25250",
                        "user": "agent",
                        "password": "agent-password"
                    }
                }
            }
        }
    }
}`)

var _ = BeforeSuite(func() {
	config, err := alicloud.NewConfigFromBytes(configForTest)
	Expect(err).NotTo(HaveOccurred())

	err = config.Validate()
	Expect(err).NotTo(HaveOccurred())

	logger := boshlog.NewWriterLogger(boshlog.LevelWarn, os.Stderr)

	mockContext = mock.NewTestContext(config)
	services := Services{
		Stemcells: mock.NewStemcellManagerMock(mockContext),
		Osses:     mock.NewOssManagerMock(mockContext),
		Instances: mock.NewInstanceManagerMock(mockContext),
		Disks:     mock.NewDiskManagerMock(mockContext),
		Networks:  mock.NewNetworkManagerMock(mockContext),
		Registry:  mock.NewRegistryMock(),
	}

	caller = NewCallerWithServices(config, logger, services)
})
