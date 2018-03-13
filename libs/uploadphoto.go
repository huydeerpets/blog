package libs

import (
	"os"
	"strings"
	"image"
	_"image/jpeg"
	"image/png"
	"fmt"
	"code.google.com/p/graphics-go/graphics"
	"time"
	"strconv"
	"math/rand"
	"regexp"
)


//middle image
func MakeMilThumb(orgpath string,width int,height int) error {
	src, err := LoadImage(orgpath)
	if err != nil {
		return err
	}
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	err = graphics.Scale(dst, src)   //缩小图片

	if err != nil {
		return err
	}
	filepath_m := strings.Replace(orgpath, "_l.", "_m.",1)
	err = SaveImage(filepath_m, dst)
	return err
}

//small image
func MakeSmallThumb(orgpath string,width int,height int) error {
	src, err := LoadImage(orgpath)
	if err != nil {
		return err
	}
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	err = graphics.Scale(dst, src)   //缩小图片

	if err != nil {
		return err
	}
	filepath_s := strings.Replace(orgpath, "_l.", "_s.",1)
	err = SaveImage(filepath_s, dst)
	return err
}

//读取文件
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()
	img, _, err = image.Decode(file) //解码图片
	fmt.Println(err)
	return
}

//保存文件
func SaveImage(path string, img image.Image) (err error) {
	imgfile, err := os.Create(path)

	defer imgfile.Close()
	err = png.Encode(imgfile, img)   //编码图片
	return
}

//重新生成文件名
func FormatFileName(filename string) string {
	filepath := strings.Split(filename, ".")
	t := time.Now().UnixNano()
	rndNum := rand.Int63()
	t1 :=  strconv.FormatInt(t,10)+strconv.FormatInt(rndNum,10)
	newname := t1+"_l."+filepath[1]
	return newname
}

/**
 * 处理获取图片类型（大，中，小）
 * xxxx_l 原图   xxxx_m 中  xxx_s小
 */
func GetImageSizetype(path string, sizetype string) string{
	imgpath := path
	reg := regexp.MustCompile("(.*)_(.)(.*)")
	params := reg.FindStringSubmatch(imgpath)

	if(sizetype == "middle"){
		imgpath = params[1] + "_m" + params[3]
	}else if (sizetype == "small"){
		imgpath = params[1] + "_s" + params[3]
	}else{
		imgpath = path
	}
	fmt.Printf(imgpath)
	return imgpath
}

