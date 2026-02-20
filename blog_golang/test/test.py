
import requests

# 发送JSON数组
data = {
    "Username": "testuser1",
	"Email":    "testuser1@example.com",
	"Password": "password123",
	}

response = requests.post("http://localhost:9999" + "/api/v1/auth/register", json=data)
print(response.json())