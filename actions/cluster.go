package actions

import (
	"encoding/json"
)

type Cluster struct {
	ActionBase
}

func (this *Cluster) QuorumStatus() ([]byte, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "quorum_status",
	})
	if err != nil {
		return nil, err
	}
	result, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Cluster) Status() ([]byte, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "status",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	result, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (this *Cluster) Version() ([]byte, error) {
	cmdline, err := json.Marshal(map[string]interface{}{
		"prefix": "version",
		"format": "json",
	})
	if err != nil {
		return nil, err
	}
	result, _, err := this.CephConn.Rados.MonCommand(cmdline)
	if err != nil {
		return nil, err
	}
	return result, nil
}
