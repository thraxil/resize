// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resize

import (
	"code.google.com/p/graphics-go/graphics"
	"fmt"
	"image"
	"regexp"
	"strconv"
)

type sizeSpec struct {
	width  int
	height int
	square bool
	full   bool
}

// sizes are specified with a short string that can look like
//   full - full size, will not scale the image
//   100s - make a 100 pixel square image
//   200w - will make the image 200 pixels wide, preserving original aspect ratio
//   100h - will make it 100 pixels high, preserving original aspect ratio
//   100h300w - will make it 100 pixels high 
//              and 300 wide (width and height can be specified in either order)
//
// images will always be cropped to match the desired aspect ratio rather than 
// squished, cropping will always be centered.
//
// if 'full' or 's' are specified, they will take precedent over
// width and height specs.
// 
// see Test_MakeSizeSpec in resize_test.go for more examples

func MakeSizeSpec(str string) *sizeSpec {
	s := sizeSpec{}
	if str == "full" {
		s.full = true
		s.width = -1
		s.height = -1
		return &s
	}
	r, _ := regexp.Compile("\\d+s")
	if m := r.FindString(str); m != "" {
		w, _ := strconv.Atoi(m[:len(m)-1])
		s.width = w
		s.height = w
		s.square = true
		return &s
	}
	// not full size or square, so we need to parse individual dimensions
	s.square = false
	s.full = false

	r, _ = regexp.Compile("\\d+w")
	if m := r.FindString(str); m != "" {
		w, _ := strconv.Atoi(m[:len(m)-1])
		s.width = w
	} else {
		// width was not set
		s.width = -1
	}
	r, _ = regexp.Compile("\\d+h")
	if m := r.FindString(str); m != "" {
		h, _ := strconv.Atoi(m[:len(m)-1])
		s.height = h
	} else {
		// height was not set
		s.height = -1
	}

	return &s
}

func (self *sizeSpec) IsSquare() bool {
	return self.square
}

func (self *sizeSpec) IsFull() bool {
	return self.full
}

func (self *sizeSpec) Width() int {
	return self.width
}

func (self *sizeSpec) Height() int {
	return self.height
}

// given an image size (as image.Rect), we match it up
// to the sizeSpec and return a new image.Rect which is 
// essentially, the dimensions to crop the image to before scaling

func (self *sizeSpec) ToRect(rect image.Rectangle) image.Rectangle {
	if self.full || self.Width() == -1 || self.Height() == -1 {
		// full-size or only scaling one dimension, means we deal with the whole thing
		return rect
	}

	if self.square {
		if rect.Dx() == rect.Dy() {
			// already square. WIN.
			return rect
		}
		if rect.Dx() > rect.Dy() {
			// wider than taller, crop and center on width
			trim := (rect.Dx() - rect.Dy()) / 2
			return image.Rect(trim, 0, rect.Dx()-trim, rect.Dy())
		} else {
			// taller than wider, crop and center on height
			trim := (rect.Dy() - rect.Dx()) / 2
			return image.Rect(0, trim, rect.Dx(), rect.Dy()-trim)
		}
	}
	// scaling both width and height
	if self.Width() > self.Height() {
		if rect.Dx() == rect.Dy() {
			// keep width, trim height
			ratio := float64(self.Height()) / float64(self.Width())
			targetHeight := int(ratio * float64(rect.Dx()))
			trim := targetHeight / 2
			return image.Rect(0, trim, rect.Dx(), rect.Dy()-trim)
		} else {
			if rect.Dx() > rect.Dy() {
				ratio := float64(rect.Dy()) / float64(rect.Dx())
				outRatio := float64(self.Height()) / float64(self.Width())
				if ratio == outRatio {
					return rect
				}
				if outRatio > ratio {
					// 
				} else {

				}
				rHeight := int(float64(rect.Dy()) * ratio)
				trim := (rect.Dy() - rHeight) / 2
				return image.Rect(0, trim, rect.Dx(), rect.Dy()-trim)
			} else {
				// rect.Dy() is the keeper
				ratio := float64(rect.Dx()) / float64(self.Width())
				rHeight := ratio * float64(self.Height())
				trim := int((float64(rect.Dy()) - rHeight) / 2)
				return image.Rect(0, trim, rect.Dx(), rect.Dy()-trim)
			}
		}
	} else {
		if rect.Dx() == rect.Dy() {
			// keep height, trim width
			ratio := float64(self.Width()) / float64(self.Height())
			targetWidth := int(ratio * float64(rect.Dy()))
			trim := targetWidth / 2
			return image.Rect(trim, 0, rect.Dx()-trim, rect.Dx())
		} else {
			if rect.Dx() > rect.Dy() {
				// rect.Dx() is the keeper
				ratio := float64(self.Width()) / float64(self.Height())
				targetWidth := int(ratio * float64(rect.Dy()))
				trim := targetWidth / 2
				return image.Rect(trim, 0, rect.Dx()-trim, rect.Dx())
			} else {

			}
		}

	}
	return rect
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// size of the image that will result from resizing one of the
// specified rect to this sizeSpec
func (self *sizeSpec) TargetWH(rect image.Rectangle) (int, int) {
	if self.full {
		return rect.Dx(), rect.Dy()
	}
	if self.square {
		return self.width, self.height
	}
	if self.width == -1 {
		ratio := float64(rect.Dy()) / float64(self.height)
		x := int(float64(rect.Dx()) / ratio)
		return x, self.height
	}
	if self.height == -1 {
		ratio := float64(rect.Dx()) / float64(self.width)
		x := int(float64(rect.Dy()) / ratio)
		return self.width, x
	}

	return self.width, self.height
}

// Resize returns a scaled copy of the image slice r of m.
// The returned image has width w and height h.
func Resize(m image.Image, sizeStr string) image.Image {
	var w, h int

	ss := MakeSizeSpec(sizeStr)
	r := ss.ToRect(m.Bounds())
	w, h = ss.TargetWH(m.Bounds())

	if w < 0 || h < 0 {
		return nil
	}
	if w == 0 || h == 0 || r.Dx() <= 0 || r.Dy() <= 0 {
		return image.NewRGBA64(r)
	}

	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	if err := graphics.Thumbnail(dst, m); err != nil {
		fmt.Println("could not thumbnail")
	}
	return dst
}
