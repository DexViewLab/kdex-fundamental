package share

import (
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

// HTTPGet ..
func HTTPGet(method string, url string, httpClient *http.Client) (string, error) {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	resp, err := httpClient.Do(req)

	if err != nil {
		err := fmt.Errorf("HTTP get failed. err = %v, url = %s", err, url)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("HTTP read body failed. err = %v, url = %s", err, url)
		return "", err
	}
	bodyStr := string(body)
	if len(bodyStr) > 50 {
		glog.V(5).Infof("HTTP get %s OK.\n Body is %s ...\n", url, bodyStr[0:50])
	} else {
		glog.V(5).Infof("HTTP get %s OK.\n Body is %s\n", url, bodyStr)
	}
	return bodyStr, nil
}
