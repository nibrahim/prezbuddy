#include <xosd.h>

#define NKEYS 10

xosd *
configure_osd(int lines)
{
  xosd *osd;
  osd = xosd_create (lines);

  /* xosd_set_font(osd, "-adobe-*-*-r-*-*-*-120-*-*-*-*-*-*"); */
  /* xosd_set_pos(osd, XOSD_top); */
  /* xosd_set_align(osd, XOSD_right); */

  /* xosd_set_colour(osd, "green"); */
  /* xosd_set_outline_colour(osd, "black"); */
  /* xosd_set_outline_offset(osd, 2); */
  /* xosd_set_shadow_colour(osd, "grey"); */
  /* xosd_set_shadow_offset(osd, 3); */

  /* xosd_set_timeout(osd, -1); */
  return osd;
}

void
display_string(xosd * osd, int l, char *line) {
  xosd_display(osd, l, XOSD_printf, line);
}
