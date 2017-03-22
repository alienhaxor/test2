# Description
Service that imports test1 and outputs each received command to STDOUT. 

# Usage
Use cURL to spam the server.
```
curl -H "Content-Type: application/json" -X POST -d '{"type":1, "body": "jesus 666"}' 127.0.0.1:9999/cmd --verbose
```
