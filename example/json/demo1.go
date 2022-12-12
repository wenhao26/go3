package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `{"code":0,"message":"Request succeeded.","result":[{"icon":"https://coinsky.s3.us-west-1.amazonaws.com/coin-icon/20220211/bitcoin.png","coin_id":"bitcoin","symbol":"BTC","price":120538.88,"amplitude_24h":2.38,"rank":1,"rank_change":0},{"icon":"https://coinsky.s3.us-west-1.amazonaws.com/coin-icon/20220211/ethereum.png","coin_id":"ethereum","symbol":"ETH","price":9078.25,"amplitude_24h":5.32,"rank":2,"rank_change":21},{"icon":"https://coinsky.s3.us-west-1.amazonaws.com/coin-icon/20220211/tether.png","coin_id":"tether","symbol":"USDT","price":7.16,"amplitude_24h":0.04,"rank":3,"rank_change":17},{"icon":"https://coinsky.s3.us-west-1.amazonaws.com/coin-icon/20220211/tether.png","coin_id":"tether","symbol":"USDT2","price":7.16,"amplitude_24h":0.04,"rank":4,"rank_change":74},{"icon":"https://coinsky.s3.us-west-1.amazonaws.com/coin-icon/20220211/tether.png","coin_id":"tether","symbol":"USDT3","price":7.16,"amplitude_24h":0.04,"rank":5,"rank_change":7}],"request_id":"DEE7F988-3EA7-FF49-170B-187415940B16","response_time":1669875871}`

func main() {
	value := gjson.Get(json, "result.0.coin_id")

	fmt.Println(value.String())
}
