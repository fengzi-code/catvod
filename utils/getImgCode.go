package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func GetImgCode(imgUrl string) (imgCode string, cookies []*http.Cookie) {
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	cookies = res.Cookies()
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	f, err := os.OpenFile(`test.png`, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return
	} else {
		_, err = f.Write(b)
		if err != nil {
			return
		}
	}
	//url := "http://api.95man.com:8888/api/Http/Recog?Taken=foGqbzNHS9SXlt6reCPqyglUPVTU&imgtype=2&len=4"
	url := "http://119.23.145.51:8082//ocr/file/text"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("test.png")
	defer file.Close()
	part1,
		//errFile1 := writer.CreateFormFile("imgfile", filepath.Base("d:/test1.png"))
		errFile1 := writer.CreateFormFile("image", filepath.Base("test.png"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	imgCode = string(body)
	fmt.Println(imgCode)
	return imgCode, cookies
}
