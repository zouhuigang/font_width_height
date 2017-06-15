/*
读取字体文件的大小
作者:邹慧刚
联系方式：952750120@qq.com
*/

package font_width_height

import (
	"fmt"
	"image/draw"
	//_ "image/jpeg" //解析图片时需要用到，导入但是不使用
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

const (
	//图片格式
	ImageFormatPng  = iota //0
	ImageFormatJpeg        //1
	ImageFormatGif         //2
)

//返回图片的宽，高
//	b := img.Bounds()
//	width, height := ImgWidthHeight(b)
func imgWidthHeight(b image.Rectangle) (int, int) {
	min, max := b.Min, b.Max
	height, width := max.Y-min.Y, max.X-min.X
	return width, height
}

//打开图片并且解码,支持gif,jpg,png格式,需要导入图片库image/png,image/jpeg,image/gif
func OpenImgDecode(imageFileName string) (image.Image, error) {
	reader, err := os.Open(imageFileName)
	if err != nil {
		fmt.Println("打开图片失败,自动跳过")
		return nil, err
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println("解析失败")
		return nil, err
	}
	return img, nil
}

//得到宽,start_x,end_x
func GetImgWidth(img image.Image) (start_x int, end_x int) {
	b := img.Bounds()
	width, height := imgWidthHeight(b)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			oldPixel := img.At(x, y)
			r, g, b, _ := oldPixel.RGBA() //uint32
			rf, gf, bf := float64(r>>8), float64(g>>8), float64(b>>8)

			if rf == gf && rf == bf { //白色或黑色背景,可能会有其他rgb相同的出现
				continue
			}
			if start_x == 0 {
				start_x = x
			} else {
				end_x = x
			}

		}
	}

	return
}

//start_y,end_y
func GetImgHeight(img image.Image) (start_y int, end_y int) {
	b := img.Bounds()
	width, height := imgWidthHeight(b)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			oldPixel := img.At(x, y)
			r, g, b, _ := oldPixel.RGBA() //uint32
			rf, gf, bf := float64(r>>8), float64(g>>8), float64(b>>8)

			if rf == gf && rf == bf { //白色或黑色背景,可能会有其他rgb相同的出现
				continue
			}
			if start_y == 0 {
				start_y = y
			} else {
				end_y = y
			}

		}
	}

	return
}

//输出图像,将img拷贝一部分输出
func OutPutImg(img image.Image, start_x, start_y, width, height int) image.Image {

	out := image.NewRGBA(image.Rect(0, 0, width, height)) //输出图片

	img_start_point := image.Point{start_x, start_y} //开始拷贝的点

	//待拷贝的形状
	dp := image.Point{0, 0}
	sr := image.Rect(0, 0, width, height)               //图片大小
	copy_area := image.Rectangle{dp, dp.Add(sr.Size())} //是图片绘图区域

	//拷贝
	draw.Draw(out, copy_area, img, img_start_point, draw.Src)

	return out
}

//保存图片对象,/将图片保持到输出流种，可以是文件或HTTP流等
func SaveImage(w io.Writer, img image.Image, imageFormat int) error {
	//f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0622)
	if imageFormat == ImageFormatPng {
		return png.Encode(w, img)
	}
	if imageFormat == ImageFormatJpeg {
		return jpeg.Encode(w, img, &jpeg.Options{100})
	}
	if imageFormat == ImageFormatGif {
		return gif.Encode(w, img, &gif.Options{NumColors: 256})
	}

	return errors.New("Not supported image format")
}
