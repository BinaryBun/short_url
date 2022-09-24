# URL Shortener
URL shortener based on GoLang and Redis. <br>
All data is stored in Redis. <br>
Client abbreviations are stored in cookies and are checked for TTL

## Running
 > Starting ```redis``` server <br>
``` bash
$ sudo redis-server
```

> Starting ```GoLand``` server
``` bash
$ go run shorter.go
```

## Сustomization
You can customize url length, TTL and host.
```Go
16. const deadline = 1*time.Hour  // your time
17. const host = "your url"  // format: "http://host/ref/"
18. const url_length = 6  // your length
```

## Interface
![Interface](./images/interface.png)

## Languages and Tools:
<p align="left"> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://redis.io" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/redis/redis-original-wordmark.svg" alt="redis" width="40" height="40"/> </a> </p>
