package cp

import (
	"image/color"

	"github.com/jakecoffman/cp/v2"
	"github.com/sago35/koebiten"
)

const DrawPointLineScale = 1.0

type Theme struct {
	Outline                         color.RGBA
	Shape, ShapeSleeping, ShapeIdle color.RGBA
	Constraint, CollisionPoint      color.RGBA
}

func toFColor(c color.RGBA) cp.FColor {
	r := float32(c.R) / 255.0
	g := float32(c.G) / 255.0
	b := float32(c.B) / 255.0
	a := float32(c.A) / 255.0
	return cp.FColor{R: r, G: g, B: b, A: a}
}

func DefaultTheme() *Theme {
	return &Theme{
		Outline:        color.RGBA{0xC8, 0xD2, 0xE6, 0xFF},
		ShapeSleeping:  color.RGBA{0x33, 0x33, 0x33, 0x80},
		ShapeIdle:      color.RGBA{0xA8, 0xA8, 0xA8, 0x80},
		Shape:          color.RGBA{0xB2, 0x4C, 0x99, 0x80},
		Constraint:     color.RGBA{0x00, 0xBF, 0x00, 255},
		CollisionPoint: color.RGBA{0xFF, 0x19, 0x33, 255},
	}
}

type Drawer struct {
	Screen       *koebiten.Image
	ScreenWidth  int
	ScreenHeight int
	StrokeWidth  float32
	FlipYAxis    bool
	// Drawing colors
	Theme *Theme
	// Deprecated: Use OptStroke and OptFill instead of AntiAlias
	AntiAlias bool
}

func NewDrawer(screenWidth, screenHeight int) *Drawer {
	antiAlias := true
	return &Drawer{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		AntiAlias:    antiAlias,
		StrokeWidth:  1,
		FlipYAxis:    false,
		Theme:        DefaultTheme(),
	}
}

func (d *Drawer) WithScreen(screen *koebiten.Image) *Drawer {
	d.Screen = screen
	return d
}

func (d *Drawer) DrawCircle(pos cp.Vector, angle, radius float64, outline, fill cp.FColor, data interface{}) {

}

func (d *Drawer) DrawSegment(a, b cp.Vector, fill cp.FColor, data interface{}) {

}

func (d *Drawer) DrawFatSegment(a, b cp.Vector, radius float64, outline, fill cp.FColor, data interface{}) {

}

func (d *Drawer) DrawPolygon(count int, verts []cp.Vector, radius float64, outline, fill cp.FColor, data interface{}) {
}

func (d *Drawer) DrawDot(size float64, pos cp.Vector, fill cp.FColor, data interface{}) {
}

func (d *Drawer) Flags() uint {
	return 0
}

func (d *Drawer) OutlineColor() cp.FColor {
	return toFColor(d.Theme.Outline)
}

func (d *Drawer) ShapeColor(shape *cp.Shape, data interface{}) cp.FColor {
	body := shape.Body()
	if body.IsSleeping() {
		return toFColor(d.Theme.ShapeSleeping)
	}

	if body.IdleTime() > shape.Space().SleepTimeThreshold {
		return toFColor(d.Theme.ShapeIdle)
	}
	return toFColor(d.Theme.Shape)
}

func (d *Drawer) ConstraintColor() cp.FColor {
	return toFColor(d.Theme.Constraint)
}

func (d *Drawer) CollisionPointColor() cp.FColor {
	return toFColor(d.Theme.CollisionPoint)
}

func (d *Drawer) Data() interface{} {
	return nil
}
