# 邮件告警示例
本例子通过RESTAPI获取当前值，并对当前值于阈值进行比较，如果超出阈值，则发邮件告警。

[代码传送门](tradeUSDT)
## 代码结构解析
- config：配置信息，包括URL、邮件信息、阈值等
- models：模型结构定义
- services：具体的业务逻辑，请求RESTAPI获取相应数据
- utils：工具，包括Http请求封装、发送邮件等
## 代码片段
- utils.go用来学习封装Http Get请求和发送邮件

		package utils

		import (
			"net/http"
			"io/ioutil"
			"strings"
			"net/smtp"
		)
		
		// Http Get请求基础函数, 通过封装Go语言Http请求
		// strUrl: 请求的URL
		// strParams: string类型的请求参数, user=lxz&pwd=lxz
		// return: 请求结果
		func HttpGetRequest(strUrl string, mapParams map[string]string) string {
			httpClient := &http.Client{}
		
			var strRequestUrl string
			if nil == mapParams {
				strRequestUrl = strUrl
			} else {
				strParams := Map2UrlQuery(mapParams)
				strRequestUrl = strUrl + "?" + strParams
			}
		
			// 构建Request, 并且按官方要求添加Http Header
			request, err := http.NewRequest("GET", strRequestUrl, nil)
			if nil != err {
				return err.Error()
			}
			//request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			// 发出请求
			response, err := httpClient.Do(request)
			defer response.Body.Close()
			if nil != err {
				return err.Error()
			}
		
			// 解析响应内容
			body, err := ioutil.ReadAll(response.Body)
			if nil != err {
				return err.Error()
			}
		
			return string(body)
		}
		//发送邮件通用方法
		func SendToMail(user, password, host, to, body, mailtype string) error {
			hp := strings.Split(host, ":")
			auth := smtp.PlainAuth("", user, password, hp[0])
			var contentType string
			if mailtype == "html" {
				contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
			} else {
				contentType = "Content-Type: text/plain" + "; charset=UTF-8"
			}
		
			msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "价格预警\r\n" + contentType + "\r\n\r\n" + body)
			sendTos := strings.Split(to, ";")
			err := smtp.SendMail(host, auth, user, sendTos, msg)
			return err
		}
- market.go用来学习string转json
		
		func GetGateioUsdtCny() string {
			mapParams := make(map[string]string)
			mapParams["symbol"] = "USDT_CNY"
			mapParams["type"] = "push_main_rates"
			jsonKLineReturn := utils.HttpGetRequest(config.Gateio_URL, mapParams)
		
			var dat map[string]interface{}
			json.Unmarshal([]byte(jsonKLineReturn), &dat)
			var sellPrice string
		
			if v, ok := dat["appraised_rates"]; ok {
				appraisedRates := v.(map[string]interface{})
				if ar, ok := appraisedRates["sell_rate"]; ok {
					sellPrice = ar.(string)
				}
			}
			return sellPrice
		}
- tradeUSDT用来学习时间格式化和字符串转数字等

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
			for range time.Tick(config.CYCLE * time.Second) {
				checkPrice()
			}
		}