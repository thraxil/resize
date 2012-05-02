package resize

import (
	"image"
	"testing"
)

func Test_MakeSizeSpec(t *testing.T) {

	ss := MakeSizeSpec("100s")
	if ss.IsFull() {
		t.Error("thinks it is full-size")
	}
	if !ss.IsSquare() { 
		t.Error("not square") 
	}
	if ss.Width() != 100 { 
		t.Error("not 100 wide") 
	}
	if ss.Height() != 100 {
		t.Error("not 100 high")
	}

	ss2 := MakeSizeSpec("100w")
	if ss2.IsFull() {
		t.Error("thinks it is full-size")
	}
	if ss2.IsSquare() { 
		t.Error("think's it is square") 
	}
	if ss2.Width() != 100 { 
		t.Error("not 100 wide") 
	}
	if ss2.Height() != -1 {
		t.Error("not -1 high")
	}

	ss3 := MakeSizeSpec("100h")
	if ss3.IsFull() {
		t.Error("thinks it is full-size")
	}
	if ss3.IsSquare() { 
		t.Error("think's it is square") 
	}
	if ss3.Width() != -1 { 
		t.Error("not -1 wide") 
	}
	if ss3.Height() != 100 {
		t.Error("not 100 high")
	}

	ss4 := MakeSizeSpec("100h200w")
	if ss4.IsFull() {
		t.Error("thinks it is full-size")
	}
	if ss4.IsSquare() { 
		t.Error("think's it is square") 
	}
	if ss4.Width() != 200 { 
		t.Error("not 200 wide") 
	}
	if ss4.Height() != 100 {
		t.Error("not 100 high")
	}

	ss5 := MakeSizeSpec("100w200h")
	if ss5.IsFull() {
		t.Error("thinks it is full-size")
	}
	if ss5.IsSquare() { 
		t.Error("think's it is square") 
	}
	if ss5.Width() != 100 { 
		t.Error("not 100 wide") 
	}
	if ss5.Height() != 200 {
		t.Error("not 200 high")
	}

	ss6 := MakeSizeSpec("full")
	if !ss6.IsFull() {
		t.Error("not full-size")
	}
	if ss6.IsSquare() { 
		t.Error("not square") 
	}
	if ss6.Width() != -1 { 
		t.Error("width set on full") 
	}
	if ss6.Height() != -1 {
		t.Error("height set on full")
	}

}

type toRectTestCase struct {
	Label string
	SizeSpec *sizeSpec
	Rect image.Rectangle
	ExpectedWidth int
	ExpectedHeight int
}

func Test_ToRect(t *testing.T) {
	full_sized := MakeSizeSpec("full")
	square_sized := MakeSizeSpec("100s")
	width_constrained := MakeSizeSpec("100w")
	height_constrained := MakeSizeSpec("100h")
	height_and_width_constrained_wh := MakeSizeSpec("100w50h")
	height_and_width_constrained_hw := MakeSizeSpec("100h50w")
	wider_than_tall := image.Rect(0,0,1000,500)
	taller_than_wide := image.Rect(0,0,500,1000)
	square_image := image.Rect(0,0,1000,1000)

	cases := []toRectTestCase{
		{
		Label: "full-sized should be no-op",
		SizeSpec: full_sized,
		Rect: square_image,
		ExpectedWidth: square_image.Dx(),
		ExpectedHeight: square_image.Dy(),
		},

		{
		Label: "square == square",
		SizeSpec: square_sized,
		Rect: square_image,
		ExpectedWidth: square_image.Dx(),
		ExpectedHeight: square_image.Dy(),
		},

		{
		Label: "square != square (wider than taller)",
		SizeSpec: square_sized,
		Rect: wider_than_tall,
		ExpectedWidth: 500,
		ExpectedHeight: 500,
		},

		{
		Label: "square != square (taller than wide)",
		SizeSpec: square_sized,
		Rect: taller_than_wide,
		ExpectedWidth: 500,
		ExpectedHeight: 500,
		},

		{
		Label: "width constrained square image",
		SizeSpec: width_constrained,
		Rect: square_image,
		ExpectedWidth: 1000,
		ExpectedHeight: 1000,
		},

		{
		Label: "width constrained wider than tall",
		SizeSpec: width_constrained,
		Rect: wider_than_tall,
		ExpectedWidth: 1000,
		ExpectedHeight: 500,
		},

		{
		Label: "width constrained taller than wide",
		SizeSpec: width_constrained,
		Rect: taller_than_wide,
		ExpectedWidth: 500,
		ExpectedHeight: 1000,
		},

		{
		Label: "height constrained square image",
		SizeSpec: height_constrained,
		Rect: square_image,
		ExpectedWidth: 1000,
		ExpectedHeight: 1000,
		},

		{
		Label: "height constrained wider than tall",
		SizeSpec: height_constrained,
		Rect: wider_than_tall,
		ExpectedWidth: 1000,
		ExpectedHeight: 500,
		},

		{
		Label: "height constrained taller than wide",
		SizeSpec: height_constrained,
		Rect: taller_than_wide,
		ExpectedWidth: 500,
		ExpectedHeight: 1000,
		},

		{
		Label: "height and width constrained (w>h) square",
		SizeSpec: height_and_width_constrained_wh,
		Rect: square_image,
		ExpectedWidth: 1000,
		ExpectedHeight: 500,
		},

		{
		Label: "height and width constrained (w>h) taller than wide",
		SizeSpec: height_and_width_constrained_wh,
		Rect: taller_than_wide,
		ExpectedWidth: 500,
		ExpectedHeight: 250,
		},

		{
		Label: "height and width constrained (w>h) wider than tall",
		SizeSpec: height_and_width_constrained_wh,
		Rect: wider_than_tall,
		ExpectedWidth: 1000,
		ExpectedHeight: 500,
		},

		{
		Label: "height and width constrained (h>w) square",
		SizeSpec: height_and_width_constrained_hw,
		Rect: square_image,
		ExpectedWidth: 500,
		ExpectedHeight: 1000,
		},

		{
		Label: "height and width constrained (h>w) taller than wide",
		SizeSpec: height_and_width_constrained_hw,
		Rect: taller_than_wide,
		ExpectedWidth: 500,
		ExpectedHeight: 1000,
		},

		{
		Label: "height and width constrained (h>w) wider than tall",
		SizeSpec: height_and_width_constrained_hw,
		Rect: wider_than_tall,
		ExpectedWidth: 500,
		ExpectedHeight: 1000,
		},

	}

	for i := range cases {
		c := cases[i]
		r := c.SizeSpec.ToRect(c.Rect)
		if r.Dx() != c.ExpectedWidth {
			t.Error(c.Label, "-- bad width",r.Dx(),"expected",c.ExpectedWidth)
		}
		if r.Dy() != c.ExpectedHeight {
			t.Error(c.Label, "-- bad height",r.Dy(),"expected",c.ExpectedHeight)
		}
	}

}
