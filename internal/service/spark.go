package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	hostUrl = "wss://spark-api.xf-yun.com/v1.1/chat"
	ctx     = gctx.New()
)

func genParams1(appid, question, ver string) map[string]interface{} { // 根据实际情况修改返回的数据结构和字段名
	messages := []Message{
		{Role: "user", Content: question},
	}
	domain := "general"
	if ver == "v2" {
		hostUrl = "wss://spark-api.xf-yun.com/v2.1/chat"
		domain = "generalv2"
	} else if ver == "v3" {
		hostUrl = "wss://spark-api.xf-yun.com/v3.1/chat"
		domain = "generalv3"
	} else if ver == "v3.5" {
		hostUrl = "wss://spark-api.xf-yun.com/v3.5/chat"
		domain = "generalv3.5"
	} else {
		domain = "generalv3.5"
	}
	data := map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
		"header": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"app_id": appid, // 根据实际情况修改返回的数据结构和字段名
		},
		"parameter": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"chat": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"domain":      domain,     // 根据实际情况修改返回的数据结构和字段名
				"temperature": 0.5,        // 根据实际情况修改返回的数据结构和字段名
				"top_k":       int64(1),   // 根据实际情况修改返回的数据结构和字段名
				"max_tokens":  int64(150), // 根据实际情况修改返回的数据结构和字段名
				"auditing":    "default",  // 根据实际情况修改返回的数据结构和字段名
			},
		},
		"payload": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"message": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"text": messages, // 根据实际情况修改返回的数据结构和字段名
			},
		},
	}

	return data // 根据实际情况修改返回的数据结构和字段名
}

// 创建鉴权url  apikey 即 hmac username
func assembleAuthUrl1(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		fmt.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	// fmt.Println(sgin)
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
	// fmt.Println(sha)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func Gen(input, ver string) (error, string, float64, float64, float64, float64) {
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	appidVar, _ := gcfg.Instance().Get(ctx, "spark.appid")
	apiKeyVar, _ := gcfg.Instance().Get(ctx, "spark.apiKey")
	apiSecretVar, _ := gcfg.Instance().Get(ctx, "spark.apiSecret")
	asId, _ := gcfg.Instance().Get(ctx, "spark.asId")
	q := input + "的性味以及不宜事项"
	//握手并建立websocket 连接
	if ver == "v2" {
		hostUrl = "wss://spark-api.xf-yun.com/v2.1/chat"
	} else if ver == "v3" {
		hostUrl = "wss://spark-api.xf-yun.com/v3.1/chat"
	} else if ver == "v3.5" {
		hostUrl = "wss://spark-api.xf-yun.com/v3.5/chat"
	} else if ver == "vswxg" {
		hostUrl = "wss://spark-openapi.cn-huabei-1.xf-yun.com/v1/assistants/" + asId.String()
		q = input
	}
	//log.Println("host", hostUrl)
	conn, resp, err := d.Dial(assembleAuthUrl1(hostUrl, apiKeyVar.String(), apiSecretVar.String()), nil)
	if err != nil {
		panic(readResp(resp) + err.Error())
		return err, "", 0, 0, 0, 0
	} else if resp.StatusCode != 101 {
		panic(readResp(resp) + err.Error())
	}

	go func() {

		//data := genParams1(appid, "角色设定：你是一位根据名字测试性格的大师\n目标任务：根据我提供的姓名进行分析性格特征\n需求说明：要求有理有据，分析内容积极向上，进行详细的分析解释\n风格设定：轻松愉快\n接下来我的输入是：{{猪八戒}}")
		data := genParams1(appidVar.String(), q, ver)
		log.Println(data)
		conn.WriteJSON(data)

	}()

	var answer = ""
	var qt, pt, tt, ct float64
	//获取返回的数据
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read message error:", err)
			break
		}

		var data map[string]interface{}
		err1 := json.Unmarshal(msg, &data)
		if err1 != nil {
			fmt.Println("Error parsing JSON:", err)
			return err1, "", 0, 0, 0, 0
		}
		//log.Printf("%#v", data)
		//fmt.Println(string(msg))
		//解析数据
		payload := data["payload"].(map[string]interface{})
		choices := payload["choices"].(map[string]interface{})
		header := data["header"].(map[string]interface{})
		code := header["code"].(float64)

		if code != 0 {
			//fmt.Println(data["payload"])
			return errors.New("parse data error"), "", 0, 0, 0, 0
		}
		status := choices["status"].(float64)
		//fmt.Println(status)
		text := choices["text"].([]interface{})
		content := text[0].(map[string]interface{})["content"].(string)
		if status != 2 {
			answer += content
		} else {
			//fmt.Println("收到最终结果")
			answer += content
			usage := payload["usage"].(map[string]interface{})
			temp := usage["text"].(map[string]interface{})
			//totalTokens := temp["total_tokens"].(float64)
			qt = temp["question_tokens"].(float64)
			pt = temp["prompt_tokens"].(float64)
			ct = temp["completion_tokens"].(float64)
			tt = temp["total_tokens"].(float64)
			//fmt.Println("total_tokens:", totalTokens)
			conn.Close()
			break
		}

	}
	return nil, answer, qt, pt, ct, tt
	//输出返回结果
	//fmt.Println(answer)

	//time.Sleep(1 * time.Second)
}
func SaveSpark(name, value, openid string, qt, pt, ct, tt float64) {
	db := g.DB("default")
	db.Model("usage").Data(g.Map{
		"user_input":     name,
		"spark_response": value,
		"qt":             qt,
		"pt":             pt,
		"ct":             ct,
		"tt":             tt,
		"openid":         openid,
	}).Insert()
}
