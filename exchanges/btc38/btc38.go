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

	b3All, err := btc38.allTicker("cny")
	if err != nil {
		fmt.Println("无法获取btc38的cny市场的全部币的ticker")
	} else {
		fmt.Println("BTC38.com All Coins:")
		for k, v := range b3All {
			fmt.Println(k, *v)
		}
	}

	return btc38
}

func build() {
	//生成一个btc38的实例
	btc38 = instance()

	//执行btc38的各项任务
	btc38.checkNewCoin()
	btc38.watching()
}