package client

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/open-vector/mallPromotie/backend/model"
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

func New() JdClient {
	return JdClient{AppKey: "2ac1c5db7e31d4a85d145ac19fa3d4e8", V: "1.0", SecretKey: "47c30eb69dbb422097842d039df4133e"}
}

func (jdClient JdClient) Execute(method string, request []byte) model.JingfenQueryResult {
	url := "https://api.jd.com/routerjson"

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

	var jfQueryResponse model.JFQueryResponse
	json.Unmarshal(body, &jfQueryResponse)

	var jingfenQueryResult model.JingfenQueryResult
	json.Unmarshal([]byte(jfQueryResponse.Body.QueryResult), &jingfenQueryResult)

	return jingfenQueryResult
}

func sign(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return strings.ToTitle(res)
}

func getPreSignStr(jdClient JdClient) string {
	return fmt.Sprintf("%s360buy_param_json%sapp_key%smethod%stimestamp%sv%s%s",
		jdClient.SecretKey, jdClient.Param, jdClient.AppKey, jdClient.Method, jdClient.Timestamp, jdClient.V, jdClient.SecretKey)
}

func getReader(jdClient JdClient) string {
	appKey := "2ac1c5db7e31d4a85d145ac19fa3d4e8"
	v := "1.0"
	return fmt.Sprintf("app_key=%s&v=%s&method=%s&timestamp=%s&sign=%s&360buy_param_json=%s",
		appKey, v, jdClient.Method, jdClient.Timestamp, jdClient.Sign, jdClient.Param)
}
