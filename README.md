# URL Shortener
URL shortener based on GoLang and Redis. <br>
All data is stored in Redis. <br>
Client abbreviations are stored in cookies and are checked for TTL

## Running

<table>
  <tr>
    <td>
      <b>Redis</b>
    </td>
    <td>

    ``` bash
    $ sudo redis-server
    ```

    </td>
  </tr>
  <tr>
    <td>
      <b>GoLang</b>
    </td>
    <td>

    Starting redis server

    ``` bash
    $ go run shorter.go
    ```
    </td>
  </tr>
</table>

## Interface
![Interface](./images/interface.png)

## Languages and Tools:
<p align="left"> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://redis.io" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/redis/redis-original-wordmark.svg" alt="redis" width="40" height="40"/> </a> </p>
