package database

import (
	"bytes"
	"fmt"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.Debug = true
	//mysql驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//连接数据库 别名default
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(10.95.136.114:3306)/mini_didi?charset=utf8")
	// create table
	orm.RunSyncdb("default", false, true)

}

func getSql(pid string, did string, status string) string {
	var sql bytes.Buffer
	if pid != "" {
		sql.WriteString(fmt.Sprintf("select id, pid, did, from_addr, to_addr, status, create_time from orders where pid = %s order by create_time desc", pid))
	} else if did != "" {
		sql.WriteString(fmt.Sprintf("select id, pid, did, from_addr, to_addr, status, create_time from orders where did = %s order by create_time desc", did))
	}

	if status != "" {
		sql.WriteString(fmt.Sprintf(" and status = %s", status))
	}

	return sql.String()
}
