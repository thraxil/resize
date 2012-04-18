// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resize

import (
	"image"
	"image/color"
)

// Resize returns a scaled copy of the image slice r of m.
// The returned image has width w and height h.
func Resize(m image.Image, r image.Rectangle, w, h int) image.Image {
	if w < 0 || h < 0 {
		return nil
	}
	if w == 0 || h == 0 || r.Dx() <= 0 || r.Dy() <= 0 {
		return image.NewRGBA64(r)
	}
	switch m := m.(type) {
	case *image.RGBA:
		return resizeRGBA(m, r, w, h)
	}
	return image.NewRGBA64(r)
}

// average convert the sums to averages and returns the result.
func average(sum []uint64, w, h int, n uint64) image.Image {
	r := image.Rect(0, 0, w, h)
	ret := image.NewRGBA(r)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			i := y*ret.Stride + x*4
			j := 4 * (y*w + x)
			ret.Pix[i+0] = uint8(sum[j+0] / n)
			ret.Pix[i+1] = uint8(sum[j+1] / n)
			ret.Pix[i+2] = uint8(sum[j+2] / n)
			ret.Pix[i+3] = uint8(sum[j+3] / n)
		}
	}
	return ret
}

// resizeRGBA returns a scaled copy of the RGBA image slice r of m.
// The returned image has width w and height h.
func resizeRGBA(m *image.RGBA, r image.Rectangle, w, h int) image.Image {
	ww, hh := uint64(w), uint64(h)
	dx, dy := uint64(r.Dx()), uint64(r.Dy())
	// See comment in Resize.
	n, sum := dx*dy, make([]uint64, 4*w*h)
	for y := r.Min.Y; y < r.Max.Y; y++ {
		pix := m.Pix[(y-m.Rect.Min.Y)*m.Stride:]
		for x := r.Min.X; x < r.Max.X; x++ {
			// Get the source pixel.
			p := pix[(x-m.Rect.Min.X)*4:]
			r64 := uint64(p[0])
			g64 := uint64(p[1])
			b64 := uint64(p[2])
			a64 := uint64(p[3])
			// Spread the source pixel over 1 or more destination rows.
			py := uint64(y-r.Min.Y) * hh
			for remy := hh; remy > 0; {
				qy := dy - (py % dy)
				if qy > remy {
					qy = remy
				}
				// Spread the source pixel over 1 or more destination columns.
				px := uint64(x-r.Min.X) * ww
				index := 4 * ((py/dy)*ww + (px / dx))
				for remx := ww; remx > 0; {
					qx := dx - (px % dx)
					if qx > remx {
						qx = remx
					}
					qxy := qx * qy
					sum[index+0] += r64 * qxy
					sum[index+1] += g64 * qxy
					sum[index+2] += b64 * qxy
					sum[index+3] += a64 * qxy
					index += 4
					px += qx
					remx -= qx
				}
				py += qy
				remy -= qy
			}
		}
	}
	return average(sum, w, h, n)
}

// Resample returns a resampled copy of the image slice r of m.
// The returned image has width w and height h.
// plain old Nearest Neighbor algorithm
func Resample(m image.Image, r image.Rectangle, w, h int) image.Image {
	if w < 0 || h < 0 {
		return nil
	}
	if w == 0 || h == 0 || r.Dx() <= 0 || r.Dy() <= 0 {
		return image.NewRGBA64(r)
	}
	curw, curh := r.Dx(), r.Dy()
	img := image.NewRGBA(r)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Get a source pixel.
			subx := x * curw / w
			suby := y * curh / h
			r32, g32, b32, a32 := m.At(subx, suby).RGBA()
			r := uint8(r32 >> 8)
			g := uint8(g32 >> 8)
			b := uint8(b32 >> 8)
			a := uint8(a32 >> 8)
			img.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}
	return img
}
