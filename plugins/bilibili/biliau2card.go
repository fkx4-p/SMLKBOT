package bilibili

import (
	"SMLKBOT/data/botstruct"
	"SMLKBOT/utils/cqfunction"
	"log"
	"regexp"
	"strings"
)

//GetAu : Get audio number by regexp.
func GetAu(msg string) (au string) {
	if strings.Contains(msg, "CQ:") {
		return ""
	}

	reg := regexp.MustCompile("(?i)au[0-9]+")

	str := strings.Join(reg.FindAllString(msg, 1), "")
	return str
}

// Au2Card : Handle meassage and send music card.
func Au2Card(MsgInfo *botstruct.MsgInfo, BotConfig *botstruct.BotConfig) {
	au := GetAu(MsgInfo.Message)

	if au != "" {
		log.SetPrefix("BiliAu2Card: ")
		log.Println("Known command:", au)
		AuInfo := GetAuInfo(au)

		if !AuInfo.AuStatus {
			msgMake := "BiliAu2Card: AU" + AuInfo.AuNumber + AuInfo.AuMsg
			switch MsgInfo.MsgType {
			case "private":
				go cqfunction.CQSendPrivateMsg(MsgInfo.SenderID, msgMake, BotConfig)
				break
			case "group":
				go cqfunction.CQSendGroupMsg(MsgInfo.GroupID, msgMake, BotConfig)
				break
			}
		} else {
			cqCodeMake := "[CQ:music,type=custom,url=" + AuInfo.AuJumpURL + ",audio=" + AuInfo.AuURL + ",title=" + AuInfo.AuTitle + ",content=" + AuInfo.AuDesp + ",image=" + AuInfo.AuCoverURL + "@180w_180h]"
			switch MsgInfo.MsgType {
			case "private":
				go cqfunction.CQSendPrivateMsg(MsgInfo.SenderID, cqCodeMake, BotConfig)
				break
			case "group":
				go cqfunction.CQSendGroupMsg(MsgInfo.GroupID, cqCodeMake, BotConfig)
				break
			}
		}
	}
}
