package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {

}
