package helpers

import "fmt"

func (c *Client) clientWriter() {
	for msg := range c.send {
		_, err := c.c_conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("Error %v\n", err)
			return
		}
	}
}
