# websocket

websocket is a websocket command line websocket client designed for exploring and debugging websocket servers. ws includes readline-style keyboard shortcuts, persistent history, and colorization.

## Installation

```bash

go get -u -v hidevops.io/websocket

```

## Usage

type `websocket -h` for help

```bash

websocket -h

websocket command line client

Usage:
  websocket [flags]

Flags:
  -h, --help           help for websocket
  -t, --token string   token. e.g. --token=the-token-string
  -u, --url string     websocket origin. e.g. --token=the-token-string
  -v, --version        print version, e.g. websocket version or websocket -v websocket --version

```

### Basic usage

#### Run websocket server

First, get the example server source code [here](https://hidevops.io/hiboot)

```bash
# get the source code
go get -u hidevops.io/hiboot

# change the working directory
cd $GOPATH/src/hidevops.io/hiboot/examples/web/websocket

# run the example
go run main.go

___  / / /__(_)__  /_______________  /_
__  /_/ /__  /__  __ \  __ \  __ \  __/
_  __  / _  / _  /_/ / /_/ / /_/ / /_     Hiboot Application Framework
/_/ /_/  /_/  /_.___/\____/\____/\__/     https://hidevops.io/hiboot

[INFO] 2018/11/30 14:46 Starting Hiboot web application hiboot on localhost with PID 59877
[INFO] 2018/11/30 14:46 Working directory: /Users/johnd/.gvm/pkgsets/go1.10/hidevops/src/hidevops.io/hiboot
[INFO] 2018/11/30 14:46 The following profiles are active: local, [websocket logging web]
[INFO] 2018/11/30 14:46 Initializing Hiboot Application
[INFO] 2018/11/30 14:46 Auto configure websocket starter on websocket.configuration
[INFO] 2018/11/30 14:46 Auto configure web starter on web.configuration
[INFO] 2018/11/30 14:46 Auto configure logging starter on logging.configuration
[INFO] 2018/11/30 14:46 Mapped "/websocket" onto controller.websocketController.Get()
[INFO] 2018/11/30 14:46 Mapped "/websocket/status" onto controller.websocketController.GetStatus()
[INFO] 2018/11/30 14:46 Add command line properties is enabled
[INFO] 2018/11/30 14:46 Hiboot started on port(s) http://localhost:8080
[INFO] 2018/11/30 14:46 Started hiboot in 0.007884 seconds

```

#### Connect to websocket server


```bash
websocket ws://localhost:8080/websocket/status

______  ____________             _____
___  / / /__(_)__  /_______________  /_
__  /_/ /__  /__  __ \  __ \  __ \  __/
_  __  / _  / _  /_/ / /_/ / /_/ / /_     Hiboot Application Framework
/_/ /_/  /_/  /_.___/\____/\____/\__/     https://hidevops.io/hiboot

> test
< Status: Up
>

```

For ssl connection, use `wss` protocol instead

```bash

websocket wss://your-host/websocket

```

### For jwt connection

You may want to connect to the server that request jwt token, how do we connect it?

First, get the example server source code [here](https://hidevops.io)

```bash
# get the source code
go get -u hidevops.io/hiboot

# change the working directory
cd $GOPATH/src/hidevops.io/hiboot/examples/web/jwt

# run the example
go run main.go

______  ____________             _____
___  / / /__(_)__  /_______________  /_
__  /_/ /__  /__  __ \  __ \  __ \  __/
_  __  / _  / _  /_/ / /_/ / /_/ / /_     Hiboot Application Framework
/_/ /_/  /_/  /_.___/\____/\____/\__/     https://hidevops.io/hiboot

[INFO] 2018/11/30 14:40 Starting Hiboot web application jwt-example on localhost with PID 59655
[INFO] 2018/11/30 14:40 Working directory: /Users/johnd/.gvm/pkgsets/go1.10/hidevops/src/hidevops.io/hiboot/examples/web/jwt
[INFO] 2018/11/30 14:40 The following profiles are active: local, [locale logging jwt websocket web]
[INFO] 2018/11/30 14:40 Initializing Hiboot Application
[INFO] 2018/11/30 14:40 Auto configure jwt starter on jwt.configuration
[INFO] 2018/11/30 14:40 Auto configure websocket starter on websocket.configuration
[INFO] 2018/11/30 14:40 Auto configure web starter on web.configuration
[INFO] 2018/11/30 14:40 Auto configure actuator starter on actuator.configuration
[INFO] 2018/11/30 14:40 Auto configure locale starter on locale.configuration
[INFO] 2018/11/30 14:40 Auto configure logging starter on logging.configuration
[INFO] 2018/11/30 14:40 Mapped "/foo" onto controller.fooController.Get()
[INFO] 2018/11/30 14:40 Mapped "/foo" onto controller.fooController.Post()
[INFO] 2018/11/30 14:40 Mapped "/health" onto actuator.healthController.Get()
[INFO] 2018/11/30 14:40 Mapped "/login" onto controller.loginController.Post()
[INFO] 2018/11/30 14:40 Mapped "/bar" onto controller.barController.Get()
[INFO] 2018/11/30 14:40 Mapped "/websocket" onto controller.websocketController.Get()
[INFO] 2018/11/30 14:40 Mapped "/websocket/status" onto controller.websocketController.GetStatus()
[INFO] 2018/11/30 14:40 Add command line properties is enabled
[INFO] 2018/11/30 14:40 Hiboot started on port(s) http://localhost:8080
[INFO] 2018/11/30 14:40 Started jwt-example in 0.011171 seconds

```

Once the server is running, try to get the token

```bash
http :8080/login 'username=someone' 'password=any-pwd'


HTTP/1.1 200 OK
Content-Length: 354
Content-Type: application/json; charset=UTF-8
Date: Fri, 30 Nov 2018 06:51:24 GMT
Set-Cookie: app.language=; Path=/; Expires=Fri, 30 Nov 2018 08:51:24 GMT; Max-Age=7200; HttpOnly

{
    "code": 200,
    "data": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDM1NjI0ODQsImlhdCI6MTU0MzU2MDY4NCwicGFzc3dvcmQiOiJhbnktcHdkIiwidXNlcm5hbWUiOiJzb21lb25lIn0.fg4vCKsy-0v-diKzXyOI_BjheGLL5CUjajcgz62hz-fM-wD7wkku3N7DPVfLJ6fwby7hr53iKxahyN7FdHlml-5Vg23s-DtIsgPMv0Juat8fRULmZ_BPDXdfcr9hRtbwYJhbYKUD5bgJ3kojME8ownjv0OWyqVPFADINN6N9h7s",
    "message": "Success"
}

```

Finally, copy the token above and connect to the server with the token

```bash

websocket ws://localhost:8080/websocket/status -t "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDM1NjI0ODQsImlhdCI6MTU0MzU2MDY4NCwicGFzc3dvcmQiOiJhbnktcHdkIiwidXNlcm5hbWUiOiJzb21lb25lIn0.fg4vCKsy-0v-diKzXyOI_BjheGLL5CUjajcgz62hz-fM-wD7wkku3N7DPVfLJ6fwby7hr53iKxahyN7FdHlml-5Vg23s-DtIsgPMv0Juat8fRULmZ_BPDXdfcr9hRtbwYJhbYKUD5bgJ3kojME8ownjv0OWyqVPFADINN6N9h7s"

______  ____________             _____
___  / / /__(_)__  /_______________  /_
__  /_/ /__  /__  __ \  __ \  __ \  __/
_  __  / _  / _  /_/ / /_/ / /_/ / /_     Hiboot Application Framework
/_/ /_/  /_/  /_.___/\____/\____/\__/     https://hidevops.io/hiboot

> test
< Status: Up
>

```


