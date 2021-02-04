package tableimage

import (
	"bytes"
	"encoding/base64"
	"image"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//FileType the image format png or jpg
type FileType string

//TD a table data container
type TD struct {
	Text  string
	Color string
}

//TR the table row
type TR struct {
	BorderColor string
	Tds         []TD
}

type TableImage struct {
	width           int
	height          int
	th              TR
	trs             []TR
	backgroundColor string
	fileType        FileType
	filePath        string
	img             *image.RGBA
	firacode        font.Face
}

const (
	rowSpace         = 48
	tablePadding     = 14
	letterPerPx      = 24
	separatorPadding = 10
	wrapWordsLen     = 12
	columnSpace      = wrapWordsLen * letterPerPx
	// PNG is a png image format
	PNG FileType = "png"
	// JPEG is a jpg image format
	JPEG FileType = "jpg"
)

//Init initialise the table image receiver
func Init(backgroundColor string, fileType FileType, filePath string) TableImage {
	firacodeBin, _ := base64.StdEncoding.DecodeString(firacodeTTF)
	f, err := truetype.Parse(firacodeBin)
	if err != nil {
		log.Fatal(err)
	}
	ti := TableImage{
		backgroundColor: backgroundColor,
		fileType:        fileType,
		filePath:        filePath,
		firacode: truetype.NewFace(f, &truetype.Options{
			Size: 24,
			DPI:  96,
		}),
	}
	ti.setRgba()
	return ti
}

//AddTH adds the table header
func (ti *TableImage) AddTH(th TR) {
	ti.th = th
}

//AddTRs add the table rows
func (ti *TableImage) AddTRs(trs []TR) {
	ti.trs = trs
}

//Save saves the table
func (ti *TableImage) Save() {
	ti.calculateHeight()
	ti.calculateWidth()

	ti.setRgba()

	ti.drawTH()
	ti.drawTR()

	ti.saveFile()
}

// Get return table image in memory bytes
func (ti *TableImage) Get() *bytes.Buffer {
	ti.calculateHeight()
	ti.calculateWidth()

	ti.setRgba()

	ti.drawTH()
	ti.drawTR()

	return ti.get()
}
