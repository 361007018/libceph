package actions

import (
	"encoding/json"
	"libceph/models"
)

// Cluster ...
type Cluster struct {
	ActionBase
}

// Status ...
func (this *Cluster) Status() (*models.ResponseStatus, error) {
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
	var responseStatus = new(models.ResponseStatus)
	err = json.Unmarshal(result, &responseStatus)
	if err != nil {
		return nil, err
	}
	return responseStatus, nil
}

func (this *Cluster) Version() (*models.ResponseVersion, error) {
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
	var responseVersion = new(models.ResponseVersion)
	err = json.Unmarshal(result, responseVersion)
	if err != nil {
		return nil, err
	}
	return responseVersion, nil
}
