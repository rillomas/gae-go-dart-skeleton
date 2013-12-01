gae-go-dart-skeleton
====================

A skeleton project for GAE/Go + Dart

# Usage
1. Setup Google App Engine/Go environment.
2. Setup Dart enviroment.
3. Clone the repository.
4. Execute "pub install" at the repository's root directory to get packages.
5. Execute "dev_appserver.py --port=8081 gae-go-dart-skeleton" to start a local server, and access http://localhost:8081/ with Dartium to check if it's working.
6. Change the application ID inside app.yaml to an ID that matches your own app ID, and execute "appcfg.py update --oauth2 gae-go-dart-skeleton" to upload to App Engine.

# About the source tree
## server
Contains the Go backend code for the server.

## web
Contains the Dart frontend code for the client.

### web/components
Contains all web components used in the web app.
