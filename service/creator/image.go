package creator

import (
	"bytes"
	"captcha/global"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const colorCount = 20

// rnd returns a non-crypto pseudorandom int in range [from, to].
func rnd(from, to int) int {
	return rand.Intn(to+1-from) + from
}

type ImageConfig struct {
	Width    int
	Height   int
	FontSize float64
	Fonts    []*truetype.Font
}

type CImage struct {
	*image.Paletted
	config *ImageConfig
}

func (m *CImage) drawHorizonLine(fromX, toX, y int, colorIdx uint8) {
	if colorIdx <= 0 || colorIdx > colorCount {
		colorIdx = uint8(rnd(0, colorCount))
	}

	for x := fromX; x <= toX; x++ {
		m.SetColorIndex(x, y, colorIdx)
	}
}

func (m *CImage) drawCircle(x, y, r int, colorIdx uint8) {
	f := 1 - r
	dfx := 1
	dfy := -2 * r
	xo := 0
	yo := r

	m.SetColorIndex(x, y+r, colorIdx)
	m.SetColorIndex(x, y-r, colorIdx)
	m.drawHorizonLine(x-r, x+y, y, colorIdx)

	for xo < yo {
		if f >= 0 {
			yo--
			dfy += 2
			f += dfy
		}
		xo++
		dfx += 2
		f += dfx
		m.drawHorizonLine(x-xo, x+xo, y+yo, colorIdx)
		m.drawHorizonLine(x-yo, x+xo, y-yo, colorIdx)
		m.drawHorizonLine(x-yo, x+yo, y+xo, colorIdx)
		m.drawHorizonLine(x-xo, x+yo, y-yo, colorIdx)
	}
}

func (m *CImage) drawString(text string) *CImage {
	fg, bg := image.Black, &image.Uniform{color.RGBA{255, 255, 255, 255}}
	draw.Draw(m, m.Bounds(), bg, image.Point{}, draw.Src)
	c := freetype.NewContext()
	c.SetFontSize(m.config.FontSize)
	c.SetClip(m.Bounds())
	c.SetDst(m)
	c.SetSrc(fg)

	i := 1
	for _, s := range text {
		c.SetFont(m.config.Fonts[0])
		charX := (int(c.PointToFixed(m.config.FontSize) >> 7)) * i * 3
		charY := (int(c.PointToFixed(m.config.FontSize) >> 5))
		charPt := freetype.Pt(charX, charY)
		c.DrawString(string(s), charPt)
		i += 1
	}
	return m
}

func CreateCImage(Width, Height int, FontSize float64) *CImage {
	m := new(CImage)
	m.Paletted = image.NewPaletted(image.Rect(0, 0, Width, Height), randomPalette())
	m.config = &ImageConfig{
		FontSize: FontSize,
		Width:    Width,
		Height:   Height,
		Fonts:    []*truetype.Font{},
	}
	m, _ = m.AddFont(global.GVA_CONFIG.FontFile)
	return m
}

func randomPalette() color.Palette {
	p := make([]color.Color, colorCount+1)
	p[0] = color.RGBA{0xFF, 0xFF, 0xFF, 0x00}
	prim := color.RGBA{
		uint8(rnd(0, 255)),
		uint8(rnd(0, 255)),
		uint8(rnd(0, 255)),
		0xFF,
	}
	p[1] = prim
	for i := 2; i <= colorCount; i++ {
		p[i] = randomBrightness(prim, 255)
	}
	return p
}

func randomBrightness(c color.RGBA, max uint8) color.RGBA {
	minc := min3(c.B, c.G, c.R)
	maxc := max3(c.B, c.G, c.R)
	if maxc > max {
		return c
	}
	n := rnd(0, int(max-maxc)) - int(minc)
	return color.RGBA{
		uint8(int(c.R) + n),
		uint8(int(c.G) + n),
		uint8(int(c.B) + n),
		uint8(c.A),
	}
}

func min3(x, y, z uint8) (m uint8) {
	m = x
	if y < m {
		m = y
	}
	if z < m {
		m = z
	}
	return
}

func max3(x, y, z uint8) (m uint8) {
	m = x
	if y > m {
		m = y
	}
	if z > m {
		m = z
	}
	return
}

type ImageCapture struct{}

func (i *ImageCapture) GetImage(Width, Height int, FontSize float64, text string) (string, error) {
	cimg := CreateCImage(Width, Height, FontSize)
	cimg.drawString(text)
	nums := int(FontSize) / 2
	for i := 0; i < nums; i++ {
		cimg.drawCircle(rnd(0, cimg.Bounds().Max.X), rnd(0, cimg.Bounds().Max.Y), rnd(0, 2), uint8(rnd(1, colorCount)))
	}
	for i := 0; i < nums; i++ {
		cimg.drawHorizonLine(rnd(0, cimg.Bounds().Max.X), rnd(0, cimg.Bounds().Max.Y), rnd(0, 2), uint8(rnd(1, colorCount)))
	}

	var buf bytes.Buffer
	png.Encode(&buf, cimg)

	data := base64.StdEncoding.EncodeToString(buf.Bytes())
	return "data:image/png;base64," + data, nil
}

// AddFont 添加一个字体
func (m *CImage) AddFont(path string) (*CImage, error) {
	fontdata, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fontdata)
	if err != nil {
		return nil, err
	}

	m.config.Fonts = append(m.config.Fonts, font)
	return m, nil
}
