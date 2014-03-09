gae-go-dart-skeleton
====================

A skeleton project for GAE/Go + Dart

# Usage
1. Setup Google App Engine/Go environment.
2. Setup Dart enviroment.
3. Clone the repository.
4. Execute "pub get" inside "client" directory to get packages.
5. Execute "goapp serve" at the root of the repository to start a local server, and access http://localhost:8080/ with Dartium to check if it's working.
6. Change the application ID inside app.yaml to an ID that matches your own app ID, and execute "appcfg.py update --oauth2 gae-go-dart-skeleton" to upload to App Engine.

# About the source tree
## server
Contains the Go backend code for the server.

## client
Contains the Dart frontend code for the client.

### client/lib/components
Contains all web components used in the web app.
