package collector

import (
	_ "bytes"
	. "dataserver/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "os"
	_ "os/exec"
	_ "path/filepath"
	_ "strings"
	"time"
)

const (
	apikey         string = "c480facf3199797f836e2122a60e3274"                       //开发者key
	servicesUrl    string = "http://apis.baidu.com/showapi_open_bus/weather_showapi" //易源数据服务固定url（show）
	areaidService  string = "/areaid"                                                //获取名字对应cityid服务地址
	addressservice string = "/address"                                               //获取cityid对应天气信息的服务地址
)

var cityidcontent *citycode //内存驻留cityid列表

//获取json数据建模
type citycode struct {
	Cityids []areamodel
}

type areamodel struct {
	Area string
	City []city
}

type city struct {
	Cityname string
	Cityid   string
}

//初始化cityid列表
func confInit() (err error) {
	//fmt.Printf("conf init\n")
	buf, err := ioutil.ReadFile("./conf/cityid.json") //读取文件
	if err != nil {
		return
	}
	cityidcontent = &citycode{}
	err = json.Unmarshal(buf, cityidcontent) //序列化文件为对象
	if err != nil {
		return
	}
	//fmt.Printf("areamodel:%s\n",cityread.Cityids[0].City[0].Cityid)
	return
}

//公开方法  运行采集器
func Run() {
	//fmt.Printf("collector run\n")
	err := confInit() //初始化
	if err != nil {
		fmt.Printf("collector init error:%s\n", err.Error())
	}
	go loop() //开始循环
}

//循环采集
func loop() {
	wcollectortimer := time.NewTimer(time.Hour)
	for {
		select {
		case <-wcollectortimer.C: //1秒后开始第一次采集
			fmt.Printf("collecting\n")
			collect()
			wcollectortimer.Reset(time.Hour * 12) //以后每12小时采集一次
		}
	}
}

//开始采集
func collect() {
	for i := 0; i < len(cityidcontent.Cityids); i++ {
		fmt.Printf("begin area %s \n ", cityidcontent.Cityids[i].Area)
		go collectArea(cityidcontent.Cityids[i])
	}
	//fmt.Printf("alldone:%b\n",done)
}

//采集省级数据
func collectArea(areamodel areamodel) {
	for i := 0; i < len(areamodel.City); i++ {
		fmt.Printf("begin %s \n", areamodel.City[i].Cityname)
		collectCity(areamodel.City[i].Cityid)
		fmt.Printf("end %s \n", areamodel.City[i].Cityname)
	}
}

//采集城市数据
func collectCity(cityid string) {
	v2 := url.Values{}
	v2.Set("areaid", cityid)
	v2.Set("needMoreDay", "0")
	v2.Set("needIndex", "0")

	addresstestUrl := fmt.Sprintf("%s%s?%s", servicesUrl, addressservice, v2.Encode())
	//fmt.Printf("areaidtest:%s\n",addresstestUrl) //输出请求天气预报路径

	req, err := http.NewRequest("GET", addresstestUrl, nil)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	req.Header.Set("apikey", apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("bodyread err:%s\n", err.Error())
	}
	//fmt.Printf("resp body:%s\n",string(body))

	addressResp := &AddressResp{}
	err = json.Unmarshal(body, addressResp)
	if err != nil {
		fmt.Printf("jsonerr %s:%s\n", cityid, err.Error())
		return
	}
	err = UpsertData(&addressResp.Showapi_res_body)
	if err != nil {
		fmt.Printf("upsert %s fail:%s\n", cityid, err.Error())
	}
	fmt.Printf("upsert %s success\n", cityid)
}
