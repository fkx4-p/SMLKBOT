package main

import (
	"SMLKBOT/data/botstruct"
	"SMLKBOT/utils/cqfunction"
	"SMLKBOT/utils/smlkshell"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"github.com/tidwall/gjson"
)

type functionFormat func(MsgInfo *botstruct.MsgInfo, BotConfig *botstruct.BotConfig)

func judgeandrun(name string, functionFormat functionFormat, MsgInfo *botstruct.MsgInfo, BotConfig *botstruct.BotConfig) {
	config := gjson.Get(*cqfunction.ConfigFile, "Feature.0").String()
	if gjson.Get(config, name).Bool() {
		go functionFormat(MsgInfo, BotConfig)
	}
}

//MsgHandler converts HTTP Post Body to MsgInfo Struct.
func MsgHandler(raw []byte) (MsgInfo *botstruct.MsgInfo) {
	var mi = new(botstruct.MsgInfo)
	mi.TimeStamp = gjson.GetBytes(raw, "time").Int()
	mi.MsgType = gjson.GetBytes(raw, "message_type").String()
	mi.GroupID = gjson.GetBytes(raw, "group_id").String()
	mi.Message = gjson.GetBytes(raw, "message").String()
	mi.SenderID = gjson.GetBytes(raw, "user_id").String()
	mi.Role = gjson.GetBytes(raw, "sender.role").String()
	str := []byte(strconv.FormatInt(mi.TimeStamp, 10) + mi.MsgType + mi.GroupID + mi.Message + mi.SenderID)
	mi.MD5 = md5.Sum(str)
	return mi
}

//HTTPhandler : Handle request type before handling message.
func HTTPhandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "POST" {
		w.WriteHeader(400)
		fmt.Fprint(w, "<body><img src=\"https://api.smlk.org/mirror/ink33/what.png\" style=\"vertical-align: top\" alt=\"Bad request.\"/><ln><p>Bad request.</p></body>")
	} else {
		rid := r.Header.Get("X-Self-ID")
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		hmacsh1 := hmac.New(sha1.New, []byte(gjson.Get(*cqfunction.ConfigFile, "CoolQ.Api."+rid+".HTTPAPIPostSecret").String()))
		hmacsh1.Reset()
		hmacsh1.Write(body)
		var signature string = strings.Replace(r.Header.Get("X-Signature"), "sha1=", "", 1)
		var hmacresult string = fmt.Sprintf("%x", hmacsh1.Sum(nil))
		if signature == "" {
			w.WriteHeader(401)
			fmt.Fprint(w, "Unauthorized.")
		} else if signature != hmacresult {
			w.WriteHeader(401)
			fmt.Fprint(w, "Unauthorized.")
		} else {
			if gjson.GetBytes(body, "meta_event_type").String() != "heartbeat" {
				var msgInfoTmp = MsgHandler(body)
				msgInfoTmp.RobotID = rid
				var bc = new(botstruct.BotConfig)
				bc.HTTPAPIAddr = gjson.Get(*cqfunction.ConfigFile, "CoolQ.Api."+msgInfoTmp.RobotID+".HTTPAPIAddr").String()
				bc.HTTPAPIToken = gjson.Get(*cqfunction.ConfigFile, "CoolQ.Api."+msgInfoTmp.RobotID+".HTTPAPIToken").String()
				bc.MasterID = gjson.Get(*cqfunction.ConfigFile, "CoolQ.Master").Array()
				log.SetPrefix("SMLKBOT: ")
				go log.Println("RobotID:", rid, "Received message:", msgInfoTmp.Message, "from:", "User:", msgInfoTmp.SenderID, "Group:", msgInfoTmp.GroupID, "Role:", smlkshell.RoleHandler(msgInfoTmp, bc).RoleName)
				if msgInfoTmp.Message == "<SMLK reload" {
					if smlkshell.RoleHandler(msgInfoTmp, bc).RoleLevel == 3 {
						cqfunction.ConfigFile = cqfunction.ReadConfig()
						functionReload()
						log.Println("Succeed.")
						smlkshell.ShellLog(msgInfoTmp, bc, "succeed")
					} else {
						smlkshell.ShellLog(msgInfoTmp, bc, "deny")
					}
				} else {
					smlkshell.SmlkShell(msgInfoTmp, bc)
					for k, v := range functionList {
						go judgeandrun(k, v, msgInfoTmp, bc)
					}
				}
			}
		}
	}
}

func closeSignalHandler() {
	channel := make(chan os.Signal, 2)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-channel
		log.Println("Program stop.")
		os.Exit(0)
	}()
}

func main() {
	log.SetPrefix("SMLKBOT: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	functionLoad()

	runtime.GOMAXPROCS(runtime.NumCPU())
	closeSignalHandler()
	path := gjson.Get(*cqfunction.ConfigFile, "CoolQ.HTTPServer.ListeningPath").String()
	port := gjson.Get(*cqfunction.ConfigFile, "CoolQ.HTTPServer.ListeningPort").String()

	log.Println("Powered by Ink33")
	log.Println("Start listening", path, port)

	http.HandleFunc(path, HTTPhandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln("ListenAndServe", err)
	}
}
