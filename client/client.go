package client

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/open-vector/open-jd-sdk/model"
	"github.com/open-vector/open-jd-sdk/util/jsonUtil"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type JdClient struct {
	Method    string `json:"method"`
	AppKey    string `json:"app_key"`
	Timestamp string `json:"timestamp"`
	V         string `json:"v"`
	Sign      string `json:"sign"`
	Param     string `json:"360buy_param_json"`
	SecretKey string `json:"secretKey"`
}

// New 实例化jdClient
func New(appKey string, secretKey string) JdClient {
	return JdClient{AppKey: appKey, V: "1.0", SecretKey: secretKey}
}

// Execute 发送http请求获取响应结果
func (jdClient JdClient) Execute(requestInterface model.RequestInterface, response interface{}) {
	url := "https://api.jd.com/routerjson"

	request := jsonUtil.GetByte(requestInterface)
	method := requestInterface.GetMethod()

	jdClient.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	jdClient.Method = method
	jdClient.Param = string(request)
	// 此处顺序不能乱
	jdClient.Sign = sign(getPreSignStr(jdClient))
	payload := strings.NewReader(getReader(jdClient))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	// 返回结果统一处理
	model.ResponseHandle(body, response, jdClient.Method)
}

// md5加密
func sign(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return strings.ToTitle(res)
}

// 获取签名前字符串
func getPreSignStr(jdClient JdClient) string {
	return fmt.Sprintf("%s360buy_param_json%sapp_key%smethod%stimestamp%sv%s%s",
		jdClient.SecretKey, jdClient.Param, jdClient.AppKey, jdClient.Method, jdClient.Timestamp, jdClient.V, jdClient.SecretKey)
}

// 构造payload
func getReader(jdClient JdClient) string {
	return fmt.Sprintf("app_key=%s&v=%s&method=%s&timestamp=%s&sign=%s&360buy_param_json=%s",
		jdClient.AppKey, jdClient.V, jdClient.Method, jdClient.Timestamp, jdClient.Sign, jdClient.Param)
}
