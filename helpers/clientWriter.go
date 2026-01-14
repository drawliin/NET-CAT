package helpers

import (
	"fmt"
	"time"
)

func (c *Client) clientWriter() {
	for msg := range c.send {
		var prompt string
		var newLine string
		if !msg.history && msg.showPrompt {
			promptTime := time.Now().Format("2006-01-02 15:04:05")
			prompt = fmt.Sprintf("[%s][%s]:", promptTime, c.c_name)
			newLine = "\n"
		}
		_, err := c.c_conn.Write([]byte(newLine + msg.message + prompt))
		if err != nil {
			fmt.Printf("Error %v\n", err)
			return
		}
	}
}
