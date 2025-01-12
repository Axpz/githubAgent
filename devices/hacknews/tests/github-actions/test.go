package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

// GitHub Secret
const secret = "your-webhook-secret" // 请替换为你在 GitHub webhook 配置中的 secret

// PullRequestEvent 定义了 GitHub webhook 的 payload 中与 Pull Request 相关的数据结构
type PullRequestEvent struct {
	Action      string `json:"action"`
	PullRequest struct {
		Merged bool `json:"merged"`
		Base   struct {
			Ref string `json:"ref"`
		} `json:"base"`
	} `json:"pull_request"`
}

// PushEvent 定义了 GitHub webhook 的 payload 中与 Push 相关的数据结构
type PushEvent struct {
	Ref     string `json:"ref"`
	Before  string `json:"before"`
	After   string `json:"after"`
	Commits []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"commits"`
}

// 验证签名
func verifySignature(r *gin.Context) bool {
	// 获取 GitHub 发送的签名
	signature := r.GetHeader("X-Hub-Signature")
	if signature == "" {
		log.Println("No signature found")
		return false
	}

	// 读取请求体内容
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		log.Println("Error reading body:", err)
		return false
	}

	// 计算签名
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write(body)
	expectedSignature := "sha1=" + hex.EncodeToString(mac.Sum(nil))

	// 比较签名
	return subtle.ConstantTimeCompare([]byte(signature), []byte(expectedSignature)) == 1
}

// 处理 pull_request 事件
func handlePullRequestEvent(prEvent PullRequestEvent) {
	// 如果 PR 被合并并且目标分支是 main
	if prEvent.Action == "closed" && prEvent.PullRequest.Merged && prEvent.PullRequest.Base.Ref == "main" {
		fmt.Println("PR has been merged into the main branch!")
		// 这里可以添加你想要执行的操作
	}
}

// 处理 push 事件
func handlePushEvent(pushEvent PushEvent) {
	// 检查推送是否到主分支
	if pushEvent.Ref == "refs/heads/main" {
		// 如果是强制推送，before 和 after 提交哈希可能会有所不同
		if pushEvent.Before != "0000000000000000000000000000000000000000" {
			// 可以根据 before 和 after 字段判断是否是强制推送
			fmt.Println("Force push detected!")
		}
		// 打印提交信息
		for _, commit := range pushEvent.Commits {
			fmt.Printf("Commit ID: %s\n", commit.ID)
			fmt.Printf("Commit Message: %s\n", commit.Message)
		}
	}
}

// GitHub Webhook handler
func githubWebhookHandler(c *gin.Context) {
	// 创建日志记录
	klog.InitFlags(nil)
	defer klog.Flush()

	// 验证 GitHub 请求的签名
	if !verifySignature(c) {
		klog.Warningf("Invalid GitHub signature. Event: github, Status: unauthorized")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// 解析 GitHub 发送的 JSON 数据
	eventType := c.GetHeader("X-GitHub-Event")

	klog.Infof("Received GitHub event: %s, Content-Length: %d", eventType, c.Request.ContentLength)

	switch eventType {
	case "push":
		var pushEvent PushEvent
		if err := c.ShouldBindJSON(&pushEvent); err != nil {
			klog.Errorf("Invalid push event payload. Error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}
		handlePushEvent(pushEvent)

	case "pull_request":
		var prEvent PullRequestEvent
		if err := c.ShouldBindJSON(&prEvent); err != nil {
			klog.Errorf("Invalid pull request event payload. Error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}
		handlePullRequestEvent(prEvent)

	default:
		klog.Warningf("Unhandled GitHub event: %s", eventType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unhandled event"})
		return
	}

	// 返回成功响应
	klog.Infof("Webhook successfully received and processed")
	c.JSON(http.StatusOK, gin.H{"message": "Webhook received"})
}

func main() {
	// 创建 Gin 引擎
	r := gin.Default()

	// 设置 Webhook 端点
	r.POST("/github/webhook", githubWebhookHandler)

	// 启动服务器
	port := "8080"
	log.Printf("Starting server on :%s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
