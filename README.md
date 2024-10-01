how to use this server application

tools & language required:
1. Go languange
2. Docker Desktop / cli

steps to reproduce:
1. clone this repository
2. run docker compose build
3. run docker compose up
4. now the api url is available at http://localhost:3009/api/v1



API Documentation (Swagger)
 - http://localhost:3009/api/v1/swagger/index.html

Database UI (Adminer)
 - http://localhost:9001/ (can use the .env credentials to logged in to the database, make sure the server name is "db")

if you encounter issues when u run the command "docker compose up" containing texts "go.sum", try running "go mod tidy" in your project directory and try run "docker compose up" once again

if any errors still persisted during first build. try run:
 - docker compose build --no-cache && docker compose up


Known Bug:
- sometimes mysql is refused to connect and you had to run "docker compose build --no-cache && docker compose up" once again. This whole repository was rebuild based on PostgreSQL project.


Known Issue:
- due to lacks information about the technical test requirement. i only tested on my end using react native simulator and had my DeviceID or Device token registered on this Backend.
- You can use your own Device Token to make sure this whole system is worked. Please navigate to folder /fb/main.go line.45, change the token variable with working DeviceID token.

Possible Issue:
- due to security concern, github or any third party including firebase would not be able to publish any secrets to the public. whilst this repository are remain open to public. you can create your own firebase config and place in the project with the same name if possible.


Current System Flows:
1. once the /fcm/send (POST) API triggered, it will publish the rabbitmq to the queue and exchanges
2. after finished the rabbitmq publish, it will send to the Firebase Cloud Messaging.
3. after finished with FCM, it will save { identifier: string, deliveryAt: Date } to the database table "fcm_job"
4. you can enable rabbitmq message listener for the topic "notification.done" or "notification.fcm" if you enable on main.go line.60
5. finished and thanks
