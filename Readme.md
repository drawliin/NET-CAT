#### 1- Sanitize user mssg (Prevent empty messages)
#### 2- When a client disconnected.. we should remove his conn from clients slice so other users can connect
#### 3- Validate Name (user can have up to 3 attempts when entering an invalid name)
#### 4- Writing to connections without error handling
```
we ignore write errors everywhere:

c.c_conn.Write([]byte(msg))

If a client disconnects silently, this will keep failing.

maybe we can try:
if _, err := c.c_conn.Write([]byte(msg)); err != nil {
    // optional: remove client here
}
```