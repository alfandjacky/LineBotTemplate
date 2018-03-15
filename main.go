// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"regexp"
	"net/http"
	"os"
	"math/rand"
	"strconv"
	"time"
	
	"github.com/line/line-bot-sdk-go/linebot"
)
var Str1 string

var bot *linebot.Client

func main() {
	var err error
	rand.Seed(time.Now().UnixNano())
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}


func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	 
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//以上已經篩選好訊息 純文字
				title := cut(message.Text)
				var wordtitle string
				//這邊要想辦法切文字然後回傳結果
				//切好的文字傳到要帶入的數值
				//下面判斷條件
				switch title {
				case "te":
					wordtitle = tetitle()
				case "cc":
					wordtitle = coc7thtitle()
				case "AS":
					wordtitle = astitle()
				case "D66":
					wordtitle = d66title()
				}
				//負責穿傳出訊息
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(wordtitle)).Do()
				        err != nil {
					log.Print(err)
				} 
				//負責傳出訊息
			}
		}
	}
}
//測試
func tetitle() string {
	word := "測試輸出:"
	return word
}
//克蘇魯7th擲骰
func coc7thtitle() string {
	word := "CoC7th擲骰:"
	return word
}
//絕對奴隸擲骰
func astitle() string {
	word := "絕對隸奴擲骰:"
	return word
}
//D66擲骰
func d66title() string {
	word := "D66擲骰:"
	return word
}

//文字切片+判斷
func cut(testword string) string {
	var word string
	a, _ := regexp.MatchString("(?i:^cc)", testword)
	b, _ := regexp.MatchString("^AS", testword)
	c, _ := regexp.MatchString("^D66", testword)
	d, _ := regexp.MatchString("^te", testword)
	if a {
		word = "cc"
	} else if b {
		word = "AS"
	} else if c {
		word = "D66"
	} else if d {
		word = "te"
	} 
	return word
}

//產生隨機數
func diceroll(diceside int) int {
	san := rand.Intn(diceside)+1
	return san
}
//D66判定
func d66() string {
	var dice1 = diceroll(6)
	var dice2 = diceroll(6)
	diceresult := dice1*10 + dice2
	Str1 := strconv.Itoa(diceresult)
	return Str1
}

	

//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+ Str1)).Do(); err != nil {log.Print(err)}
