package client

import (
	"libceph/actions"
	"libceph/conn"
)

var cephClient *CephClient

func GetCephClient() (*CephClient, error) {
	if cephClient == nil {
		InitCephClient("admin")
	}
	return cephClient, nil
}

func InitCephClient(cephXName string) error {
	cephClient = new(CephClient)
	cephClient.CephXName = cephXName
	if err := cephClient.Initial(); err != nil {
		return err
	}
	return nil
}

type CephClient struct {
	CephXName string
	Cluster   *actions.Cluster
	Pool      *actions.Pool
	Mon       *actions.Mon
	Osd       *actions.Osd
	Pg        *actions.Pg
	Auth      *actions.Auth
}

func (this *CephClient) Initial() error {
	cephConn, err := conn.GetCephConn(this.CephXName)
	if err != nil {
		return err
	}
	this.Cluster = &actions.Cluster{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	this.Pool = &actions.Pool{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	this.Mon = &actions.Mon{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	this.Osd = &actions.Osd{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	this.Pg = &actions.Pg{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	this.Auth = &actions.Auth{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	return nil
}
