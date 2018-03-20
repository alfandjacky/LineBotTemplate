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
	"strings"
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
				title := titleread(message.Text)
				//傳出句首是什麼
				var wordtitle string
				//先宣告要傳的文字再來組合
				//下面判斷條件什麼字首做什麼事
				switch title {
				case "te":
					wordtitle = tetitle()
				case "cc":
					wordtitle = coc7thtitle()
				case "AS":
					wordtitle = astitle()
				case "D66":
					wordtitle = d66title()
				case "DD":
					wordtitle = ddtitle(message.Text)
				}
				//負責傳出訊息
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(wordtitle)).Do()
				        err != nil {
					log.Print(err)
				} 
				//負責傳出訊息
			}
		}
	}
}

//判斷有沒有多次擲骰
func ddtitle(wordin string) string {
	word := "基本擲骰:\n"
	var firsttime = regexp.MustCompile(`^\S+`)
	var reg = regexp.MustCompile(`\d+(?i:d)\d+`)
	fstword := firsttime.FindString(wordin)
	times := strings.Count(fstword, "+") + 1
	totleresult := 0
	ttresolt := reg.FindAllString(fstword, times)
	for i:= 0 ; i < times ; i++ {
		word1 , number1 := ddone(ttresoltp[i])
		totleresult = totleresult + number1
		if i == times-1 {
			word = word + word1 + ""
		}else{
			word = word + word1 + "+"
		}
	}
	word = word + "\n→" + strconv.Itoa(totleresult)
	return word
	
}

//普通擲骰正則表達式提取數字
func cutmath(wordin string) (int, int) {
        var mdm = regexp.MustCompile(`\d+(?i:d)\d+`)
	var md = regexp.MustCompile(`^\d+`)
	var dm = regexp.MustCompile(`\d+$`)
	word := mdm.FindString(wordin)
	math1 ,_ :=strconv.Atoi(md.FindString(word))
	math2 ,_ :=strconv.Atoi(dm.FindString(word))
	return math1 , math2
}

//普通擲骰
func ddone(wordin string) ( string , int ) {
	word := "("
	dicenumber, diceside := cutmath(wordin)
	diceresult := make([]int, dicenumber)
	number := 0
	
	for i:=0; i < dicenumber ; i++ { 
		tmemath := diceroll(diceside) 
		diceresult[i] = tmemath
	}
	for i:=0; i < dicenumber ; i++ { 
		word1 := strconv.Itoa(diceresult[i])
		number = number + diceresult[i]
		if i == 0 {
			word = word + word1
		}else{
		word = word +"+"+ word1
		}
	}
	word = word + ")"
	return word , number
	
}

//測試
func tetitle() string {
	word := "測試輸出:"
	return word
}

//克蘇魯7th擲骰
func coc7thtitle() string {
	word := "CoC7th"
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

//句首判斷
func titleread(testword string) string {
	var word string
	aa := true
	a, _ := regexp.MatchString("(?i:^cc)", testword)
	b, _ := regexp.MatchString("(?i:^AS)", testword)
	c, _ := regexp.MatchString("(?i:^D66)", testword)
	d, _ := regexp.MatchString("(?i:^te)", testword)
	e, _ := regexp.MatchString("^[0-9]+(?i:d)[0-9]+", testword)
	switch aa{
		case a :
		word = "cc"
		case b :
		word = "AS"
		case c :
		word = "D66"
		case d :
		word = "te"
		case e :
		word = "DD"
} 
	return word
}

//產生隨機數
func diceroll(diceside int) int {
	san := rand.Intn(diceside)+1
	return san
}


//D66判定
func d66() (int,int) {
	var dice1 = diceroll(6)
	var dice2 = diceroll(6)
	return dice1 , dice2
}

	

//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+ Str1)).Do(); err != nil {log.Print(err)}
