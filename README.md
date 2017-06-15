本库用法：

	go get github.com/zouhuigang/font_width_height

头部导入

	import (
		ft "github.com/zouhuigang/font_width_height"
	)

使用：

		start_x, end_x := ft.GetImgWidth(img)//得到宽(end_x-start_x)
		start_y, end_y := ft.GetImgHeight(img)//得到高(end_y-start_y)

		out := ft.OutPutImg(img, start_x, start_y, (end_x - start_x), (end_y - start_y))//img为原始图片,输出图片，可保存


效果:

原始图：

![https://github.com/zouhuigang/font_width_height/raw/master/example/test1.png](https://github.com/zouhuigang/font_width_height/raw/master/example/test1.png)

效果图:

![https://github.com/zouhuigang/font_width_height/blob/master/example/out1.jpg](https://github.com/zouhuigang/font_width_height/blob/master/example/out1.jpg)



原始图：

![https://github.com/zouhuigang/font_width_height/raw/master/example/test.png](https://github.com/zouhuigang/font_width_height/raw/master/example/test.png)

效果图:

![https://github.com/zouhuigang/font_width_height/raw/master/example/out.jpg](https://github.com/zouhuigang/font_width_height/raw/master/example/out.jpg)