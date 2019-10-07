/*
 * Copyright (C) 2017-2019 Alibaba Group Holding Limited
 */
package action

import (
	"bosh-alicloud-cpi/alicloud"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

type HasVMMethod struct {
	CallContext
	instances alicloud.InstanceManager
}

func NewHasVMMethod(cc CallContext, instances alicloud.InstanceManager) HasVMMethod {
	return HasVMMethod{cc, instances}
}

func (a HasVMMethod) HasVM(cid apiv1.VMCID) (bool, error) {
	//
	//
	instCid := cid.AsString()
	inst, err := a.instances.GetInstance(instCid)

	if err != nil {
		return false, bosherr.WrapErrorf(err, "Finding VM Failed '%s'", cid)
	}

	if inst != nil {
		return true, nil
	} else {
		return false, nil
	}
}
