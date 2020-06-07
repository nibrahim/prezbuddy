package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
	// "io/ioutil"
	"log"
	"os"
	"xosd"
	// "unsafe"
)

type section struct {
	duration float64
	topic    string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func display(topics []section) {
	for _, topic := range topics {
		fmt.Println(topic.topic)
		time.Sleep(time.Duration(topic.duration) * time.Second)
	}
}

func obtain(fname string) []section {
	var ret []section
	file, err := os.Open(fname)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		components := strings.SplitAfterN(scanner.Text(), " ", 2)
		duration, err := strconv.ParseFloat(strings.Trim(components[0], " \t\n"), 64)
		check(err)
		topic := strings.Trim(components[1], " \t\n")
		ret = append(ret, section{duration, topic})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ret
}

func clock(items []section) {
	font := "-*-itc bookman-*-i-*-*-*-*-*-*-*-*-*-*"
	clock_osd := xosd.New(5, xosd.XOSD_bottom)
	clock_osd.SetFont(font)
	clock_osd.SetColour("yellow")
	for {
		now := time.Now().Format("03:04:05 PM")
		clock_osd.DisplayString(2, now)
		time.Sleep(time.Duration(100000) * time.Nanosecond)
	}
}

func main() {
	items := obtain(os.Args[1])
	font := "-*-courier 10 pitch-*-r-*-*-*-*-*-*-*-*-*-*"
	// Display all items initially
	osd3 := xosd.New(10, xosd.XOSD_bottom)
	osd3.SetFont(font)
	osd3.SetColour("dark green")
	go clock(items)
	for lno, item := range items {
		t := fmt.Sprintf("%.1f  %20s", item.duration, item.topic)
		osd3.DisplayString(lno, t)
	}

	for lno, item := range items {
		for lno, item := range items { // Set everything to dark colours
			osd2 := xosd.New(10, xosd.XOSD_bottom)
			osd2.SetFont(font)
			osd2.SetColour("dark green")
			t := fmt.Sprintf("%.1f  %20s", item.duration, item.topic)
			osd2.DisplayString(lno, t)
			osd2.Destroy()
		}
		log.Printf("Highlighting %s (%d) \n", item.topic, lno)
		osd1 := xosd.New(10, xosd.XOSD_bottom)
		osd1.SetFont(font)
		osd1.SetColour("green")
		t := fmt.Sprintf("%.1f  %20s", item.duration, item.topic)
		osd1.DisplayString(lno, t)
		time.Sleep(time.Duration(item.duration) * time.Second)
		osd1.Destroy()
	}
	fmt.Printf("Done!")

	time.Sleep(time.Duration(10) * time.Second)
	osd3.Destroy()
}
