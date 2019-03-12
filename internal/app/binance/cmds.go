package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"cryptocurrency/pkg/config"
	"cryptocurrency/pkg/util"
	"fmt"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type assetOfBtc struct {
	asset string
	value float64
	valueOfBtc float64
}

func init(){
	req.SetTimeout(10 * time.Second)
	//req.Debug = true
}

func newCmdPrompt() {
	io.WriteString(os.Stdout, "please input any command: ")
}

func Ping() {
	resp, err := req.Get(config.EndPoint + config.TestConnectivity)
	util.CheckError(err)
	result := make(map[string]interface{})
	err = resp.ToJSON(&result)
	util.CheckError(err)
	fmt.Println("ping ok")
	newCmdPrompt()
}

func ServerTime() {
	resp, err := req.Get(config.EndPoint + config.ServerTimeUrl)
	util.CheckError(err)
	result := make(map[string]int64)
	err = resp.ToJSON(&result)
	util.CheckError(err)
	dateTime := time.Unix(result["serverTime"] / 1000, result["serverTime"] % 1000)
	fmt.Printf("\nserverTime: %s\n", dateTime.Format(config.TimeYMDHMS))
	newCmdPrompt()
}

func AccountInfo(hashHandler util.UnsignedString) {
	header := req.Header{"X-MBX-APIKEY": hashHandler.ApiKey}
	params := req.Param{
		"timestamp": strconv.FormatInt(time.Now().Unix() * 1000, 10),
		"recvWindow": 50000,
	}
	hashHandler.SetMapData(params)
	params["signature"] = hashHandler.HashEncode()
	resp, err := req.Get(config.EndPoint + config.AccountInfoUrl, params, header)
	util.CheckError(err)
	response := resp.Response()
	defer response.Body.Close()
	r, err := ioutil.ReadAll(response.Body)
	util.CheckError(err)
	result := gjson.Get(string(r), "balances")
	if !result.Exists() {
		fmt.Printf("get account info error: %s\n", string(r))
		newCmdPrompt()
		return
	}

	myAssets := make(map[string]*assetOfBtc)
	c := make(chan assetOfBtc)
	cBtcPrice := make(chan float64)
	fmt.Printf("\nassets list: \nasset, value, value of btc, value of USDT, percentage\n")
	result.ForEach(func(k, v gjson.Result) bool {
		asset, free, locked := v.Get("asset").String(), v.Get("free").Float(), v.Get("locked").Float()
		if free + locked > 0.001 {
			myAssets[asset] = &assetOfBtc{asset: asset, value: free + locked}
			go func(data assetOfBtc) {
				if asset == "BTC" {
					data.valueOfBtc = data.value
					c <- data
				} else if asset == "USDT" {
					resp, err := req.Get(config.EndPoint + config.CurrentAveragePriceUrl, req.Param{"symbol": "BTCUSDT"})
					util.CheckError(err)
					response := resp.Response()
					defer response.Body.Close()
					r, _ := ioutil.ReadAll(response.Body)
					price := gjson.Get(string(r), "price").Float()
					data.valueOfBtc = data.value / price
					c <- data
					cBtcPrice <- price
					close(cBtcPrice)
				} else {
					resp, err := req.Get(config.EndPoint + config.CurrentAveragePriceUrl, req.Param{"symbol": data.asset + "BTC"})
					util.CheckError(err)
					response := resp.Response()
					defer response.Body.Close()
					r, _ := ioutil.ReadAll(response.Body)
					if gjson.Get(string(r), "code").Exists() {
						data.valueOfBtc = 0
					} else {
						price := gjson.Get(string(r), "price").Float()
						data.valueOfBtc = data.value * price
					}
					c <- data
				}
			}(*myAssets[asset])
		}
		return true
	})

	counter := len(myAssets)
	for assetElement := range c {
		myAssets[assetElement.asset].valueOfBtc = assetElement.valueOfBtc
		counter--
		if counter == 0 {
			close(c)
		}
	}
	btcPrice := <- cBtcPrice

	myAssetSlice := []assetOfBtc{}
	var totalValueOfBtc float64 = 0
	for _, v := range myAssets {
		myAssetSlice = append(myAssetSlice, *v)
		totalValueOfBtc += v.valueOfBtc
	}
	sort.SliceStable(myAssetSlice, func(i, j int) bool {
		return myAssetSlice[i].valueOfBtc > myAssetSlice[j].valueOfBtc
	})
	for _, e := range myAssetSlice {
		if e.valueOfBtc > 0 {
			fmt.Printf("%s: %f, %f BTC, %.2f $, %.3f%%\n",
				e.asset, e.value, e.valueOfBtc, e.valueOfBtc * btcPrice, e.valueOfBtc / totalValueOfBtc * 100)
		}
	}
	newCmdPrompt()
}

func SignedTest() {
	hash := util.UnsignedString{
		ApiKey: "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		HashHandler: hmac.New(sha256.New, []byte("NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j")),
	}
	hash.SetData("symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559")
	value := hash.HashEncode()
	correctValue := "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71"
	log.Printf("signed test pass: %t", value == correctValue)
	newCmdPrompt()
}
