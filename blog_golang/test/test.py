
import requests

# 发送JSON数组
data = {
    "Username": "testuser3",
	"Email":    "testuser3@example.com",
	"Password": "password123",
	}

# response = requests.post("http://localhost:9999" + "/api/v1/auth/register", json=data)

data = {
    "Username": "testuser3",
	"Password": "password123",
	}
#response = requests.post("http://localhost:9999" + "/api/v1/auth/login", json=data)


#print(response.json())

response = requests.post("http://localhost:9999" + "/api/v1/profile", 
                         headers={"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJ1c2VybmFtZSI6InRlc3R1c2VyMyIsImV4cCI6MTc3MTc2NDE0NywibmJmIjoxNzcxNjc3NzQ3LCJpYXQiOjE3NzE2Nzc3NDd9.bh4xvMCLZaUVErUJnbgvSgPUXQ0bllWPcXNMvqIgnCA" })
print(response.json())
