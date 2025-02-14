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

We can deploy this application using docker or AWS EC2 or on Heroku CLI
