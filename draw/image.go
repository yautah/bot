package draw

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype"
	. "github.com/yautah/bot/data"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

const (
	fontFile = "./assets/font.ttf"
	fontSize = 20 // 字体尺寸
	fontDPI  = 72 // 屏幕每英寸的分辨率
)

func CreateDeckImg(teamDeck []Card, opponentDeck []Card) []byte {

	dst := image.NewRGBA(image.Rect(0, 0, 400, 240))
	blue := color.White
	draw.Draw(dst, dst.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	for i, card := range teamDeck {

		file, err := os.Open(fmt.Sprintf("./assets/cards/%s.png", card.Key))
		if err != nil {
			log.Fatal(err)
		}
		img, _ := png.Decode(file)
		defer file.Close()

		offsetX := 90*(i%4) + 10
		offsetY := i/4*110 + 10
		offset := image.Pt(offsetX, offsetY)
		draw.Draw(dst, img.Bounds().Add(offset), img, image.ZP, draw.Src)
	}

	// imga, _ := os.Open("F:\\cr\\after\\battle.png")
	// imgB, _ := png.Decode(imga)
	// defer imga.Close()
	// offsetVs := image.Pt(90*4+15, 30)
	// draw.Draw(dst, imgB.Bounds().Add(offsetVs), imgB, image.ZP, draw.Src)

	// for i, card := range opponentDeck {
	// fmt.Println(card.Key)
	// file, _ := os.Open(fmt.Sprintf("F:\\cr\\after\\%s.png", card.Key))
	// img, _ := png.Decode(file)
	// defer file.Close()

	// offsetX := 400 + 90*(i%4) + 10
	// offsetY := i/4*110 + 10
	// offset := image.Pt(offsetX, offsetY)
	// draw.Draw(dst, img.Bounds().Add(offset), img, image.ZP, draw.Src)
	// }

	emptyBuff := bytes.NewBuffer(nil) //开辟一个新的空buff
	png.Encode(emptyBuff, dst)

	return emptyBuff.Bytes()

}

func CreateChestImg(chest Chest) []byte {

	log.Println(chest)

	dst := image.NewRGBA(image.Rect(0, 0, 560, 540))
	blue := color.White
	draw.Draw(dst, dst.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println("读取字体数据出错")
		log.Println(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("转换字体样式出错")
		log.Println(err)
	}

	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(dst.Bounds())
	c.SetDst(dst)
	c.SetSrc(image.Black)

	for i, ct := range chest.Upcoming {

		file, _ := os.Open(fmt.Sprintf("F:\\cr\\chests\\chest-%s.png", ct))
		img, _ := png.Decode(file)
		defer file.Close()

		offsetX := 110*(i%4) + 10
		offsetY := i/4*130 + 12
		offset := image.Pt(offsetX, offsetY)
		draw.Draw(dst, img.Bounds().Add(offset), img, image.ZP, draw.Src)

		pt := freetype.Pt(offsetX+40, offsetY+115) // 字出现的位置
		_, err = c.DrawString(fmt.Sprintf("# %d", i), pt)
	}

	specials := []string{"giant", "magical", "epic", "legendary", "supermagical"}
	for k, v := range specials {
		file, _ := os.Open(fmt.Sprintf("./assets/chests/chest-%s.png", v))
		img, _ := png.Decode(file)
		defer file.Close()

		offsetX := 110*((k+9)%4) + 10
		offsetY := (k+9)/4*130 + 12
		offset := image.Pt(offsetX, offsetY)
		draw.Draw(dst, img.Bounds().Add(offset), img, image.ZP, draw.Src)

		pt := freetype.Pt(offsetX+30, offsetY+115) // 字出现的位置
		_, err = c.DrawString(fmt.Sprintf("# %d", chest.GetChestByKey(v)), pt)
	}

	emptyBuff := bytes.NewBuffer(nil) //开辟一个新的空buff
	png.Encode(emptyBuff, dst)

	return emptyBuff.Bytes()

}
