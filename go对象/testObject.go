package main

import (
	"encoding/json"
	"fmt"
)

type (
	// TrackingClickData
	TrackingClickData struct {
		Emei          string
		Package       string
		TrackingClick trackingClick
	}

	//存入redis的TrackingClick数据
	trackingClick struct {
		Channel    string `json:"channel"`    //渠道id
		SearchKey  string `json:"searchkey"`  //内部生成广告请求ID{32}
		Vid        string `json:"vid"`        //曝光唯一标识
		GroupID    string `json:"groupid"`    //广告组id
		CreativeID string `json:"creativeid"` //广告创意id
		AdslotID   string `json:"slotid"`     //广告位id
	}
)

func main() {
	// 这里不赋值一个对象也不会出现问题，因为我们定义一个变量，go都会为其赋值为零值
	var trackingClickData TrackingClickData
	trackingClickData.Emei = "dlfsdjlfsdfdjsl"
	trackingClickData.TrackingClick.GroupID = "11111"
	trackingClickData.TrackingClick.Channel = "309"
	trackingClickData.TrackingClick.SearchKey = "searchKey"
	trackingClickData.TrackingClick.Vid = "vid"
	trackingClickData.TrackingClick.CreativeID = "creativeID"
	trackingClickData.TrackingClick.AdslotID = "adslotId"
	bytes, err := json.Marshal(&trackingClickData.TrackingClick)
	fmt.Println(string(bytes), err)
}

//输出结果:
//{"channel":"309","searchkey":"searchKey","vid":"vid","groupid":"11111","creativeid":"creativeID","slotid":"adslotId"} <nil>
