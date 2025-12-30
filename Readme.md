### Sanitize user mssg (Prevent empty messages)
### When a client disconnected.. we should remove his conn from clients slice so other users can connect
### Validate Name (user can have up to 3 attempts when entering an invalid name)