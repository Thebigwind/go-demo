package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	Do()
}

type Data struct {
	Id     int64
	Name   string
	IdCard string
	Status int
}

var ComsumerNum = 5
var DataChanLen = 100

func Do() {
	DataChan := make(chan Data, DataChanLen)

	//生产者
	go Producer(DataChan)

	var wg = sync.WaitGroup{}
	wg.Add(ComsumerNum)
	for i := 0; i < ComsumerNum; i++ {
		go Consumer(DataChan, &wg, i)
	}
	wg.Wait()
	fmt.Println("finish...")
	//监听

}

func Producer(DataChan chan Data) {
	//从数据库查询数据
	queryData := make([]Data, 0)
	//手动构建了1条
	queryData = append(queryData, Data{12, "linklist", "2324", 1})

	for {
		//此处将查询数据写入 queryData。
		//do something

		//如果查询到的数据为0，说明已经查询完，应该关闭Datachan,退出发送的for循环
		if len(queryData) == 0 {
			close(DataChan)
			break
		}

		//将查询到的数据发送到DataChan
		for _, v := range queryData {
			DataChan <- v
		}
		queryData = make([]Data, 0)
	}
	fmt.Println("producer finish...")
}

func Consumer(DataChan chan Data, wg *sync.WaitGroup, index int) {
	defer wg.Done()
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover...")
		}
	}()

	//从DataChan读取，然后做相应逻辑处理，再写回数据库
	timer := time.NewTimer(time.Second * 2)
	defer timer.Stop()

	isStop := false
	for {
		select {
		case <-timer.C:
			fmt.Println("time.Sleep...")
		case data, ok := <-DataChan:
			if !ok {
				//channel已经关闭，并且已经取完了数据，此时可以退出for循环了
				fmt.Println("break..")
				isStop = true
				break
			}
			//可以取到数据，做业务处理写回数据库
			fmt.Printf("data:%v\n", data)
		}
		if isStop {
			break
		}
	}
	fmt.Printf("cosumer %d exit...\n", index)
}
