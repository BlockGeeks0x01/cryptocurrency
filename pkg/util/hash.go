package util

import (
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"sort"
	"strconv"
	"strings"
)

type UnsignedString struct {
	ApiKey string
	HashHandler hash.Hash
	data string
	mapData map[string]interface{}
}

func (s *UnsignedString) SetData(data string) {
	s.data = data
}

func (s *UnsignedString) SetMapData(data map[string]interface{}) {
	s.mapData = data
}

func (s UnsignedString) HashEncode() (ret string) {
	var encodeData string
	if s.data != "" {
		encodeData = s.data
	} else {
		dataArray := make([]string, 0) 
		for k, v := range s.mapData {
			switch value := v.(type) {
			case int: dataArray = append(dataArray, k + "=" + strconv.Itoa(value))
			case float64, float32: dataArray = append(dataArray, k + "=" + fmt.Sprintf("%f", value))
			case []byte: dataArray = append(dataArray, k + "=" + string(value))
			case string: dataArray = append(dataArray, k + "=" + value)
			default:
				panic(errors.New("unknown value type"))
			}
		}
		sort.Strings(dataArray)
		encodeData = strings.Join(dataArray, "&")
	}
	s.HashHandler.Write([]byte(encodeData))
	ret = hex.EncodeToString(s.HashHandler.Sum(nil))
	s.HashHandler.Reset()
	return
}
