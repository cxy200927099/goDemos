package utils

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

const (
	DB_NAME = "./dytt.db"
	MOVIE_INFO_TABLE_NAME = "movie_info"
)

type MovieInfo struct {
	Id         int
	Name       string
	Decade     string
	VType      string
	Duration   string
	Language   string
	Subtitles  string
	Format     string
	Resolution string
	Director   string
	Actors     string
	Placard    string
	Ftpurl     string
	MD5        string
}

func (mvInfo MovieInfo) Print(){
	fmt.Printf("Id=%d MD5=%s Name=%s Resolution=%s \n",
		mvInfo.Id,mvInfo.MD5,mvInfo.Name,mvInfo.Resolution)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetSQLiteDBConnect(dbName string) (*sql.DB, error){
	db, err := sql.Open("sqlite3", dbName)
	checkErr(err)
	return db, err
}

func GetMovieInfoCnt() (int){

	db, err := GetSQLiteDBConnect(DB_NAME)
	defer db.Close()

	sql := fmt.Sprintf("select count(1) from '%s'",MOVIE_INFO_TABLE_NAME)
	rows, err := db.Query(sql)
	checkErr(err)
	var cnt int
	for rows.Next() {
		err = rows.Scan(&cnt)
		checkErr(err)
	}
	rows.Close()

	return cnt
}

func GetMovieInfos(start int, limit int) ([]MovieInfo){
	var movieInfos []MovieInfo
	db, _ := GetSQLiteDBConnect(DB_NAME)
	defer db.Close()

	sql := fmt.Sprintf("select * from %s where Id >%d limit %d", MOVIE_INFO_TABLE_NAME, start, limit)
	fmt.Printf("GetMovieInfos sql=%s\n",sql)
	rows, err := db.Query(sql)
	checkErr(err)
	var index int
	for rows.Next() {
		values := []interface{}{
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
			new(interface{}),
		}

		err = rows.Scan(values...)
		if err != nil {
			fmt.Printf("get %d movieinfo failed! err=%s\n",index, err)
		}

		var tmpMovieInfo MovieInfo
		if *(values[0].(*interface{})) != nil {
			tmp, _ := Int( values[0])
			tmpMovieInfo.Id = tmp
		}
		//string需要先转换为[]uint8数组
		if *(values[1].(*interface{})) != nil {
			tmp := string((*(values[1].(*interface{}))).([]uint8))
			tmpMovieInfo.Name = tmp
		}
		if *(values[2].(*interface{})) != nil {
			tmp := string((*(values[2].(*interface{}))).([]uint8))
			tmpMovieInfo.Decade = tmp
		}
		if *(values[3].(*interface{})) != nil {
			tmp := string((*(values[3].(*interface{}))).([]uint8))
			tmpMovieInfo.VType = tmp
		}
		if *(values[4].(*interface{})) != nil {
			tmp := string((*(values[4].(*interface{}))).([]uint8))
			tmpMovieInfo.Duration = tmp
		}
		if *(values[5].(*interface{})) != nil {
			tmp := string((*(values[5].(*interface{}))).([]uint8))
			tmpMovieInfo.Language = tmp
		}
		if *(values[6].(*interface{})) != nil {
			tmp := string((*(values[6].(*interface{}))).([]uint8))
			tmpMovieInfo.Subtitles = tmp
		}
		if *(values[7].(*interface{})) != nil {
			tmp := string((*(values[7].(*interface{}))).([]uint8))
			tmpMovieInfo.Format = tmp
		}
		if *(values[8].(*interface{})) != nil {
			tmp := string((*(values[8].(*interface{}))).([]uint8))
			tmpMovieInfo.Resolution = tmp
		}
		if *(values[9].(*interface{})) != nil {
			tmp := string((*(values[9].(*interface{}))).([]uint8))
			tmpMovieInfo.Director = tmp
		}
		if *(values[10].(*interface{})) != nil {
			tmp := string((*(values[10].(*interface{}))).([]uint8))
			tmpMovieInfo.Actors = tmp
		}
		if *(values[11].(*interface{})) != nil {
			tmp := string((*(values[11].(*interface{}))).([]uint8))
			tmpMovieInfo.Placard = tmp
		}
		if *(values[12].(*interface{})) != nil {
			tmp := string((*(values[12].(*interface{}))).([]uint8))
			tmpMovieInfo.Ftpurl = tmp
		}
		if *(values[13].(*interface{})) != nil {
			tmp := string((*(values[13].(*interface{}))).([]uint8))
			tmpMovieInfo.MD5 = tmp
		}


		movieInfos = append(movieInfos, tmpMovieInfo)

		index++
	}

	return movieInfos
}