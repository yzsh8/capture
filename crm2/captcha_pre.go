// captcha 预处理
package crm2

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

const CAPURL = `http://om.jd.com/auth_authImg`

var (
	white = color.RGBA{255, 255, 255, 255}
	black = color.RGBA{0, 0, 0, 255}
)

// 验证码去则色输出 阈值..34000
func ImageClean(in, out string, threshole int) error {
	f, err := os.Open(in)
	if err != nil {
		return err
	}
	defer f.Close()
	o, err := os.Create(out)
	if err != nil {
		return err
	}
	defer o.Close()

	m, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	bounds := m.Bounds()
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	// 去除边缘
	minX, maxX, minY, maxY := bounds.Min.X+1, bounds.Max.X-1, bounds.Min.Y+1, bounds.Max.Y-1
	img := image.NewNRGBA(image.Rect(0, 0, maxX-1, maxY-1))
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			_, _, b, _ := m.At(x, y).RGBA()
			if b < uint32(threshole) {
				img.Set(x-1, y-1, black)
			} else {
				img.Set(x-1, y-1, white)
			}
		}
	}

	opt := &jpeg.Options{Quality: 100}
	err = jpeg.Encode(o, img, opt)
	if err != nil {
		return err
	}
	return nil
}

// 统计图像色彩信息,可以用excel做直方图,根据色彩范围确定验证码的阈值
func ImageColorInfo(in, out string) error {
	f, err := os.Open(in)
	if err != nil {
		return err
	}
	defer f.Close()

	m, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	bounds := m.Bounds()
	// 去除边缘
	minX, maxX, minY, maxY := bounds.Min.X+1, bounds.Max.X-1, bounds.Min.Y+1, bounds.Max.Y-1

	s := make(map[uint32]uint32)
	for y := minY; y < maxY; y++ {
		for x := minX + 1; x < maxX; x++ {
			//r, g, b, _ := m.At(x, y).RGBA()
			_, _, b, _ := m.At(x, y).RGBA()
			if _, ok := s[b]; ok {
				s[b] += 1
			} else {
				s[b] = 1
			}
		}
	}

	o, err := os.Create(out)
	if err != nil {
		return err
	}
	defer o.Close()
	for k, v := range s {
		fmt.Fprintf(o, "%d,%d\n", k, v)
	}
	return nil
}
