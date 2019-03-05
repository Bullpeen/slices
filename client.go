package slices

import (
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
)

const apiEndpoint = "https://scorescraper.herokuapp.com/"
const apiUserPath = "user/"

/*
"play_count": 151,
"rank_global": 18251,
"rank_region": 9070,
"scores": [],
"total_pp": 1386.03,
"total_score": 49711395,
"user_id": 76561198057633470,
"user_name": "tyler.morita"
 */

type User struct {
	PlayCount  int         `json:"play_count"`
	RankGlobal int         `json:"rank_global"`
	RankRegion int         `json:"rank_region"`
	Scores     []UserScore `json:"scores"`
	TotalPP    float64     `json:"total_pp"`
	TotalScore int         `json:"total_score"`
	UserId     int64       `json:"user_id"`
	UserName   string      `json:"user_name"`
}

/*
{
"accuracy": 0.7711,
"author": "Speoghi",
"difficulty": "Expert",
"max_pp": 127.45110074474177,
"net_pp": 88.99,
"raw_pp": 88.99,
"song_id": 11805,
"song_rank": 2966,
"time": "2019-03-02 04:54:47 UTC",
"title": "Katy Perry - California Gurls"
},
 */

type UserScore struct {
	Accuracy   float64 `json:"accuracy"`
	Author     string  `json:"author"`
	Difficulty string  `json:"difficulty"`
	MaxPP      float64 `json:"max_pp"`
	NetPP      float64 `json:"net_pp"`
	RawPP      float64 `json:"raw_pp"`
	SongId     int64   `json:"song_id"`
	SongRank   int64   `json:"song_rank"`
	Time       string  `json:"time"`
	Title      string  `json:"title"`
}

func GetUser(user string) (*User, error) {
	endpoint := apiEndpoint + apiUserPath + user

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client { Timeout: time.Duration(10 * time.Second) }

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ret := &User{}
	err = json.Unmarshal(body, ret)

	return ret, err
}

type ByTotalPP []*User

func (s ByTotalPP) Len() int { return len(s) }
func (s ByTotalPP) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByTotalPP) Less(i, j int) bool {
	if s[i] == nil { return false }
	if s[j] == nil { return true }
	return s[i].TotalPP > s[j].TotalPP
}