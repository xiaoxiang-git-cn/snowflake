package snowflake

import (
	"sync"
	"time"
)
/*
* Snowflake
* 1. 41位时间截(毫秒级)，注意这是时间戳的差值（当前时间截 - 开始时间截)。可以使用约70年: (1L << 41) / (1000 * 60 * 60 * 24 * 365) = 69
* 2. 10位数据机器位，可以部署在1024个节点
* 3. 12位序列，毫秒内的计数，同一机器，同一时间截并发4096个序号
 */
type snowFlake struct {
	timeStamp 	int64 			// 本次时间戳ms
	workId		int64
	index 		int64					// 序号
	beginTime	int64
	sync.Mutex
}

const (
	//timeBits = 41
	workIdBits = 10
	indexBits = 12
	indexBitsMax = 4096
	beginTime = 1638288000000 //  2021-12-01 00:00:00  (ms)
)

var sf *snowFlake



func (this *snowFlake)uuid() int64 {
	this.Lock()
	defer this.Unlock()
	now := time.Now().UnixMilli()
	if now <= this.timeStamp {
		if this.index >= indexBitsMax {
			this.timeStamp ++
			this.index = 0
		}
	}else{
		this.timeStamp = now
		this.index = 0
	}
	this.index ++
	return this.formatUid()
}

func (this *snowFlake)formatUid() int64 {
	return (this.timeStamp-this.beginTime) << (workIdBits+indexBits) | this.workId <<indexBits | this.index

}
//func init()  {
//	wId := flag.Int64("work_id", 1, "snow_flake_work_id")
//	workId := *wId
//	if workId >= 1024 {
//		panic("snowflake init err")
//	}
//	sf = &snowFlake{workId: workId, beginTime: beginTime}
//}

func Uuid() int64 {
	return sf.uuid()
}

func Init(workId, beginTime int64)  {
	if workId >= 1024 {
		panic("snowflake init err")
	}
	sf = &snowFlake{workId: workId, beginTime: beginTime}
}
