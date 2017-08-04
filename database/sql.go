package main

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"log"


	"time"
)

var dbConn *sql.DB = nil

func GetConn() (db *sql.DB) {

	dbUser := conf.GetValue("database", "dbUser")
	dbPass := conf.GetValue("database", "dbPass")
	dbHost := conf.GetValue("database", "dbHost")
	dbPort := conf.GetValue("database", "dbPort")
	dbType := conf.GetValue("database", "dbType")
	dbName := conf.GetValue("database", "dbName")

	if dbConn == nil {
		dbStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)

		dbConn, err := sql.Open(dbType, dbStr)
		if err == nil {
			dbConn.SetMaxOpenConns(2000)
			dbConn.SetMaxIdleConns(1000)
			return dbConn
		}
	} else {
		return dbConn
	}
	return nil

}

func DbQuery(sql_text string) ([]map[string]interface{}, error) {

	/*db, err := sql.Open("mysql", "root:password@tcp(10.95.136.114:3306)/mini_didi")
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.Ping()
	checkErr(err)*/

	db := GetConn()

	defer db.Close()

	rows, err := db.Query(sql_text)

	checkErr(err)


	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var tags []map[string]interface{}

	for rows.Next() {
		//将行数据保存到recorscanArgsd字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {

				record[columns[i]] = string(col.([]byte))
			}
		}

		tags = append(tags, record)

		//slice[i] = record

	}

	return tags, err

}

func queryOrdersById(oid string) (map[string]interface{}, error) {

	//sql := fmt.Sprintf("select id, pid, did, from_addr, to_addr, status from orders where id = %s", oid)

	db := GetConn()
	rows, err := db.Query("select id, pid, did, from_addr, to_addr, status, create_time from orders where id = ?", oid)
	defer db.Close()

	//rows, err := db.Query("select id, pid, did, from_addr, to_addr from orders where id = ?", oid)

	//rows, err := db.Query(fmt.Sprintf("select oid, pid, did, from_addr, to_addr from orders where id = %d", oid))
	checkErr(err)


	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	//slice := make([]interface{}, 1000000)
	var tags map[string]interface{}
	for rows.Next() {
		//将行数据保存到recorscanArgsd字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				/*log.Println(col)
				log.Println(columns[i])
				bytes := col.([]byte)
				println(bytes)*/
				switch inst := col.(type){
				case []uint8:
					record[columns[i]] = string(col.([]byte))
				case int64:
					record[columns[i]] = inst

				}
				tags = record

			}
		}

	}

	return tags, err

}

func insert(pid int, from_addr string, to_addr string) (int64, error) {
	/*db, err := sql.Open("mysql", "root:password@tcp(10.95.136.114:3306)/mini_didi")
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.Ping()
	checkErr(err)*/
	db := GetConn()

	stmt, err := db.Prepare("INSERT orders (oid, pid, from_addr, to_addr, status) VALUES (?,?,?,?,0)")

	defer db.Close()
	checkErr(err)
	//stmt, err := db.Prepare("INSERT orders (pid, from_addr, to_addr) VALUES (?,?,?)")
	oid, err := oIDGener.NextId()

	_, err = stmt.Exec(oid, pid, from_addr, to_addr)

	return oid, err
}

func PKOrder(oid int, did int, updateStatus int, expectStatus int) (int, error) {
	/*db, err := sql.Open("mysql", "root:password@tcp(10.95.136.114:3306)/mini_didi")
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.Ping()
	checkErr(err)*/
	db := GetConn()

	stmt, err := db.Prepare("UPDATE orders set status = ?, did = ?, update_time = ? where id = ? and status = ?")

	checkErr(err)
	res, err := stmt.Exec(updateStatus, did, time.Now(), oid, expectStatus)

	affectRows, _ := res.RowsAffected();

	if err != nil || affectRows <= 0 {
		return 0, err;
	}

	return int(affectRows), err

}

func finishedOrder(oid int, did int, updateStatus int, expectStatus int) (int, error) {

	/*db, err := sql.Open("mysql", "root:password@tcp(10.95.136.114:3306)/mini_didi")
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.Ping()
	checkErr(err)*/
	db := GetConn()

	stmt, err := db.Prepare("UPDATE orders set status = ?, update_time = ? where id = ? and status = ? and did = ?")

	checkErr(err)
	res, err := stmt.Exec(updateStatus, time.Now(), oid, expectStatus, did)
	checkErr(err)

	affectRows, _ := res.RowsAffected()

	return int(affectRows), err
}

func checkErr(err error) {
	if err != nil {
		Logger.Println(err.Error())
	}
}

func update(sql_text string) (int64, error) {
	// dbStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
	// db, err := sql.Open(dbType, dbStr)
	db := GetConn()
	//defer db.Close()

	if db == nil {
		log.Println("dbConn is nil")
	}
	stmt, err := db.Prepare(sql_text)
	checkErr(err)
	res, err := stmt.Exec()
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	return num, err

}


//判断用户是否注册
func queryUser(phone string, role string) ([]map[string]interface{}, error) {

	sql := "select * from %s where phone = '%s'"
	if role == "driver" {
		sql = fmt.Sprintf(sql, "driver", phone)
	} else {
		sql = fmt.Sprintf(sql, "passenger", phone)
	}
	res, err := DbQuery(sql)

	return res, err

}

//更新用户Token
func updateToken(phone string, role string, token string) (int64, error) {

	sql := "update %s set token = '%s' where phone = '%s'"
	if role == "driver" {
		sql = fmt.Sprintf(sql, "driver", token, phone)
	} else {
		sql = fmt.Sprintf(sql, "passenger", token, phone)
	}
	res, err := update(sql)

	return res, err

}