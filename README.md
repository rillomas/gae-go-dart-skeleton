gae-go-dart-skeleton
====================

A skeleton project for GAE/Go + Dart + Polymer.dart.

You can see the actual website on App Engine at http://gae-go-dart-skeleton.appspot.com/

# Usage
1. Setup Google App Engine/Go environment.
2. Setup Dart enviroment.
3. Clone the repository.
4. Set GOPATH like the following 
   * For Windows: `<path to repository>\vendor;<path to repository>`
   * For Linux: `<path to repository>/vendor:<path to repository>`
5. Execute `goapp get google.golang.org/appengine` in repository root.
5. Execute `pub get` inside `src/web` directory to get packages.
6. Execute `goapp serve` at the repository root to start a local server, and access `http://localhost:8080/` with Dartium to check if it's working.
7. Execute `pub build` inside `src/web` directory to convert dart files to JavaScript. You will have to change the `app.yaml` handlers according to which files you want to serve with `goapp`. You will also have to change the root directory path in `main.go`.
   *  Handler for dev server

        ```
        - url: /(.*\.dart)
          static_files : web/web/\1
          upload: web/web/(.*\.dart)
        ```

   * Handler for deployed server (after pub build)

        ```
        - url: /(.*\.js)
          static_files : web/build/web/\1
          upload: web/build/web/(.*\.js)
        ```

   * root directory for dev server

        ```
          rootDirectory string = "web/web"
        ```

   * root directory for deployed server (after pub build)

        ```
          rootDirectory string = "web/build/web"
        ```

8. Change the application ID inside app.yaml to an ID that matches your own app ID, and execute `appcfg.py update .` inside `src` to upload to App Engine.

# About the source tree
## src
Contains the Go backend code for the server.

## src/web
Contains the Dart frontend code for the client.

## src/web/lib/components
Contains all web components used in the web app.

## vendor
Contains all external Go packages. The GOPATH must be set correctly for this directory to work properly. This directory is empty at the beginning, and external code + libraries will be created by `goapp` after `goapp get`.