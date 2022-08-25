package middleware

import "github.com/gin-gonic/gin"

func CrosMiddleWare(c *gin.Context) {
	origin := c.Request.Header.Get("origin")
	if len(origin) == 0 {
		origin = c.Request.Header.Get("Origin")
	}
	//该字段是必须的。它的值要么是请求时Origin字段的值，要么是一个*，表示接受任意域名的请求
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)

	//该字段可选。它的值是一个布尔值，表示是否允许发送Cookie。默认情况下，Cookie不包括在CORS请求之中。
	//设为true，即表示服务器明确许可，Cookie可以包含在请求中，一起发给服务器。这个值也只能设为 true，
	//如果服务器不要浏览器发送Cookie，删除该字段即可。
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	//如果浏览器请求包括Access-Control-Request-Headers字段，则Access-Control-Allow-Headers字段是必需的。
	//它也是一个逗号分隔的字符串，表明服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段。
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	//该字段必需，它的值是逗号分隔的一个字符串，表明服务器支持的所有跨域请求的方法。
	//注意，返回的是所有支持的方法，而不单是浏览器请求的那个方法。这是为了避免多次"预检"请求。
	c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()

}
