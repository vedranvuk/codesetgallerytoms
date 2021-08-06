package main

import (
	"io/ioutil"

	"github.com/vedranvuk/jsonobj"
)

func main() {
	var buf []byte
	var err error
	if buf, err = ioutil.ReadFile("/usr/lib/code/product.json"); err != nil {
		panic(err)
	}
	var obj *jsonobj.JSON
	if obj, err = jsonobj.Unmarshal(buf); err != nil {
		panic(err)
	}
	if err = obj.Set("extensionsGallery.serviceUrl", "https://marketplace.visualstudio.com/_apis/public/gallery"); err != nil {
		panic(err)
	}
	if buf, err = obj.Export("\t"); err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile("/usr/lib/code/product.json", buf, 0644); err != nil {
		panic(err)
	}
}