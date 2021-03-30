# Simple Conditional Server

## Setup

Three services have been setup :
(click on the links to open, run tests, build, and run the services)

1. [Auth Service](https://github.com/Avarsa/playground/tree/main/golang/simple-conditional-server/auth_service)  (user-authentication microservice)
2. [User Service](https://github.com/Avarsa/playground/tree/main/golang/simple-conditional-server/user_service)  (user microservice)
3. Landing Server  (proxy microservice)

## Checklist
(GET requests to these paths from root of the Landing Server)

1. - [x] /user/profile - secured - should respond with a data object related to the user
2. - [x] /user/profile - unsecured - should return with unauthorized status code
3. - [x] /microservice/name - should respond with the name of the microservice
 

