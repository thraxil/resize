package resize

import (
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
