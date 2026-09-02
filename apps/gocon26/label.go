package main

import (
	"image/color"
)

type label struct {
	buf   []uint16
	w, h  int16
	scale int16
}

func newLabel(w, h int16) *label {
	return &label{
		buf:   make([]uint16, int(w)*int(h)),
		w:     w,
		h:     h,
		scale: 1,
	}
}

func (l *label) Display() error { return nil }
func (l *label) Size() (int16, int16) {
	return l.w / l.scale, l.h / l.scale
}
func (l *label) SetPixel(x, y int16, c color.RGBA) {
	v := rgbaTo565(c)
	for dy := int16(0); dy < l.scale; dy++ {
		for dx := int16(0); dx < l.scale; dx++ {
			ry := y*l.scale + dy
			rx := x*l.scale + dx
			if ry >= 0 && ry < l.h && rx >= 0 && rx < l.w {
				l.buf[int(rx)+int(ry)*int(l.w)] = v
			}
		}
	}
}

func (l *label) DrawBitmap(src []uint16, srcW, srcH, x, y int16) {
	for sy := int16(0); sy < srcH; sy++ {
		for sx := int16(0); sx < srcW; sx++ {
			px, py := x+sx, y+sy
			if px >= 0 && px < l.w && py >= 0 && py < l.h {
				l.buf[py*l.w+px] = src[sy*srcW+sx]
			}
		}
	}
}

func (l *label) FillScreen(c color.RGBA) {
	for i := range l.buf {
		l.buf[i] = rgbaTo565(c)
	}
}

func (l *label) Scale(s int16) *label {
	return &label{
		buf:   make([]uint16, int(l.w)*int(l.h)),
		w:     l.w,
		h:     l.h,
		scale: s,
	}
}

func rgbaTo565(c color.RGBA) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF800) + ((g & 0xFC00) >> 5) + ((b & 0xF800) >> 11))
}
