/*
读取字体文件的大小
*/

package main

import (
	"fmt"
	ft "github.com/zouhuigang/font_width_height"
	"os"
)

func main() {

	img, err := ft.OpenImgDecode("test.png")
	if err != nil {
		fmt.Println(err)
	}

	start_x, end_x := ft.GetImgWidth(img)
	start_y, end_y := ft.GetImgHeight(img)
	fmt.Printf("start:[%v,%v] end:[%v,%v]\n", start_x, end_x, start_y, end_y)
	out := ft.OutPutImg(img, start_x, start_y, (end_x - start_x), (end_y - start_y))

	//保存图片
	w, _ := os.Create("out.jpg")
	defer w.Close()
	ft.SaveImage(w, out, 1)

}
