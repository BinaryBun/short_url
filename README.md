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

    Starting redis server
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
