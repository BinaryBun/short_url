package main

import ("log"
        "fmt"
        "time"
        "regexp"
        "strings"
        "context"
        "net/http"
        "math/rand"
        "html/template"

        "github.com/gorilla/mux"
        "github.com/go-redis/redis")

const deadline = 1*time.Hour
const host = "http://localhost:8080/ref/"

type URL struct {
  URLs    [][3]string  // [main_url, short_url, TTL]
}

type redisDB struct {
  ctx   context.Context
  rdb   *redis.Client
}

func (r *redisDB) init() {
  r.ctx = context.Background()
  r.rdb = redis.NewClient(&redis.Options{ Addr: "localhost:6379" })
}

func (r *redisDB) set(key, value string) {
  if r.ctx == nil || r.rdb == nil { panic("Data is nil") }
  err := r.rdb.Set(r.ctx, key, value, deadline).Err()
  if err != nil { panic(err) }
}

func (r *redisDB) getKeys() []string {
  data, _ := r.rdb.Keys(r.ctx, "*").Result()
  return data
}

func (r *redisDB) get(key string) string {
  data, _ := r.rdb.Get(r.ctx, key).Result()
  return data
}

func (r *redisDB) exists(key string) bool {
  result, _ := r.rdb.Exists(r.ctx, key).Result()
  if result == 0 { return false }
  return true
}

func (r *redisDB) ttl(key string) string {
  result := fmt.Sprintf("%v", r.rdb.TTL(r.ctx, key))
  result = strings.Split(result, " ")[2]
  if result[1]!='h' && result[2]!='h' { result = "0:"+result }
  result = strings.Replace(result, "h", ":", -1)
  result = strings.Replace(result, "m", ":", -1)
  result = strings.Replace(result, "s", "", -1)

  return result
}

func (r *redisDB) deinit() {
  r.rdb.Close()
}

func getCookie(r *http.Request) string {
  urls_cook, err := r.Cookie("urls")
  if err != nil { return "" }
  log.Println("Cookie <OK>")
  return urls_cook.Value
}

func set_cookie(w http.ResponseWriter, r *http.Request, cookieValue string) {
  http.SetCookie( w, &http.Cookie{
      Name:       "urls",
      Value:      fmt.Sprintf("%v|%v", getCookie(r), cookieValue),
      Path:       "/",
      Domain:     "",
      Expires:    time.Time{},
      RawExpires: "",
      MaxAge:     0,
      Secure:     false,
      HttpOnly:   true,
      SameSite:   0,
      Raw:        "",
      Unparsed:   nil,} )
}

func clearCookie(w http.ResponseWriter, r *http.Request) []string  {
  db := redisDB{}
  db.init()
  defer db.deinit()

  // filter urls
  result := ""
  data := strings.Split(getCookie(r), "|")[1:]
  for _, val := range(data) {
    if db.exists(val) { result += fmt.Sprintf("|%v", val) }
  }

  // reset Cookie
  http.SetCookie( w, &http.Cookie{
      Name:       "urls",
      Value:      result,
      Path:       "/",
      Domain:     "",
      Expires:    time.Time{},
      RawExpires: "",
      MaxAge:     0,
      Secure:     false,
      HttpOnly:   true,
      SameSite:   0,
      Raw:        "",
      Unparsed:   nil,} )

    return strings.Split(result, "|")[1:]
}

func randomString(length int) string {
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, length)
    rand.Read(b)
    log.Println(fmt.Sprintf("%x", b)[:length])
    return fmt.Sprintf("%x", b)[:length]
}

func homePage(w http.ResponseWriter, r *http.Request) {
  db := redisDB{}
  db.init()
  defer db.deinit()

  cookies := clearCookie(w, r)
  // get short_url: [main_url, TTL]
  u := URL{}
  for _, key := range(cookies) {
    u.URLs = append(u.URLs, [3]string {db.get(key), host+key, db.ttl(key)})
    //u.URLs[host+key] = [2]string {db.get(key), db.ttl(key)}
  }

  t, _ := template.ParseFiles("templace/index.html")
  t.Execute(w, u)
}

func redirect(w http.ResponseWriter, r *http.Request) {
  db := redisDB{}
  db.init()
  defer db.deinit()

  if r.FormValue("find") != "" {
    pattern := `^((http)|(https))\:\/\/\w+\.\w{2,}`
    matched, _ := regexp.Match(pattern, []byte(r.FormValue("find")))
    if matched {
      // redis add hash-FormValue
      rand_str := randomString(6)
      for db.exists(rand_str) { rand_str = randomString(6) }
      db.set(rand_str, r.FormValue("find") )
      set_cookie(w, r, rand_str)
    } else {
      log.Println("ERROR addr")
    }
  }

  http.Redirect(w, r, "/", http.StatusSeeOther)
}

func get_normal_url(w http.ResponseWriter, r *http.Request) {
  db := redisDB{}
  db.init()
  defer db.deinit()

  id := mux.Vars(r)["id"]
  if db.exists(id) {
    http.Redirect(w, r, db.get(id), http.StatusSeeOther)
    //log.Println("Redirect <OK>")
  } else {
    http.Redirect(w, r, "/", http.StatusSeeOther)
  }
}

func pageHeaders() {
  http.Handle("/styles/",
             http.StripPrefix("/styles/",
                              http.FileServer(http.Dir("./styles/"))))

  rout := mux.NewRouter()
  rout.HandleFunc("/", homePage).Methods("GET")
  rout.HandleFunc("/find/", redirect).Methods("POST")
  rout.HandleFunc("/ref/{id}", get_normal_url).Methods("GET", "POST")

  http.Handle("/", rout)  // перенаправление на роутер
  http.ListenAndServe(":8080", nil)
}

func main() {
  log.Println("=== SERVER IS STARTED ===")

  pageHeaders()
}
