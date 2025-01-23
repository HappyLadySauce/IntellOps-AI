package gf

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// http请求-通用
// method=请求方式(GET POST PUT DELETE)，url=请求地址，data=请求数据,headerList=请求头数据
func RequestHttp(method, url, data string, headerList map[string]string) (string, error) {
	payload := strings.NewReader(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}
	for key, val := range headerList {
		req.Header.Add(key, val)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// 发送GET请求
func HttpGet(url_text string, data map[string]interface{}) (map[string]interface{}, error) {
	u, err := url.Parse(url_text)
	if err != nil {
		log.Fatal(err)
	}
	paras := &url.Values{}
	//设置请求参数
	for k, v := range data {
		paras.Set(k, fmt.Sprintf("%v", v))
	}
	u.RawQuery = paras.Encode()
	resp, err := http.Get(u.String())
	//关闭资源
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, errors.New("request token err :" + err.Error())
	}
	jMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return nil, errors.New("request token response json parse err :" + err.Error())
	} else {
		return jMap, nil
	}

}

// 发送POST请求
func HttpPost(url_text string, urldata map[string]interface{}, postdata map[string]interface{}, contentType string) (map[string]interface{}, error) {
	u, err := url.Parse(url_text)
	if err != nil {
		log.Fatal(err)
	}
	paras := &url.Values{}
	//设置请求参数
	for k, v := range urldata {
		paras.Set(k, v.(string))
	}
	u.RawQuery = paras.Encode()
	//json序列化
	jsonData := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(jsonData)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(postdata); err != nil {
		return nil, errors.New("请求错误 :" + err.Error())
	}
	body := bytes.NewBufferString(string(jsonData.Bytes()))
	resp, erro := http.Post(u.String(), contentType, body)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if erro != nil {
		return nil, errors.New("请求错误 :" + erro.Error())
	}
	jMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return nil, errors.New(" 返回结果解析错误 :" + err.Error())
	} else {
		return jMap, nil
	}
}

// 请求失败返回
type Response struct {
	Code      int         `json:"code"`
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

// 返回错误
func Get_x(url string) (string, error) {
	// 超时时间：2秒
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}
	return result.String(), nil
}

/**
*  post 请求
*  data string
 */
func PostStrdata(url string, data string, contentType string) (string, error) {
	if contentType == "" {
		contentType = "application/json"
	}
	payload := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", contentType)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		return "", error
	}
	defer resp.Body.Close()
	result, _ := io.ReadAll(resp.Body)
	return string(result), nil
}

// tool
// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}

/**
* 发送POST请求
* url：         请求地址
* data：        POST请求提交的数据 interface{}
* contentType： 请求体格式，如：application/json
* content：     请求放回的内容
 */
func Post(url string, data interface{}, contentType string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := io.ReadAll(resp.Body)
	return string(result)
}
