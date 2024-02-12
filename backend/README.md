# HC Automation System

## Steps to run the project

1. Change secret.yml.template to secret.yml and enter the passwords for your email and db. \

*NOTE*: Use iitk email address, and just add your username, for e.g. : bmerchant22, not bmerchant22@iitk.ac.in

2. Run ```go mod tidy```

3. Run the project:
``` go run cmd/main.go cmd/auth.go```
*NOTE*: Configs are set for /backend as root, so don't run ```go run main.go auth.go``` in cmd/