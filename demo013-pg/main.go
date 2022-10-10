package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime" //for goroutine
	"time"
	//"encoding/json"
	//"io/ioutil"
)

var logfile *os.File
var logger *log.Logger

//var datainfo *Datainfo

var r *rand.Rand
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Pgbench struct {
	host     string
	port     int
	userName string
	passWord string
	dname    string
	cpunum   int
	datanum  int
	procnum  int
	db       *sql.DB
}

func GetPgbench() *Pgbench {
	return pgbench
}

var pgbench *Pgbench = nil

/*
type Datainfo struct {
	DataId       string
	ParentCollId string
	FsId         string
	DataName     string
	Size         int64
	Mode         int
	Uid          int
	Gid          int
	CreateTs     int64
	ModifyTs     int64
}

type Collinfo struct {
	CollId       string
	ParentCollId string
	FsId         string
	CollName     string
	Mode         int
	Uid          int
	Gid          int
	CreateTs     int64
	ModifyTs     int64
}
*/
func GetPgDBUrl(host, userName, passWord string, port int, dbName string) string {
	pgdbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, userName, passWord, dbName)
	fmt.Println(pgdbUrl)
	return pgdbUrl
}

func Newpgbench(host string, userName string, passWord string, port int, dbName string, cpunum int, datanum int, procnum int) *Pgbench {

	pgbench := &Pgbench{host, port, userName, passWord, dbName, cpunum, datanum, procnum, nil}

	pgurl := GetPgDBUrl(host, userName, passWord, port, dbName)

	db, err := sql.Open("postgres", pgurl)
	if err != nil {
		os.Exit(1)
	}
	pgbench.db = db

	return pgbench
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var counter int64 = 0

func (pgbench *Pgbench) InsertData(ch chan int) {

	var p_u uuid.UUID
	var err error
	for i := 0; i < pgbench.datanum; i++ {

		timestamp := time.Now().Unix()

		if i%10 == 0 {
			p_u, err = uuid.FromString("00000000-0000-0000-0000-000000000001")
			if err != nil {
				os.Exit(1)
			}
		}

		c_u, err := uuid.NewV4()
		if err != nil {
			os.Exit(1)
		}

		coll_name := RandSeq(4)

		insertSql := "insert into r_coll_main values($1,$2,$3,$4,$5,$6,$7,$8,$9)"
		stmt, err := pgbench.db.Prepare(insertSql)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		_, err = stmt.Exec(c_u, p_u, "c5070f54-7ad0-48de-96ee-8f9a1fc85166", coll_name, 16877, 0, 0, timestamp, timestamp)

		if err != nil {
			//log.Fatal(err)
			logger.Println(err)
		}
		//counter
		counter++
		if counter%1000 == 0 {
			logger.Println(counter)
		}

		//insert data
		if i%10 != 0 {
			//every dir inert 10
			for j := 0; j < 10; j++ {
				d_u, err := uuid.NewV4()
				if err != nil {
					os.Exit(1)
				}

				data_name := RandSeq(5)
				randSize := rand.Intn(1000000)

				insertSql := "insert into r_data_main values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
				stmt, err := pgbench.db.Prepare(insertSql)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}

				_, err = stmt.Exec(d_u, p_u, "c5070f54-7ad0-48de-96ee-8f9a1fc85166", data_name, randSize, 33188, 0, 0, timestamp, timestamp)

				if err != nil {
					//log.Fatal(err)
					logger.Println(err)
				}
				//counter
				counter++
				if counter%1000 == 0 {
					logger.Println(counter)

				}
			}

		}

		p_u = c_u

	}
	ch <- 1

}

func main() {
	var host, userName, passWord, logpath string
	var port int
	var dbname string
	var cpunum, datanum, procnum int

	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	flag.StringVar(&host, "host", "", "The pgdb host")
	flag.StringVar(&userName, "userName", "", "The pgdb username")
	flag.StringVar(&passWord, "passWord", "", "The pgdb password")
	flag.IntVar(&port, "port", 5432, "The pgdb port")
	flag.IntVar(&cpunum, "cpunum", 1, "The cpu number wanna use")
	flag.IntVar(&datanum, "datanum", 10000, "The data count per proc")
	flag.IntVar(&procnum, "procnum", 4, "The proc num ")
	flag.StringVar(&logpath, "logpath", "./log.log", "the log path ")
	flag.StringVar(&dbname, "db", "pgbench", "The postgresql database name for testing")

	flag.Parse()

	//log

	var multi_logfile []io.Writer
	var err1 error
	logfile, err1 := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE, 0666)
	os.Truncate(logpath, 0)
	defer logfile.Close()
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	multi_logfile = []io.Writer{
		logfile,
		os.Stdout,
	}

	logfiles := io.MultiWriter(multi_logfile...)
	logger = log.New(logfiles, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)

	if host != "" {
		logger.Println("=====job start.=====")
		pgbench := Newpgbench(host, userName, passWord, port, dbname, cpunum, datanum, procnum)

		defer pgbench.db.Close()

		chs := make([]chan int, pgbench.procnum)
		runtime.GOMAXPROCS(pgbench.cpunum)
		for i := 0; i < pgbench.procnum; i++ {
			fmt.Println(i)

			chs[i] = make(chan int)

			go pgbench.InsertData(chs[i])

		}

		for _, cha := range chs {
			<-cha

		}

		logger.Println("=====Done.=====")
	} else {
		fmt.Println("Please use -help to check the usage")
	}

}
