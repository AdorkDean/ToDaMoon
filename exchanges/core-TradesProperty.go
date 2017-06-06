package exchanges

import (
	"ToDaMoon/util"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/imkira/go-observer"
)

//TradeSubject 会发布最新的交易数据。
//会通过observer.Property来更新最新的交易数据。
//updateCycleCh 可以修改Property的更新周期
type TradeSubject struct {
	observer.Property
	UpdateCycleCh chan<- time.Duration
}

//TradesProperty 是exchnge的trades监听属性
type TradesProperty map[string]map[string]TradeSubject

func makeProperties(e Exchanger, tdbs TradesDBs, checkCycle time.Duration) TradesProperty {
	tp := make(TradesProperty)
	wg := &sync.WaitGroup{}
	for money, coinDBs := range tdbs {
		tp[money] = make(map[string]TradeSubject)
		for coin, db := range coinDBs {
			wg.Add(1)
			tp[money][coin] = makePropertyAndSaveToDB(e, money, coin, db, checkCycle, wg)
		}
	}

	wg.Wait()
	text := fmt.Sprintln("已经创建了所有相关的监听属性")
	log.Println(text)

	return tp
}

func makePropertyAndSaveToDB(e Exchanger, money, coin string, db *TradesDB, checkCycle time.Duration, wg *sync.WaitGroup) TradeSubject {
	th := Trades{}
	p := observer.NewProperty(th)
	ch := updatePropertyAndSaveToDB(e, money, coin, p, db, checkCycle)
	wg.Done()
	return TradeSubject{
		Property:      p,
		UpdateCycleCh: ch,
	}
}

func updatePropertyAndSaveToDB(e Exchanger, money, coin string, p observer.Property, db *TradesDB, checkCycle time.Duration) chan<- time.Duration {
	maxTid, err := db.MaxTid()
	if err != nil {
		msg := fmt.Sprintf("updatePropertyAndSaveToDB(): 无法获取%s数据库的MaxTid: %s", db.Name(), err)
		log.Fatalln(msg)
	}
	waitCh, wait := util.WaitFunc(checkCycle)

	go func() {
		for {
			th, err := e.Trades(money, coin, maxTid)
			if err != nil {
				msg := fmt.Sprintf("updatePropertyAndSaveToDB(): 获取%s交易所的%s市场的%s的历史交易数据失败：%s\n。。。。5秒后重试", e.Name, money, coin, err)
				log.Println(msg)
				time.Sleep(time.Second * 5)
				continue
			}

			if len(th) > 0 {
				p.Update(th)

				if err := db.Insert(th); err != nil {
					msg := fmt.Sprintf("插入%s交易所的%s市场的%s的历史交易数据失败：%s", e.Name, money, coin, err)
					log.Fatalln(msg)
				}
			}
			wait()
		}
	}()

	return waitCh
}
