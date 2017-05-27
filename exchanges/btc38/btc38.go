package btc38

import (
	"ToDaMoon/Interface"
	"ToDaMoon/exchanges"
	"ToDaMoon/pubu"
	"sync"
)

var once sync.Once
var notify Interface.Notify

func init() {
	notify = pubu.New()
}

//Run 会启动btc38模块
func Run() exchanges.Exchanger {
	once.Do(build)
	return btc38
}

func build() {
	//生成一个btc38的实例
	btc38 = instance()

	//执行btc38的各项任务
	btc38.checkNewCoin()
	btc38.watching()
}
