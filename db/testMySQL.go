package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
	"goDemo/utils"
)

//221上数据库VideoDB,表video_info_dytt的结构
//+------------+---------------+------+-----+---------+-------+
//| Field      | Type          | Null | Key | Default | Extra |
//+------------+---------------+------+-----+---------+-------+
//| id         | int(11)       | NO   | PRI | NULL    |       |
//| md5        | varchar(128)  | YES  | UNI | NULL    |       |
//| name       | varchar(200)  | YES  |     | NULL    |       |
//| decade     | varchar(30)   | YES  |     | NULL    |       |
//| type       | varchar(100)  | YES  |     | NULL    |       |
//| duration   | varchar(30)   | YES  |     | NULL    |       |
//| language   | varchar(30)   | YES  |     | NULL    |       |
//| subtitles  | varchar(100)  | YES  |     | NULL    |       |
//| format     | varchar(20)   | YES  |     | NULL    |       |
//| resolution | varchar(20)   | YES  |     | NULL    |       |
//| director   | varchar(50)   | YES  |     | NULL    |       |
//| actors     | varchar(1000) | YES  |     | NULL    |       |
//| placard    | varchar(200)  | YES  |     | NULL    |       |
//| ftpurl     | varchar(200)  | YES  |     | NULL    |       |
//+------------+---------------+------+-----+---------+-------+

var (
	db_host     = "127.0.0.1"
	db_port     = 3306
	db_username = "root"
	db_password = "sd-9898w"
	db_name     = "VideoDB"
	db_table_name = "video_info_dytt"
)

var (
	db *sql.DB
)

func clearTransaction(tx *sql.Tx){
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil{
		log.Fatalln(err)
	}
}

func GetDBConnect() *sql.DB {
	if db != nil {
		return db
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8&timeout=5s",
		db_username, db_password, db_host, db_port, db_name)
	fmt.Println("GetDBConnect, connStr=", connStr)
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Connect to database failed:", err)
		log.Panicln(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	return db
}

func getVersion(){
	sqlStr := "Select version()"
	dbConn := GetDBConnect()
	stmt, err := dbConn.Prepare(sqlStr)
	if err != nil{
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			values := []interface{}{
				new(interface{}),
			}

			err = rows.Scan(values...)
			if *(values[0].(*interface{})) != nil {
				tmp := string((*(values[0].(*interface{}))).([]uint8))
				fmt.Printf("get version:%s\n",tmp)
			}
		}
	}
}

func AddVideoInfo(mvInfo *utils.MovieInfo){
	sqlstr := fmt.Sprintf("INSERT IGNORE INTO `%s` (`md5`,`name`,`decade`,`type`,`duration`,`language`,`subtitles`,`format`,`resolution`,`director`,`actors`,`placard`,`ftpurl`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", db_table_name, )
	dbConn := GetDBConnect()
	stmt, err := dbConn.Prepare(sqlstr)
	if err != nil{
		log.Println(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Exec(mvInfo.MD5, mvInfo.Name, mvInfo.Decade, mvInfo.VType, mvInfo.Duration, mvInfo.Language,
		mvInfo.Subtitles, mvInfo.Format, mvInfo.Resolution, mvInfo.Director, mvInfo.Actors, mvInfo.Placard, mvInfo.Ftpurl)
	if err != nil {
		log.Println(err.Error())
	}
	i, err := rows.LastInsertId()
	log.Println("insert LastInsertId=",i)
}

func AddVideoInfos(mvInfos []utils.MovieInfo){
	sqlstr := fmt.Sprintf("INSERT IGNORE INTO `%s` (`md5`,`name`,`decade`,`type`,`duration`,`language`,`subtitles`,`format`,`resolution`,`director`,`actors`,`placard`,`ftpurl`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", db_table_name, )
	dbConn := GetDBConnect()
	tx, _ := dbConn.Begin()
	defer clearTransaction(tx)
	for k, _ := range mvInfos{
		mvInfo := mvInfos[k]

		_, err := tx.Exec(sqlstr,mvInfo.MD5, mvInfo.Name, mvInfo.Decade, mvInfo.VType, mvInfo.Duration, mvInfo.Language,
			mvInfo.Subtitles, mvInfo.Format, mvInfo.Resolution, mvInfo.Director, mvInfo.Actors, mvInfo.Placard, mvInfo.Ftpurl)
		if err != nil{
			log.Println(err.Error())
			return
		}
		//i, err := rows.LastInsertId()
		//log.Println("insert LastInsertId=",i)
	}
	tx.Commit()
}

func main() {
	getVersion()
	//添加一条测试
	//movieInfos := utils.GetMovieInfos(0, 1)
	//AddVideoInfo(&movieInfos[0])

	//批量添加
	cnt := utils.GetMovieInfoCnt()
	fmt.Printf("movie info cnt=%d\n",cnt)
	batch := 1000
	times := cnt/batch +1
	for i:= 0; i < times; i++  {
		movieInfos := utils.GetMovieInfos(i*batch, batch)
		for k, _ := range movieInfos{
			mvInfo := movieInfos[k]
			AddVideoInfo(&mvInfo)
		}
		//AddVideoInfos(movieInfos)
	}
}
