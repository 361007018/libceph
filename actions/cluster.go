package actions

import (
	"encoding/json"
	"libceph/common"
)

type Cluster struct {
	ActionBase
}

func (this *Cluster) Df() (*common.ResDf, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "df",
		"detail": "detail",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResDf)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Cluster) QuorumStatus() (*common.ResQuorumStatus, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "quorum_status",
	})
	if err != nil {
		return nil, err
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResQuorumStatus)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Cluster) Status() (*common.ResStatus, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "status",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResStatus)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Cluster) Version() (*common.ResVersion, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "version",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	bytes, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	result := new(common.ResVersion)
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, err
	}
	return result, nil
}
