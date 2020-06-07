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

// xosd *
// configure_osd(int lines)
// {
//   xosd *osd;
//   osd = xosd_create (NKEYS);

//   xosd_set_font(osd, "-adobe-courier-bold-r-normal--60-320-*-*-*-*-*-*");
//   xosd_set_pos(osd, XOSD_top);
//   xosd_set_align(osd, XOSD_right);

//   xosd_set_colour(osd, "green");
//   xosd_set_outline_colour(osd, "black");
//   xosd_set_outline_offset(osd, 2);
//   xosd_set_shadow_colour(osd, "grey");
//   xosd_set_shadow_offset(osd, 3);

//   xosd_set_timeout(osd, -1);
//   return osd;
// }

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

func main() {
	items := obtain(os.Args[1])
	font := "-*-courier 10 pitch-*-r-*-*-*-*-*-*-*-*-*-*"
	// Display all items initially
	osd3 := xosd.New(10)
	osd3.SetFont(font)
	osd3.SetColour("dark green")
	for lno, item := range items {
		t := fmt.Sprintf("%.1f  %20s", item.duration, item.topic)
		osd3.DisplayString(lno, t)
	}

	for lno, item := range items {

		for lno, item := range items { // Set everything to dark colours
			osd2 := xosd.New(10)
			osd2.SetFont(font)
			osd2.SetColour("dark green")
			t := fmt.Sprintf("%.1f  %20s", item.duration, item.topic)
			osd2.DisplayString(lno, t)
			osd2.Destroy()
		}
		log.Printf("Highlighting %s (%d) \n", item.topic, lno)
		osd1 := xosd.New(10)
		osd1.SetFont(font)
		osd1.SetColour("green")
		t := fmt.Sprintf("%.1f  %20s", item.duration, item.topic)
		osd1.DisplayString(lno, t)
		time.Sleep(time.Duration(item.duration) * time.Second)
		osd1.Destroy()
	}
	fmt.Printf("Done!")

	time.Sleep(time.Duration(1) * time.Second)
	osd3.Destroy()
}
