package main

import ("log"
        "net/http"
        "strings"
        "encoding/json")

type response struct {
  Answer    [][]string  `json:"answ"`
  Not_good  []string    `json:"notGood"`
}

func api(w http.ResponseWriter, r *http.Request) {
  log.Println("New Request API")
  db := redisDB{}
  db.init()
  defer db.deinit()

  cookie := r.URL.Query().Get("urls")
  data := strings.Split(cookie, "|")[1:]
  retur := response{}
  // clear cookie
  for _, key := range(data) {
    log.Println(db.exists(key))
    if db.exists(key) {
      retur.Answer = append(retur.Answer, []string {db.get(key), host+key, db.ttl(key), key})
    } else {
      retur.Not_good = append(retur.Not_good, key)
    }
  }
	js, _ := json.Marshal(retur)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
