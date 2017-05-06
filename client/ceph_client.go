package client

import (
	"libceph/actions"
	"libceph/common"
)

var cephClient *CephClient

// GetCephConn ...
func GetCephClient() (*CephClient, error) {
	if cephClient == nil {
		cephClient = new(CephClient)
		if err := cephClient.Initial(); err != nil {
			return nil, err
		}
	}
	return cephClient, nil
}

type CephClient struct {
	Cluster *actions.Cluster
	Pool    *actions.Pool
}

func (this *CephClient) Initial() error {
	cephConn, err := common.GetCephConn()
	if err != nil {
		return err
	}
	this.Cluster = new(actions.Cluster)
	this.Cluster.CephConn = cephConn
	this.Pool = new(actions.Pool)
	this.Pool.CephConn = cephConn
	return nil
}
