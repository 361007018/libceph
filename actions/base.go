package actions

import (
	"github.com/361007018/libceph/conn"
)

type ActionBase struct {
	CephConn *conn.CephConn
}
