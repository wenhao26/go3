package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	twepoch        = int64(1525705533000)             // 开始时间截
	workeridBits   = uint(10)                         // 机器id所占的位数
	sequenceBits   = uint(12)                         // 序列所占的位数
	workeridMax    = int64(-1 ^ (-1 << workeridBits)) // 支持的最大机器id数量
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) // 毫秒内存列的最大值
	workeridShift  = sequenceBits                     // 机器id左移位数
	timestampShift = sequenceBits + workeridBits      // 时间戳左移位数 22位
)

type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerid  int64
	sequence  int64
}

// 实例化snowflake
func NewSnowFlake(workerid int64) (*Snowflake, error) {
	// 判定workerid是否在合理范围内
	if workerid < 0 || workerid > workeridMax {
		return nil, errors.New("error happen")
	}
	return &Snowflake{
		timestamp: 0,
		workerid:  workerid,
		sequence:  0,
	}, nil
}

func (s *Snowflake) Generate() int64 {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixNano() / 1000000
	// 如果是微妙时间戳、且处于同一秒
	if s.timestamp == now {
		// 确保sequence不会溢出,最大值为4095、保证1ms内最多生成4096个ID值
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前时间小于等于当前时间戳
			for now <= s.timestamp {
				// 死循环等待下一个毫秒值
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}

	// todo: 时钟回拨判断、如果获取到的时间戳比上一个小,则存在时钟回拨状态,需要抛出异常

	s.timestamp = now
	// 时间戳左移22位 | workeid左移12位 | sequence兜底
	r := (now-twepoch)<<timestampShift | (s.workerid << workeridShift) | (s.sequence)
	return r
}

// 创建订单ID
func CreateOrder(num int) {
	// YYYY-MM-DD
	ids := time.Now().Format("20060102")

	// 一年中的第几天
	day := strconv.Itoa(GetDayInYear())
	count := len(day)
	if count < 3 {
		day = strings.Repeat("0", 3-count) + day
	}

	ids += day
	fmt.Println(ids)
}

// 一年里的第几天
func GetDayInYear() int {
	now := time.Now()
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := now.Date()
	m := int(month)
	for i := 0; i < m-1; i++ {
		total = total + arr[i]
	}
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
		total = total + d + 1
	} else {
		total = total + d
	}
	return total
}

func main() {
	// CreateOrder(10)
	s := Snowflake{
		timestamp: time.Now().Unix(),
		workerid:  1688,
		sequence:  1688,
	}
	ret := s.Generate()
	fmt.Println(ret)
}
