# HC Automation System

## Steps to run the project

1. Change secret.yml.template to secret.yml and enter the passwords for your email and db. \

*NOTE*: Use iitk email address, and just add your username, for e.g. : bmerchant22, not bmerchant22@iitk.ac.in

2. Run ```go mod tidy```

3. Run the project:
``` go run cmd/main.go cmd/auth.go cmd/appointment.go```
*NOTE*: Configs are set for /backend as root, so don't run ```go run main.go auth.go``` in cmd/ \

To test all the routes, you can use postman, this is the invite link to our workspace (if you don't have an account, do create it): https://app.getpostman.com/join-team?invite_code=ff5a70067344ce9abce913acbf956218&target_code=1cdc0e69034c0c96c1a2e8e61eb703bc 
