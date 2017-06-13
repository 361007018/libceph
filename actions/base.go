package actions

import (
	"libceph/conn"
)

type ActionBase struct {
	CephConn *conn.CephConn
}
