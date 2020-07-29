package vtbmusic

import "github.com/tidwall/gjson"

//Translate json to go by using https://www.sojson.com/json/json2go.html

//getMusicList also can be used as GetHotMusicList
type getMusicList struct {
	Total int `json:"total"`
	Data  []struct {
		ID              string `json:"id"`
		CreateTime      string `json:"createTime"`
		PublishTime     string `json:"publishTime"`
		CreatorID       string `json:"creatorId"`
		CreatorRealName string `json:"creatorRealName"`
		Deleted         bool   `json:"deleted"`
		OriginName      string `json:"originName"`
		VocalID         string `json:"vocalId"`
		VocalName       string `json:"vocalName"`
		CoverImg        string `json:"coverImg"`
		Music           string `json:"music"`
		Lyric           string `json:"lyric"`
		Cdn             string `json:"cdn"`
		Source          string `json:"source"`
		BiliBili        string `json:"biliBili"`
		YouTube         string `json:"youTube"`
		Twitter         string `json:"twitter"`
		Likes           int    `json:"likes"`
		Length          int    `json:"length"`
		Label           string `json:"label"`
		VocalList       []struct {
			ID         string `json:"id"`
			Cn         string `json:"cn"`
			Jp         string `json:"jp"`
			En         string `json:"en"`
			Originlang string `json:"originlang"`
		} `json:"vocalList"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
	Msg       string `json:"msg"`
}

type getCDNList struct {
	Total int `json:"total"`
	Data  []struct {
		ID         string `json:"id"`
		CreateTime string `json:"createTime"`
		CreatorID  string `json:"creatorId"`
		Name       string `json:"name"`
		URL        string `json:"url"`
		Info       string `json:"info"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
	Msg       string `json:"msg"`
}
type getMusicData struct {
	ID              string `json:"id"`
	CreateTime      string `json:"createTime"`
	PublishTime     string `json:"publishTime"`
	CreatorID       string `json:"creatorId"`
	CreatorRealName string `json:"creatorRealName"`
	Deleted         bool   `json:"deleted"`
	OriginName      string `json:"originName"`
	VocalID         string `json:"vocalId"`
	VocalName       string `json:"vocalName"`
	CoverImg        string `json:"coverImg"`
	Musics          string `json:"musics"`
	Lyric           string `json:"lyric"`
	Cdn             string `json:"cdn"`
	Source          string `json:"source"`
	BiliBili        string `json:"biliBili"`
	YouTube         string `json:"youTube"`
	Twitter         string `json:"twitter"`
	Likes           int    `json:"likes"`
	Length          int    `json:"length"`
	Label           string `json:"label"`
}
type getVtbsList struct {
	Total int `json:"total"`
	Data  []struct {
		ID           string `json:"id"`
		CreateTime   string `json:"createTime"`
		CreatorID    string `json:"creatorId"`
		Deleted      bool   `json:"deleted"`
		OriginalName string `json:"originalName"`
		ChineseName  string `json:"chineseName"`
		JapaneseName string `json:"japaneseName"`
		EnglistName  string `json:"englistName"`
		GroupsID     string `json:"groupsId"`
		AvatarImg    string `json:"avatarImg"`
		Bilibili     string `json:"bilibili"`
		YouTube      string `json:"youTube"`
		Twitter      string `json:"twitter"`
		Watch        int    `json:"watch"`
		Introduce    string `json:"introduce"`
	} `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
	Msg       string `json:"msg"`
}
type getVtbsData struct {
	ID           string `json:"id"`
	CreateTime   string `json:"createTime"`
	CreatorID    string `json:"creatorId"`
	Deleted      bool   `json:"deleted"`
	OriginalName string `json:"originalName"`
	ChineseName  string `json:"chineseName"`
	JapaneseName string `json:"japaneseName"`
	EnglistName  string `json:"englistName"`
	GroupsID     string `json:"groupsId"`
	AvatarImg    string `json:"avatarImg"`
	Bilibili     string `json:"bilibili"`
	YouTube      string `json:"youTube"`
	Twitter      string `json:"twitter"`
	Watch        int    `json:"watch"`
	Introduce    string `json:"introduce"`
}

//MusicInfo includes the info of a music.
type MusicInfo struct {
	MusicName  string
	MusicID    string
	MusicVocal string
	Cover      string
	MusicURL   string
	MusicCDN   string
}

//MusicList includes the result of searching for musics.
type MusicList struct {
	Total int
	Data  []gjson.Result
}

//VtbsList includes the result of searching for Vtbs.
type VtbsList struct {
	Total int
	Data  []gjson.Result
}