package main

import "github.com/gin-gonic/gin"

func main() {
	{%- if web_framework==".gin" %}
	r := gin.Default()

	r.GET("/file_attachment", func(c *gin.Context) {
		c.FileAttachment("main.go", "main.md")
	})

	r.Run(":{{ http_port }}")
	{%- endif %}
}
