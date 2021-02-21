package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	blackfriday "github.com/russross/blackfriday/v2"
	"github.com/signintech/gopdf"
)

var (
	fp = flag.String("i", "./Chinese.md", "input file")
	op = flag.String("o", "./Chinese.pdf", "output file")
)

func main() {
	flag.Parse()

	if _, err := os.Stat(filepath.Clean(*fp)); os.IsNotExist(err) {
		return
	}

	md, err := ioutil.ReadFile(filepath.Clean(*fp))
	if err != nil {
		log.Fatalln(err)
		return
	}

	unsafe := blackfriday.Run(md)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err = pdf.AddTTFFont("alib", "./Alibaba-PuHuiTi-Bold.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.Text(string(html))
	if err != nil {
		log.Panicln(err)
		return
	}
	err = pdf.WritePdf(filepath.Clean(*op))
	if err != nil {
		log.Panicln(err)
		return
	}
}
