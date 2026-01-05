#### 1- Sanitize user mssg (Prevent empty messages)
#### 2- Validate Name (user can have up to 3 attempts when entering an invalid name) + Unique Names
#### 3- Writing to connections without error handling
```
we ignore write errors everywhere:

c.c_conn.Write([]byte(msg))

If a client disconnects silently, this will keep failing.

maybe we can try:
if _, err := c.c_conn.Write([]byte(msg)); err != nil {
    // optional: remove client here
}
```
#### 4- Drop Clients When their buffer is full (it means they are slow enough)
#### 5- LImit User Name