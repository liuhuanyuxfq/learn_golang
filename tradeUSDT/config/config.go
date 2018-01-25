package config

// API KEY
//const (
//	ACCESS_KEY string = ""
//	SECRET_KEY string = ""
//)

// API请求地址, 不要带最后的/
const (
	//URL配置
	Gateio_URL string = "https://gate.io/json_svr/query_push"
	TRADE_URL  string = "https://api.huobi.pro"
	//邮件配置
	MAIL_FROM   = "yourmailaddress@gmail.com"
	MAIL_PASSWD = "yourpasswd"
	MAIL_HOST   = "smtp.gmail.com:465"
	MAIL_TOS    = "tomailaddress@gmail.com" //此处可以是邮件列表，以";"分隔，最后不要带分号
	//轮询时间间隔
	CYCLE = 60
	//发邮件的阈值
	THRESHOLD = 7.5
)
