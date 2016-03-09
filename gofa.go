package main

import (
    "fmt"
	"log"
	"math/rand"
	"time"
	"net/url"
	"flag"
	"strconv"
    "github.com/parnurzeal/gorequest"
)

func main() {
	flag.Parse()
    s1 := flag.Arg(0)
	s2 := flag.Arg(1)
	id, err1:= strconv.Atoi(s1)
	num, err2:= strconv.Atoi(s2)
    if err1 != nil { id = 5 }
	if err2 != nil { num = 10 }
	fmt.Println(id, num)
	done := make(chan bool, num)
	for i := 0; i < num; i++ {
		go geti(id, done)
	}
	for i := 0; i < num; i++ {
	    <-done
        fmt.Println("done")
	}
	fmt.Println("main done")
	
}

func geti(id int, done chan bool) {
	fmt.Println("start geti")
	defer fmt.Println("end geti")
	template := "http://events.chncpa.org/wmx2016/action/pctou.php?id=%d&user_ip=%s&time=%s"
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	ip := fmt.Sprintf("%d.%d.%d.%d", r1.Intn(250)+1, r1.Intn(250)+1, r1.Intn(250)+1, r1.Intn(250)+1)
	t := time.Now()
	ts := t.Add(-24*time.Hour)
	//fmt.Println(url.QueryEscape(ts.Format("2006-01-02 15:04:05")))
	dtime := url.QueryEscape(ts.Format("2006-01-02 15:04:05")) //"2016-02-26+28%3A35%3A38"
	url := fmt.Sprintf(template, id, ip, dtime)
	//url = "http://wtfismyip.com/text"
	fmt.Println(url)
    request := gorequest.New()
	resp, body, errs := request.Get(url).
		Set("X-Requested-With", "XMLHttpRequest").
		Set("Referer", "http://events.chncpa.org/wmx2016/mobile/pages/jmpx.php").
		Set("User-Agent", "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405").
		End()
	if errs != nil {
		log.Fatal(errs)
        log.Fatal("error")
		log.Fatal(body)
		log.Fatal(resp)
		return
	} else {
		fmt.Println(body)
	}
	
	done <- true
}
