package common

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var SF *SnowFlake

type SnowFlake struct {
	startTime             int64 //开始时间的偏移量。
	workerIdBits          uint  //worker_id的长度
	dataCenterIdBits      uint
	sequenceBits          uint
	workerIdLeftShift     uint
	dataCenterIdLeftShift uint
	timestampLeftShift    uint
	sequenceMask          int64
	workerId              int64
	dataCenterId          int64
	sequence              int64
	lastTimestamp         int64 //这里存的是毫秒
	idLock                *sync.Mutex
}

func InitSnowFlake(workerId, dataCenterId int64) error {
	//机房ID 2进制5位  32位减掉1位 31个
	//机器ID  2进制5位  32位减掉1位 31个
	if workerId >= 32 || dataCenterId >= 32 || workerId < 0 || dataCenterId < 0 {
		return errors.New("workerId or dataCenterId over max")
	}
	var baseValue int64 = -1

	SF = &SnowFlake{}
	SF.startTime = 0
	if SF.startTime > time.Now().UnixNano()/1000/1000 {
		return errors.New("startTime is over now")
	}
	if time.Now().UnixNano()/1000/1000-SF.startTime < 1099511627776 {
		return errors.New("startTime distance is too small")
	}
	if time.Now().UnixNano()/1000/1000-SF.startTime > 2199023255551 {
		return errors.New("startTime distance is too big")
	}
	SF.workerIdBits = 5
	SF.dataCenterIdBits = 5
	SF.sequenceBits = 12
	SF.workerIdLeftShift = SF.sequenceBits                                 //12
	SF.dataCenterIdLeftShift = SF.workerIdBits + SF.workerIdLeftShift      //17
	SF.timestampLeftShift = SF.dataCenterIdBits + SF.dataCenterIdLeftShift //22
	SF.sequenceMask = baseValue ^ (baseValue << SF.sequenceBits)
	SF.sequence = 0
	SF.lastTimestamp = -1
	SF.idLock = &sync.Mutex{}
	SF.workerId = workerId
	SF.dataCenterId = dataCenterId
	return nil
}

func (sf *SnowFlake) NextId() (int64, error) {
	sf.idLock.Lock()
	defer sf.idLock.Unlock()
	timestamp := time.Now().UnixNano() / 1000 / 1000
	//时间出现了回拨直接return
	if timestamp < sf.lastTimestamp {
		return -1, fmt.Errorf("Clock moved backwards.  Refusing to generate id for %d milliseconds", sf.lastTimestamp-timestamp)
	}
	//同一毫秒，生成对应的sn
	if timestamp == sf.lastTimestamp {
		sf.sequence += 1 //			sequence = (sequence + 1) & sequenceMask;
		//当某一毫秒的时间，产生的id数 超过4095，系统会进入等待，直到下一毫秒，系统继续产生ID
		if sf.sequence > 4095 {
			//超出单毫秒限制数量，等待下一毫秒
			timestamp = sf.waitNextMillis()
			sf.sequence = 0
		}
	} else {
		sf.sequence = 0
	}
	//更新最后时间
	sf.lastTimestamp = timestamp
	// 这儿就是最核心的二进制位运算操作，生成一个64bit的id
	// 先将当前时间戳左移，放到41 bit那儿。因为时间戳是41位。总长是64位，所以左移22位。高位是0代表正数，所以不是23位。
	//特别注意下：sf.lastTimestamp - sf.startTime 这个的差值必须大于1099511627776（2004/11/4 3:53:47）且小于2199023255551（2039/9/7 23:47:35）。不然可能会出现位数有问题。
	// ；将机房id左移放到5 bit那儿；左移5+12
	// 将机器id左移放到5 bit那儿；左移12
	// 将序号放最后12 bit
	// 最后拼接起来成一个64 bit的二进制数字，转换成10进制就是个long型
	id := ((timestamp - sf.startTime) << sf.timestampLeftShift) |
		(sf.dataCenterId << sf.dataCenterIdLeftShift) |
		(sf.workerId << sf.workerIdLeftShift) |
		sf.sequence
	if id < 0 {
		id = -id
	}
	return id, nil
}

func (sf *SnowFlake) waitNextMillis() int64 {
	timestamp := time.Now().UnixNano() / 1000 / 1000 //当前时间
	fmt.Println("进入了等待", timestamp, sf.lastTimestamp)
	for timestamp <= sf.lastTimestamp {
		fmt.Println("刷新了等等时间", timestamp, sf.lastTimestamp)
		timestamp = time.Now().UnixNano() / 1000 / 1000 //一直刷新到当前时间
	}
	return timestamp
}
