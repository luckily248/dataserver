package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/weather","description":"get weather information\n"},{"path":"/city","description":"get city citycontent\n"}],"info":{"title":"DataServer API","description":"a dataserver api for data","contact":"luck248@163.com","termsOfServiceUrl":"http://luckily.cc/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/city":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/city","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/citycontent","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get citycontent","responseMessages":[{"code":200,"message":"models.citycontent","responseModel":"citycontent"},{"code":404,"message":"not found","responseModel":""}]}]},{"path":"/:name","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get City by name","parameters":[{"paramType":"path","name":"name","description":"\"cityname\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.City","responseModel":"City"},{"code":403,"message":":name is empty","responseModel":""},{"code":404,"message":"not found","responseModel":""}]}]}],"models":{"Area":{"id":"Area","properties":{"Area":{"type":"string","description":"","format":""},"City":{"type":"array","description":"","items":{"$ref":"City"},"format":""}}},"City":{"id":"City","properties":{"Cityid":{"type":"string","description":"","format":""},"Cityname":{"type":"string","description":"","format":""}}},"citycontent":{"id":"citycontent","properties":{"Areas":{"type":"array","description":"","items":{"$ref":"Area"},"format":""}}}}},"/weather":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/weather","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/:id","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get ShowWeatherModel by id","parameters":[{"paramType":"path","name":"id","description":"\"The key for cityid\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.ShowWeatherModel","responseModel":"ShowWeatherModel"},{"code":404,"message":"not found","responseModel":""},{"code":403,"message":":id is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get All (NOT USE)","type":"","summary":"get ShowWeatherModel (not use)","parameters":[{"paramType":"query","name":"query","description":"\"Filter. e.g. col1:v1,col2:v2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"fields","description":"\"Fields returned. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sortby","description":"\"Sorted-by fields. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"order","description":"\"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"Limit the size of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"Start position of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.ShowWeatherModel","responseModel":"ShowWeatherModel"},{"code":403,"message":"","responseModel":""}]}]}],"models":{"AqiDetail":{"id":"AqiDetail","properties":{"Aqi":{"type":"int","description":"","format":""},"Area":{"type":"string","description":"","format":""},"Co":{"type":"float32","description":"","format":""},"No2":{"type":"int","description":"","format":""},"O3":{"type":"int","description":"","format":""},"O3_8h":{"type":"int","description":"","format":""},"Pm10":{"type":"int","description":"","format":""},"Pm2_5":{"type":"int","description":"","format":""},"Primary_pollutant":{"type":"string","description":"","format":""},"Quality":{"type":"string","description":"","format":""},"So2":{"type":"int","description":"","format":""}}},"BaseDBmodel":{"id":"BaseDBmodel","properties":{"c":{"type":"\u0026{mgo Collection}","description":"","format":""},"db":{"type":"\u0026{mgo Database}","description":"","format":""},"session":{"type":"\u0026{mgo Session}","description":"","format":""}}},"CityInfo":{"id":"CityInfo","properties":{"C1":{"type":"string","description":"","format":""},"C10":{"type":"string","description":"","format":""},"C11":{"type":"string","description":"","format":""},"C12":{"type":"string","description":"","format":""},"C15":{"type":"string","description":"","format":""},"C16":{"type":"string","description":"","format":""},"C17":{"type":"string","description":"","format":""},"C2":{"type":"string","description":"","format":""},"C3":{"type":"string","description":"","format":""},"C4":{"type":"string","description":"","format":""},"C5":{"type":"string","description":"","format":""},"C6":{"type":"string","description":"","format":""},"C7":{"type":"string","description":"","format":""},"C8":{"type":"string","description":"","format":""},"C9":{"type":"string","description":"","format":""},"Latitude":{"type":"float32","description":"","format":""},"Longitude":{"type":"float32","description":"","format":""}}},"ForeCast":{"id":"ForeCast","properties":{"Day":{"type":"string","description":"","format":""},"Day_air_temperature":{"type":"string","description":"","format":""},"Day_weather":{"type":"string","description":"","format":""},"Day_weather_direction":{"type":"string","description":"","format":""},"Day_weather_pic":{"type":"string","description":"","format":""},"Day_wind_power":{"type":"string","description":"","format":""},"Index":{"type":"Index","description":"","format":""},"Night_air_temperature":{"type":"string","description":"","format":""},"Night_weather":{"type":"string","description":"","format":""},"Night_weather_pic":{"type":"string","description":"","format":""},"Night_wind_direction":{"type":"string","description":"","format":""},"Night_wind_power":{"type":"string","description":"","format":""},"Sun_begin_end":{"type":"string","description":"","format":""},"Weekday":{"type":"int","description":"","format":""}}},"Index":{"id":"Index","properties":{"Beauty":{"type":"Indexcontent","description":"","format":""},"Clothes":{"type":"Indexcontent","description":"","format":""},"Cold":{"type":"Indexcontent","description":"","format":""},"Comfort":{"type":"Indexcontent","description":"","format":""},"Glass":{"type":"Indexcontent","description":"","format":""},"Sports":{"type":"Indexcontent","description":"","format":""},"Travel":{"type":"Indexcontent","description":"","format":""},"Uv":{"type":"Indexcontent","description":"","format":""},"Wash_car":{"type":"Indexcontent","description":"","format":""}}},"Indexcontent":{"id":"Indexcontent","properties":{"Desc":{"type":"string","description":"","format":""},"Title":{"type":"string","description":"","format":""}}},"Nowreport":{"id":"Nowreport","properties":{"Aqi":{"type":"\u0026{2980 0x111c00a0 false}","description":"","format":""},"AqiDetail":{"type":"AqiDetail","description":"","format":""},"Sd":{"type":"string","description":"","format":""},"Temperature":{"type":"string","description":"","format":""},"Temperature_time":{"type":"string","description":"","format":""},"Weather":{"type":"string","description":"","format":""},"Weather_pic":{"type":"string","description":"","format":""},"Wind_direction":{"type":"string","description":"","format":""},"Wind_power":{"type":"string","description":"","format":""}}},"ShowWeatherModel":{"id":"ShowWeatherModel","properties":{"CityInfo":{"type":"CityInfo","description":"","format":""},"F1":{"type":"ForeCast","description":"","format":""},"F2":{"type":"ForeCast","description":"","format":""},"F3":{"type":"ForeCast","description":"","format":""},"F4":{"type":"ForeCast","description":"","format":""},"F5":{"type":"ForeCast","description":"","format":""},"F6":{"type":"ForeCast","description":"","format":""},"F7":{"type":"ForeCast","description":"","format":""},"Now":{"type":"Nowreport","description":"","format":""},"Ret_code":{"type":"int","description":"","format":""},"Time":{"type":"string","description":"","format":""}}}}}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	err := json.Unmarshal([]byte(Rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(Subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = BasePath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}