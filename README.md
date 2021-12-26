# Ws-load-test
The load testing library for websockets. Provide inbuild test strategies and allow users to develope own strategies for custom load tests

## Using ws-load-test

Clone this repository and move into this directory.

```shell
git clone git@github.com:angel-one/ws-load-test.git
cd ws-load-test
```

Resolve dependencies

```shell
go mod init
go mod tidy
````

Generate binaries

```shell
go install
go build
````

Binary named ws-load-test will be present in root directory. Run this with ws server details.

```shell
./ws-load-test --host 172.31.25.37:8080 --protocol ws --requestCount 1000 --gapTime 500 --path /mds --lifeTime 1 --strategy ping_pong --writeTime 1
```

Make sure metrices are generated
Swagger should be up http://localhost:8080/swagger/index.html#/
Send messages should be visible http://localhost:8080/send


## Runtime Arguments

--gapTime int               Number of milli seconds to wait in between establishing continuous ws connections. (default 100)

--host string               Domain name of test host. Ex 172.31.25.37:8080 (default "example.com")

--lifeTime int              The duration for which each connection remains (default 5)

--messageText string        The default text which need to be sent in case a basic strategy is used. This need to be passed in case of custom pr ping-pong strategy (default "test")

--path string               Specific url path (default "/somepath")

--port int                  application.yml port (default 8080). The http server.

--protocol string           Connection type. ws or wss. (default "wss")

--requestCount int          Total number of requests to be established through out load test. (default 6000)

--strategy string           Custom strategy which need to be injected. ping_pong in case of ping pong strategy. Can also contain a custom strategy

--writeTime int             The gap time between continuous writes.This is in seconds. (default 1)


## Metrics

All metrics are exposed through a HTTP endpoint as a PNG.

### Swagger
Swagger link : http://localhost:8080/swagger/index.html#/

![img.png](img.png)

### Metrics
All metrics are emitted aggregated per minute

1. Send Message Count (http://localhost:8080/send)
![img_3.png](img_3.png)

3. Receive Message Count (http://localhost:8080/receive)
![img_2.png](img_2.png)

4. Total connections Count (http://localhost:8080/connection)
![img_1.png](img_1.png)

5. Error Count (http://localhost:8080/error)

6. Latency (http://localhost:8080/latency)

7. Dashboard. Overall dashboard containing all metrics (http://localhost:8080/dashboard)

## Strategies

There are 3 type of load testes which are supported by this library. 2 of these are supported out of the box. For custom strategy/test, a strategy logic need to be injected into code.

### Basic Message Publish
This is used to send a custom periodic message to ws server with provided configurations.

```shell
./ws-load-test --host 172.31.25.37:8080 --protocol ws --requestCount 1000 --gapTime 500 --path /smart-stream --lifeTime 10 --strategy exchange_tick --writeTime 2
```
The above command will start test where a string "exchange_tick" will be send to ws server every 2 seconds. A total of 1000 requests will be created with a start up wait time of 500 ms between 2 established connections.

### Ping-Pong Strategy Publish
This is used to send a "ping" periodic message to ws server. This also handles the "pong" message which is sent from ws server.

```shell
./ws-load-test --host 172.31.25.37:8080 --protocol ws --requestCount 1000 --gapTime 500 --path /mds --lifeTime 1 --strategy ping_pong --writeTime 1
```
The above command will start test with ping pong strategy. A total of 1000 requests will be created with a start up wait time of 500 ms between 2 established connections. Ping message will sent to ws server every 1 second.

### Custom  Load Test Strategy
This is used to send a "ping" periodic message to ws server. This also handles the "pong" message which is sent from ws server.

```shell
./ws-load-test --host 172.31.25.37:8080 --protocol ws --requestCount 1000 --gapTime 500 --path /smart-stream --lifeTime 10 --strategy exchange_tick --writeTime 9999
```
Above example shows running a load test with custom strategy named exchange_tick. This strategy has a logic of sending a messge and receive stock markets ticks in  a streaming fashion. A total of 100 connections are made with a gap up time in 500 ms.


## Advanced Users/ Custom Strategy





   

