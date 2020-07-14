package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type section struct {
	duration time.Duration
	topic    string
}

func (self section) String() string {
	return fmt.Sprintf("%6s : %30s", self.duration.String(), self.topic)
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

func parse_line(item string) section {
	pieces := strings.SplitAfterN(item, " ", 2)
	s_duration, topic := pieces[0], pieces[1]

	// Parse the duration into a time.Duration
	s_duration = strings.Trim(s_duration, " \t\n")
	s_time := strings.Split(s_duration, ":")
	minutes, err := strconv.ParseFloat(s_time[0], 64)
	check(err)
	seconds, err := strconv.ParseFloat(s_time[1], 64)
	check(err)
	duration := time.Duration(minutes*60+seconds) * time.Second
	topic = strings.Trim(topic, " \t\n")
	return section{duration, topic}
}

func obtain(fname string) []section {
	var ret []section
	file, err := os.Open(fname)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] != '#' {
			ret = append(ret, parse_line(line))
		} else {
			log.Printf("Skipping %s\n", line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ret
}

func clock(items []section) {
	font := "-*-itc bookman-*-i-*-*-*-*-*-*-*-*-*-*"
	clock_osd := xosd_create(len(items)+10, XOSD_bottom)
	clock_osd.SetFont(font)
	clock_osd.SetColour("green")

	remaining := time.Duration(0)
	for _, i := range items {
		remaining += i.duration
	}
	log.Printf("Duration is %s\n", remaining)
	end := (time.Now().Add(remaining)).Format("03:04:05 PM")
	for {
		ends_at := fmt.Sprintf("Ends at %11s", end)
		s_remaining := fmt.Sprintf("%-6s remaining", remaining.String())
		clock_osd.DisplayString(2, s_remaining)
		clock_osd.DisplayString(3, ends_at)
		delay := time.Duration(5000000000) * time.Nanosecond
		time.Sleep(delay)
		remaining -= delay
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s specfile\n", os.Args[0])
		os.Exit(-1)
	}
	items := obtain(os.Args[1])

	font := "-*-courier 10 pitch-*-r-*-*-*-*-*-*-*-*-*-*"
	// Display all items initially
	osd3 := xosd_create(len(items)+3, XOSD_bottom)
	osd3.SetFont(font)
	osd3.SetColour("gray6")
	go clock(items)
	for lno, item := range items {
		osd3.DisplayString(lno, item.String())
	}

	for lno, item := range items {
		for lno, item := range items { // Set everything to dark colours
			osd2 := xosd_create(len(items)+3, XOSD_bottom)
			osd2.SetFont(font)
			osd2.SetColour("gray6")
			osd2.DisplayString(lno, item.String())
			osd2.Destroy()
		}
		log.Printf("Highlighting %s (%d) \n", item.topic, lno)
		osd1 := xosd_create(len(items)+3, XOSD_bottom)
		osd1.SetFont(font)
		osd1.SetColour("LawnGreen")
		osd1.DisplayString(lno, item.String())
		time.Sleep(item.duration)
		osd1.Destroy()
	}
	fmt.Printf("Done!")

	time.Sleep(time.Duration(10) * time.Second)
	osd3.Destroy()
}
