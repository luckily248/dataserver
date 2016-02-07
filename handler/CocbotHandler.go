package handler

import (
	"errors"
	_ "fmt"
	"strings"
)

type CocbotHandler interface {
	getCommands() []string
	getHelp() string
	handle(text []string) (result string, err error)
}

var mainhandler *MainHandler

func HandlecocText(text string) (reptext string, err error) {
	reccom := strings.Split(text, " ")
	reptext = ""
	//fmt.Printf("reccom:%v\n", reccom)
	mainhandler = &MainHandler{}
	mainhandler.init()
	for _, handler := range mainhandler.allcommands {
		for _, com := range handler.getCommands() {
			//fmt.Printf("com:%s\n", com)
			if strings.EqualFold(reccom[0], com) {
				reptext, err = handler.handle(reccom)
				return
			}
		}
	}
	err = errors.New("command false ,try ?help")
	return
}

type MainHandler struct {
	allcommands []CocbotHandler
}

func (this *MainHandler) init() {
	this.allcommands = make([]CocbotHandler, 0)
	this.allcommands = append(this.allcommands, &HelpHandler{})
	return
}

type HelpHandler struct {
}

func (this *HelpHandler) handle(text []string) (result string, err error) {
	resultslice := make([]string, 0)
	resultslice = append(resultslice, "commands list")
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
	return "?help\n for help"
}
