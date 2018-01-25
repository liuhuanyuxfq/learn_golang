package services

import (
	"github.com/liuhuanyuxfq/learn_golang/tradeUSDT/config"
	//"github.com/liuhuanyuxfq/learn_golang/tradeUSDT/models"
	"github.com/liuhuanyuxfq/learn_golang/tradeUSDT/utils"
	"encoding/json"
)

// 获取交易细节信息
// strSymbol: 交易对, btcusdt, bccbtc......
// return: TradeDetailReturn对象
//func GetTradeDetail(strSymbol string) models.TradeDetailReturn {
//	tradeDetailReturn := models.TradeDetailReturn{}
//
//	mapParams := make(map[string]string)
//	mapParams["symbol"] = strSymbol
//
//	strRequestUrl := "/market/trade"
//	strUrl := config.MARKET_URL + strRequestUrl
//
//	jsonTradeDetailReturn := utils.HttpGetRequest(strUrl, mapParams)
//	json.Unmarshal([]byte(jsonTradeDetailReturn), &tradeDetailReturn)
//
//	return tradeDetailReturn
//}

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
		//fmt.Println(appraisedRates)
		//fmt.Println(appraisedRates["sell_rate"])
		if ar, ok := appraisedRates["sell_rate"]; ok {
			sellPrice = ar.(string)
		}
	}
	return sellPrice
}
