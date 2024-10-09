package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	easyllm "github.com/soryetong/go-easy-llm"
	"github.com/soryetong/go-easy-llm/easyai"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})
	config := easyllm.DefaultConfig("sk-2d054090ec0e458d88b3a5233cdf56ba", easyai.ChatTypeQWen)

	client := easyllm.NewChatClient(config)
	r.GET("/getMbit", func(c *gin.Context) {

		resp, _, _ := client.NormalChat(context.Background(), &easyai.ChatRequest{
			Model:   easyai.ChatModelQWenTurbo,
			Message: "给我随机出一套完整mbti测试题目，选项为a和b,只需要告诉我题目和选项,只以json\"questions\": [ {\n      \"question\": \"在面对未知情况时，你\",\n      \"options\": {\n        \"A\": \"感到兴奋并寻找机会。\",\n        \"B\": \"感到不安并寻求确定性。\"\n      }\n    }]只需要questions这个json数组，不要其他任何多余的东西,不要任何markdown格式",
		})

		print(resp.Content)
		c.JSON(http.StatusOK, gin.H{
			"question": resp.Content,
		})
	})
	r.POST("/submitAnswers", func(c *gin.Context) {
		b, _ := c.GetRawData()
		var m map[string]interface{}
		// 反序列化
		_ = json.Unmarshal(b, &m)
		// 这里可以对 answerString 进行进一步处理
		//print(len(m), "sdfsdfsdf")
		ansStr := ""
		for k, v := range m {
			ansStr = ansStr + k + v.(string)
		}
		print(ansStr, "ansstr")

		resp, _, _ := client.NormalChat(context.Background(), &easyai.ChatRequest{
			Model:   easyai.ChatModelQWenTurbo,
			Message: ansStr + "这是我的选择，给我分析一下我的mbti",
		})
		c.JSON(200, gin.H{"message": resp.Content})
	})
	r.Run()
}
