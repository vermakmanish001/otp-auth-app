# otp-auth-app
Step-1
run the server using "go run main.go"

Step-2
testing using PostMan
Request Body:
{
  "phone": "1234567890",
  "device_id": "device123"
}
 Response:
{
  "message": "OTP sent"
}

login
Request Body:
{
  "phone": "1234567890",
  "otp": "123456",
}
{
  "message": "Login successful",
  "phone": "1234567890"
}

DEPLOYMENT STRATEGY

1. Containerization with Docker
Package the application into a Docker container for easy portability.
Deploy it using Docker Compose to manage MongoDB.
2️. Cloud Deployment Options
AWS EC2 : Deploy the Go app on an EC2 instance and use MongoDB Atlas for managing database.
Heroku: Use Heroku for deployment.
