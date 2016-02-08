package handler

import (
	"dataserver/models"
	"errors"
	"fmt"
	_ "fmt"
	"sort"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type CocbotHandler interface {
	getCommands() []string
	getHelp() string
	handle(text []string) (result string, err error)
}

var mainhandler *MainHandler

func HandlecocText(rec models.GMrecModel) (reptext string, err error) {
	reccom := strings.Split(rec.Text, " ")
	reccomfix := []string{}
	for _, s := range reccom {
		if s != "" {
			reccomfix = append(reccomfix, s)
		}
	}
	reptext = ""
	//fmt.Printf("reccom:%v\n", reccom)
	mainhandler = &MainHandler{}
	mainhandler.init(rec)
	for _, handler := range mainhandler.allcommands {
		for _, com := range handler.getCommands() {
			//fmt.Printf("com:%s\n", com)
			if strings.EqualFold(reccomfix[0], com) {
				reptext, err = handler.handle(reccomfix)
				return
			}
		}
	}
	err = errors.New("command false ,try ?help")
	return
}

type MainHandler struct {
	allcommands []CocbotHandler
	rec         models.GMrecModel
}

func (this *MainHandler) init(rec models.GMrecModel) {
	this.allcommands = make([]CocbotHandler, 0)
	this.allcommands = append(this.allcommands, &HelpHandler{})
	this.allcommands = append(this.allcommands, &NewwarHandler{})
	this.allcommands = append(this.allcommands, &ShowwarHandler{})
	this.allcommands = append(this.allcommands, &ScoutHandler{})
	this.rec = rec
	return
}

func (this *MainHandler) getGroupName() string {
	mapforGroupName := map[string]string{"19624531": "luckbot"}
	for k, v := range mapforGroupName {
		//fmt.Printf("gid:%s\n", this.rec.Group_id)
		//fmt.Printf("gn:%s\n", v)
		//fmt.Printf("boolean:%t\n", this.rec.Group_id == k)
		if this.rec.Group_id == k {
			return v
		}
	}
	return ""
}

type HelpHandler struct {
}

func (this *HelpHandler) handle(text []string) (result string, err error) {
	resultslice := make([]string, 0)
	resultslice = append(resultslice, "commands list:")
	resultslice = append(resultslice, "--------------")
	for _, handler := range mainhandler.allcommands {
		resultslice = append(resultslice, handler.getHelp())
	}
	result = strings.Join(resultslice, "\n")
	return
}
func (this *HelpHandler) getCommands() []string {
	return []string{"?help", "?h"}
}
func (this *HelpHandler) getHelp() string {
	return "?help/?h \n for help"
}

type NewwarHandler struct {
}

func (this *NewwarHandler) handle(text []string) (result string, err error) {
	fmt.Printf("len text:%d\n", len(text))
	if len(text) < 2 {
		err = errors.New("i need more info\n" + this.getHelp())
		return
	}

	if len(text) == 2 {
		groupname := mainhandler.getGroupName()
		if groupname == "" {
			err = errors.New("group not found groupid:" + mainhandler.rec.Group_id)
			return
		}
		id, err := models.AddWarData(groupname, text[1])
		if err != nil {
			fmt.Println(err.Error())
			err = errors.New("server error")
		} else {
			result = fmt.Sprintf("done \n %s VS %s \n id:%s", groupname, text[1], id)
		}
	} else {
		id, err := models.AddWarData(text[2], text[1])
		if err != nil {
			fmt.Println(err.Error())
			err = errors.New("server error")
		} else {
			result = fmt.Sprintf("done \n %s VS %s \n id:%s", text[2], text[1], id)
		}
	}
	return

}
func (this *NewwarHandler) getCommands() []string {
	return []string{"?newwar", "?new"}
}
func (this *NewwarHandler) getHelp() string {
	return "?newwar/?new [:enemyclanname] [:yourclanname] \n for new War \n usage: \n ?new enemy \n ?new enemy myclan"
}

type ShowwarHandler struct {
}

func (this *ShowwarHandler) handle(text []string) (result string, err error) {
	groupname := mainhandler.getGroupName()
	if groupname == "" {
		err = errors.New("group not found groupid:" + mainhandler.rec.Group_id)
		return
	}
	content, err := models.GetWarData(groupname)
	if err != nil {
		fmt.Println(err.Error())
		err = errors.New("server error")
		return
	}
	if !content.IsEnable {
		result = "no war being"
		return
	}
	if len(text) == 1 {
		result = fmt.Sprintf("%s VS %s \n id:%s \n %s \n",
			content.TeamA,
			content.TeamB,
			content.GetWarId(),
			content.Timestamp.Format("2006-01-02 15:04:05 +0800"))
		keys := []string{}
		for num, _ := range content.Battles {
			keys = append(keys, num)
		}
		sort.Strings(keys)
		for _, num := range keys {
			result = result + fmt.Sprintf("%s.%s ", num, content.Battles[num].Scoutstate)
			for _, caller := range content.Battles[num].Callers {
				if time.Now().After(caller.Calledtime.Add(6 * time.Hour)) {
					result = result + fmt.Sprintf("%s expried")
				} else {
					expried := caller.Calledtime.Add(6 * time.Hour).Sub(time.Now())
					result = result + fmt.Sprintf("%s %sh%sm ", caller.Callername, expried.Hours(), expried.Minutes())
				}
			}
			result = result + "\n"
		}
	} else {

		for num, battle := range content.Battles {
			if num == text[1] {
				result = fmt.Sprintf("%s.%s ", num, battle.Scoutstate)
				for _, caller := range battle.Callers {
					if time.Now().After(caller.Calledtime.Add(6 * time.Hour)) {
						result = result + fmt.Sprintf("%s expried")
					} else {
						expried := caller.Calledtime.Add(6 * time.Hour).Sub(time.Now())
						result = result + fmt.Sprintf("%s %sh%sm ", caller.Callername, expried.Hours(), expried.Minutes())
					}
				}
			}
		}

	}

	return
}
func (this *ShowwarHandler) getCommands() []string {
	return []string{"?showwar", "?show"}
}
func (this *ShowwarHandler) getHelp() string {
	return "?showwar/?show [:number]\n for show current war condition ,show all or just what u want\n usage:\n ?show\n ?show 1"
}

type ScoutHandler struct {
}

func (this *ScoutHandler) handle(text []string) (result string, err error) {
	if len(text) < 2 {
		err = errors.New("i need more info\n" + this.getHelp())
		return
	}
	groupname := mainhandler.getGroupName()
	if groupname == "" {
		err = errors.New("group not found groupid:" + mainhandler.rec.Group_id)
		return
	}
	content, err := models.GetWarData(groupname)
	if err != nil {
		fmt.Println(err.Error())
		err = errors.New("server error")
		return
	}
	if !content.IsEnable {
		result = "no war being"
		return
	}
	for num, battle := range content.Battles {
		if num == text[1] {
			if battle.Scoutstate == "needscout" || battle.Scoutstate == "scouted" {
				result = fmt.Sprintf("%s already %s", num, battle.Scoutstate)
				return
			}
		}
	}

	err = models.UpdateWarData(content.GetWarId(), bson.M{"$set": bson.M{fmt.Sprintf("battles.%s.scoutstate", text[1]): "needscout"}})
	result = "done\n"
	return
}
func (this *ScoutHandler) getCommands() []string {
	return []string{"?scout"}
}
func (this *ScoutHandler) getHelp() string {
	return "?scout [:number]\n for someone need scout \n usage:\n ?scout 1"
}
