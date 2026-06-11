package response

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func ResponseStream(c *gin.Context, data string) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	intervals := getInternalTime(data)

	go func() {
		for _, char := range data {
			_, err := c.Writer.WriteString(fmt.Sprintf("data: %c\n\n", char))
			if err != nil {
				fmt.Println(err)
			}
			c.Writer.Flush()
			time.Sleep(intervals)
		}

		_, err := c.Writer.WriteString("data: \n\n")
		if err != nil {
			fmt.Println(err)
		}
		c.Writer.Flush()
	}()

	<-c.Writer.CloseNotify()
}

func getInternalTime(data string) time.Duration {
	if len(data) < 20 {
		return 200 * time.Millisecond
	}
	if len(data) < 100 {
		return 100 * time.Millisecond
	}
	if len(data) < 500 {
		return 50 * time.Millisecond
	}
	if len(data) < 5000 {
		return 20 * time.Millisecond
	}
	return 10 * time.Millisecond
}
