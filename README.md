## Cryptocurrency Tools

### Binance client
Sometimes I want to known the percentage of my cryto-currency in Binance account, but unfortunately there is not such tools.So I do it with the api of Binance.

#### Config
* create file `.env` in the project root directory,for example:
```$txt
https_proxy="socks5://127.0.0.1:2333"
http_proxy="socks5://127.0.0.1:2333"
binace_apiKey = "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A"
binance_apiSecretKey = "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j"
```
`https_proxy` and `http_proxy` are used for network proxy settings,all of them are not required,you can set in the system environment.

#### Installation
```$shell
cd cmd/binance
go build
./binance
```

#### Usage
```$shell
please input any command: help
commands: ping, server time, account info, signed test, quit
please input any command: server time

serverTime: 2019-03-12 14:13:21
please input any command: account info

assets list: 
asset, value, value of btc, value of USDT, percentage
BNB: 508.009978, 5.653364 BTC, 21651.88 $, 57.762%
BTC: 0.382343, 0.782343 BTC, 2996.31 $, 10.761%
ETH: 2.010854, 0.137693 BTC, 527.35 $, 1.894%
LOOM: 4500.000000, 0.122700 BTC, 469.93 $, 1.688%
BCHABC: 1.500000, 0.082827 BTC, 317.22 $, 1.139%
CMT: 3500.832000, 0.070433 BTC, 269.75 $, 0.969%
AION: 778.300700, 0.060818 BTC, 232.93 $, 0.837%
TRX: 10480.670000, 0.059425 BTC, 227.59 $, 0.817%
ONT: 255.053000, 0.059058 BTC, 226.19 $, 0.812%
SNT: 10000.515100, 0.054903 BTC, 210.27 $, 0.755%
BLZ: 3304.300000, 0.043782 BTC, 167.68 $, 0.602%
USDT: 118.600295, 0.030967 BTC, 118.60 $, 0.426%
POWR: 800.874700, 0.021848 BTC, 83.68 $, 0.301%
LRC: 1000.000000, 0.015630 BTC, 59.86 $, 0.215%
KNC: 250.355826, 0.015397 BTC, 58.97 $, 0.212%
AST: 1500.000000, 0.014085 BTC, 53.94 $, 0.194%
OMG: 41.288306, 0.013903 BTC, 53.25 $, 0.191%
ADA: 1000.000000, 0.012000 BTC, 45.96 $, 0.165%
GTO: 1100.000000, 0.009174 BTC, 35.14 $, 0.126%
RDN: 70.244600, 0.005398 BTC, 20.67 $, 0.074%
CDT: 1933.993100, 0.003868 BTC, 14.81 $, 0.053%
ONG: 3.500000, 0.000444 BTC, 1.70 $, 0.006%
GAS: 0.006500, 0.000004 BTC, 0.02 $, 0.000%
CHAT: 0.789040, 0.000002 BTC, 0.01 $, 0.000%
please input any command: 
```

#### Who?
This is written by Eric Sun, please feel free to contact me via ericsgy@163.com. 