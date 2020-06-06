package xosd

import "fmt"

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lxosd
// #include <xosd.h>
// xosd *configure_osd(int lines);
// void display_string(xosd *osd , int l, char *line);
import "C"

type osd struct {
	c_xosd *C.struct_xosd
}

func New(lines int) osd {
	var ret osd
	ret.c_xosd = C.configure_osd(C.int(lines))
	C.xosd_set_timeout(ret.c_xosd, -1)
	C.xosd_set_pos(ret.c_xosd, C.XOSD_top)
	C.xosd_set_align(ret.c_xosd, C.XOSD_right)
	return ret
}

func (self osd) Print() {
	fmt.Printf("%v\n", self)
}

func (self osd) SetFont(fontname string) {
	c_fontname := C.CString(fontname)
	C.xosd_set_font(self.c_xosd, c_fontname)
	// xosd_set_font(osd, "-adobe-*-*-r-*-*-*-120-*-*-*-*-*-*")
	// self.xosd = C.configure_osd()
}

func (self osd) SetColour(colourname string) {
	c_colourname := C.CString(colourname)
	C.xosd_set_colour(self.c_xosd, c_colourname)
}

func (self osd) DisplayString(pos int, text string) {
	c_text := C.CString(text)
	C.display_string(self.c_xosd, C.int(pos), c_text)
}

func (self osd) Destroy() {
	C.xosd_destroy(self.c_xosd)
}

func OsdPrint(text string) {
	// xosd := osd.Init()

	// xosd.SetColour("violet red")
	// xosd.DisplayString("Hi there")

	// c_string := C.CString(text)

	// C.display_string(xosd, 1, c_string)
	// C.xosd_set_colour(xosd, C.CString("violet red"))
	// C.display_string(xosd, 1, c_string)
	// C.display_string(xosd, 2, c_string)

	// for i := 1; i < 5; i++ {
	// 	C.display_string(xosd, C.int(i), c_string)
	// }
	// c.xosd_destroy(xosd)
}
