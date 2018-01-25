package main

import (
	"fmt"

	"github.com/liuhuanyuxfq/learn_golang/tradeUSDT/services"
	"github.com/liuhuanyuxfq/learn_golang/tradeUSDT/utils"
	"github.com/liuhuanyuxfq/learn_golang/tradeUSDT/config"
	"strconv"
	"time"
)

func checkPrice() {

	sellPrice := services.GetGateioUsdtCny()
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now, "的价格为：", sellPrice)

	sellNum, _ := strconv.ParseFloat(sellPrice, 64)
	if sellNum > config.THRESHOLD {
		body := `
		<html>
		<body>
		<h3>
		当前价格为：` + sellPrice +
			`,可以出售USDT了！</h3>
			</body>
			</html>
			`
		fmt.Println("send email")
		err := utils.SendToMail(config.MAIL_FROM, config.MAIL_PASSWD, config.MAIL_HOST, config.MAIL_TOS, body, "html")
		if err != nil {
			fmt.Println("Send mail error!")
			fmt.Println(err)
		} else {
			fmt.Println("Send mail success!")
		}
	}

}

func main() {
	checkPrice()
	for range time.Tick(config.CYCLE * time.Second) {
		checkPrice()
	}
}
