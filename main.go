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
				var firsttime = regexp.MustCompile(`^\S+`)
				fstword := firsttime.FindString(message.Text)
				title := titleread(fstword)
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
					wordtitle = astitle(fstword)
				case "D66":
					_, _, word66 := d66title()
					wordtitle = "D66擲骰:\n" + word66
				case "DD":
					wordtitle = ddtitle(fstword)
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

//輸入比較符號與兩數字給出真偽
func camepareto (wordin string,num1 int,num2 int)(string){
	word := ""
	switch wordin {
			case ">=":
		if num1 >= num2 {
			word = "true"
		}else{
			word = "false"
		}
			case "<=":
		if num1 <= num2 {
			word = "true"
		}else{
			word = "false"
		}
			case "=":
		if num1 == num2 {
			word = "true"
		}else{
			word = "false"
		}
			case ">":
		if num1 > num2 {
			word = "true"
		}else{
			word = "false"
		}
			case "<":
		if num1 < num2 {
			word = "true"
		}else{
			word = "false"
		}
	}
	return word
}


//執行多次普通擲骰
func ddtitle(wordin string) string {
	var reg = regexp.MustCompile(`\d+(?i:d)\d+`)
	compare, _ := regexp.MatchString("[0-9]+[>=<]{1,2}[0-9]+$", wordin)
	word := "基本擲骰:\n"+"("+wordin+")\n→"
	times := strings.Count(wordin, "+") + 1
	totleresult := 0
	ttresolt := reg.FindAllString(wordin, times)
	for i:= 0 ; i < times ; i++ {
		word1 , number1 := ddone(ttresolt[i])
		totleresult = totleresult + number1
		if i == times-1 {
			word = word + word1 + ""
		}else{
			word = word + word1 + "+"
		}
	}
	
	if compare {
		var comeparetype = regexp.MustCompile(`[>=<]{1,2}`)
		var numbercompare = regexp.MustCompile("[0-9]+$")
		ase := numbercompare.FindString(wordin)
		moon1 := comeparetype.FindString(wordin)
		int11,_ :=strconv.Atoi(ase)  
		word = word+moon1+ase+"\n→"+camepareto (moon1,totleresult,int11)	
	}else{
		word = word + "\n→" + strconv.Itoa(totleresult)
	}
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
	var a,b,c,d,e,f int =0,0,0,0,0,0
	for i := 0 ; i <20 ; i++ {
		abc := diceroll(6)
		switch abc {
			case 1 :
			a++
			case 2 :
			b++
			case 3 :
			c++
			case 4 :
			d++
			case 5 :
			e++
			case 6 :
			f++
		}
		aa:=strconv.Itoa(a)  
		bb:=strconv.Itoa(b)  
		cc:=strconv.Itoa(c)  
		dd:=strconv.Itoa(d)  
		ff:=strconv.Itoa(e)  
		ee:=strconv.Itoa(f)  
		word = word + aa+bb+cc+dd+ee+ff
	}
	return word
}

//克蘇魯7th擲骰
func coc7thtitle() string {
	word := "CoC7th"
	return word
}

//絕對奴隸擲骰
func astitle( wordin string ) string {
	word := "絕對隸奴擲骰:\n→"
	var cutas = regexp.MustCompile("[^(?i:^as)]+")
	ase := cutas.FindString(wordin)
	var num1 = regexp.MustCompile("^[0-9]+")
	number1 := num1.FindString(ase)
	word1 ,_ :=strconv.Atoi(number1)
	aresult,agreat,aword := asd66 (word1)
	awresult :=strconv.Itoa(aresult)
	word = word + aword
	if aword == "(大失敗)" {
	}else{
	a, _ := regexp.MatchString("[^0-9]+", ase)
	if a {
		var num2 = regexp.MustCompile("[0-9]+$")
		number2 := num2.FindString(ase)
		word2 ,_ :=strconv.Atoi(number2)
		var cutnum = regexp.MustCompile("[^0-9]+")
		compare := cutnum.FindString(ase)
		if compare == "v"{
			bresult,_,bword := asd66 (word2)
			bwresult :=strconv.Itoa(bresult)
			word = word + "VS" + bword + "\n→" + "(" + awresult+ ")" + "VS" + "(" + bwresult + ")" +"\n→"
			if aresult >= bresult {
				word = word + "true"
			}else{
				word = word + "false"
			}
			word = word +"\n→DK增加" + agreat
			
		}else if compare == ">="{
			if aresult >= word2 {
				word = word + ">=" + number2 + "\n→" + awresult+ ">=" + number2+ "\n→" + "true" + "\n→DK增加" + agreat
			}else{
				word = word + ">=" + number2 + "\n→" + awresult+ ">=" + number2+ "\n→" + "false" + "\n→DK增加" + agreat
			}
			
		}
	}else{
		word = word + "\n→" + awresult + "\n→DK增加" + agreat
	}
	}
	return word
}
//絕對隸奴D66
func asd66 (numin int) (int , string , string){
	a , b , word66 := d66title()
	sixtime := strings.Count(word66, "6")
	fivetime := strings.Count(word66, "5")
	resort := a + b - (6*sixtime)
	sixtime2 := sixtime*sixtime
	sixword := strconv.Itoa(sixtime2)
	numout := numin - resort
	if fivetime == 2 {
		word66 = "(大失敗)"
	}
	return numout ,sixword ,word66
}
//D66擲骰
func d66title() (int,int,string) {
	var dice1 = diceroll(6)
	var dice2 = diceroll(6)
	word1 := strconv.Itoa(dice1)
	word2 := strconv.Itoa(dice2)
	word := "(" + word1 +"," + word2 +")" 
	return dice1,dice2,word
}

//句首判斷
func titleread(testword string) string {
	var word string
	aa := true
	a, _ := regexp.MatchString("(?i:^cc)[0-9]", testword)
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




	

//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+ Str1)).Do(); err != nil {log.Print(err)}
