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

func Test_ToRect(t *testing.T) {
	full_sized := MakeSizeSpec("full")
	square_sized := MakeSizeSpec("100s")
	width_constrained := MakeSizeSpec("100w")
	height_constrained := MakeSizeSpec("100h")
	height_and_width_constrained_wh := MakeSizeSpec("200w100h")
	height_and_width_constrained_hw := MakeSizeSpec("200h100w")
	wider_than_tall := image.Rect(0,0,1000,500)
	taller_than_wide := image.Rect(0,0,500,1000)
	square_image := image.Rect(0,0,1000,1000)


	// full-size should be no-op
	results := full_sized.ToRect(square_image)
	if results.Dx() != square_image.Dx() {
		t.Error("bad width")
	}
	if results.Dy() != square_image.Dy() {
		t.Error("bad height")
	}

	// square == square
	results2 := square_sized.ToRect(square_image)
	if results2.Dx() != square_image.Dx() {
		t.Error("bad width")
	}
	if results2.Dy() != square_image.Dy() {
		t.Error("bad height")
	}

	// square != square (wider than taller)
	results3 := square_sized.ToRect(wider_than_tall)
	if results3.Dx() != 500 {
		t.Error("bad width")
	}
	if results3.Dy() != 500 {
		t.Error("bad height")
	}

	// square != square (taller than wider)
	results4 := square_sized.ToRect(taller_than_wide)
	if results4.Dx() != 500 {
		t.Error("bad width")
	}
	if results4.Dy() != 500 {
		t.Error("bad height")
	}

	// width constrained square image
	results5 := width_constrained.ToRect(square_image)
	if results5.Dx() != 1000 {
		t.Error("bad width")
	}
	if results5.Dy() != 1000 {
		t.Error("bad height")
	}

	// width constrained wider than tall
	results6 := width_constrained.ToRect(wider_than_tall)
	if results6.Dx() != 1000 {
		t.Error("bad width")
	}
	if results6.Dy() != 500 {
		t.Error("bad height")
	}

	// width constrained taller than wide
	results7 := width_constrained.ToRect(taller_than_wide)
	if results7.Dx() != 500 {
		t.Error("bad width")
	}
	if results7.Dy() != 1000 {
		t.Error("bad height")
	}

	// height constrained square image
	results8 := height_constrained.ToRect(square_image)
	if results8.Dx() != 1000 {
		t.Error("bad width")
	}
	if results8.Dy() != 1000 {
		t.Error("bad height")
	}

	// height constrained wider than tall
	results9 := height_constrained.ToRect(wider_than_tall)
	if results9.Dx() != 1000 {
		t.Error("bad width")
	}
	if results9.Dy() != 500 {
		t.Error("bad height")
	}

	// height constrained taller than wide
	results10 := height_constrained.ToRect(taller_than_wide)
	if results10.Dx() != 500 {
		t.Error("bad width")
	}
	if results10.Dy() != 1000 {
		t.Error("bad height")
	}

	// height and width constrained (w>h) square
	results11 := height_and_width_constrained_wh.ToRect(square_image)
	if results11.Dx() != 1000 {
		t.Error("bad width")
	}
	if results11.Dy() != 500 {
		t.Error("bad height")
	}

	// height and width constrained (w>h) taller than wide
	results12 := height_and_width_constrained_wh.ToRect(taller_than_wide)
	if results12.Dx() != 500 {
		t.Error("bad width")
	}
	if results12.Dy() != 250 {
		t.Error("bad height")
	}

	// height and width constrained (w>h) wider than tall
	results13 := height_and_width_constrained_wh.ToRect(wider_than_tall)
	if results13.Dx() != 500 {
		t.Error("bad width")
	}
	if results13.Dy() != 1000 {
		t.Error("bad height")
	}

	// height and width constrained (h>w) square
	results14 := height_and_width_constrained_hw.ToRect(square_image)
	if results14.Dx() != 500 {
		t.Error("bad width")
	}
	if results14.Dy() != 1000 {
		t.Error("bad height")
	}

	// height and width constrained (h>w) taller than wide
	results15 := height_and_width_constrained_hw.ToRect(taller_than_wide)
	if results15.Dx() != 500 {
		t.Error("bad width")
	}
	if results15.Dy() != 1000 {
		t.Error("bad height")
	}

	// height and width constrained (h>w) wider than tall
	results16 := height_and_width_constrained_hw.ToRect(wider_than_tall)
	if results16.Dx() != 500 {
		t.Error("bad width")
	}
	if results16.Dy() != 1000 {
		t.Error("bad height")
	}

}
