package btc38

import (
	"ToDaMoon/Interface"
	"ToDaMoon/exchanges"
	"ToDaMoon/pubu"
	"fmt"
	"sync"
)

var once sync.Once
var notify Interface.Notify

//Run 会启动btc38模块
func Run() exchanges.Exchanger {
	notify = pubu.New()
	once.Do(build)

	//以下是测试内容
	fmt.Println("=============================================================")
	b3Ticker, err := btc38.Ticker("btc", "cny")
	if err != nil {
		fmt.Println("BTC38.com BTC Ticker Error:", err)
	} else {
		fmt.Println("BTC38.com BTC Ticker", b3Ticker)
	}

	b3All, err := btc38.allTicker("cny")
	if err != nil {
		fmt.Println("无法获取btc38的cny市场的全部币的ticker")
	} else {
		fmt.Println("BTC38.com All Coins's Ticker:")
		for k, v := range b3All {
			fmt.Println(k, *v)
		}
	}

	b3Depth, err := btc38.Depth("btc", "cny")
	if err != nil {
		fmt.Println("无法获取btc38的cny市场的全部币的ticker")
	} else {
		fmt.Println("BTC38.com btc depth:")
		fmt.Println(b3Depth)

	}

	fmt.Println("=============================================================")
	b3Trades, err := btc38.Trades("btc", "cny", 0)
	if err != nil {
		fmt.Println("无法获取btc38的cny市场的全部币的ticker")
	} else {
		fmt.Println("BTC38.com btc 最新的交易记录:")
		fmt.Println(b3Trades)
	}

	//以上是测试内容

	return btc38
}

func build() {
	//生成一个btc38的实例
	btc38 = instance()

	//执行btc38的各项任务
	btc38.checkNewCoin()
	btc38.watching()
}
