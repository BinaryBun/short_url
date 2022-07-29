package main

import ("log"
        "net/http"
        "html/template"
        "time"
        "regexp"
        "math/rand"
        "fmt"

        "github.com/gorilla/mux")

type URL struct {
  Main_url    string
  Short_url   string
}
var global_data map[string]string = map[string]string {}  // sh - all

func randomString(length int) string {
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, length)
    rand.Read(b)
    log.Println(fmt.Sprintf("%x", b)[:length])
    return fmt.Sprintf("%x", b)[:length]
}

func set_cookie(w http.ResponseWriter, r *http.Request) {
  http.SetCookie( w, &http.Cookie{
      Name:       "url",
      Value:      r.FormValue("find"),
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

func home_page(w http.ResponseWriter, r *http.Request) {
  url_cook, err := r.Cookie("url")
  u := URL{Main_url: "http://loclhost:8080"}
  if err != nil {
    log.Println("Error occured while reading cookie")
	} else {
    u.Main_url = url_cook.Value
    short_url(u.Main_url, &u)
  }
  t, _ := template.ParseFiles("templace/index.html")
  t.Execute(w, u)
}

func redirect (w http.ResponseWriter, r *http.Request) {
  if r.FormValue("find") != "" {
    // regexp
    pattern := `^(http)|(https)://\w+\.\w{2,}`
    matched, _ := regexp.Match(pattern, []byte(r.FormValue("find")))
    if matched {
      log.Println(r.FormValue("find"))
    } else {
      log.Println("ERROR addr")
    }
  }
  set_cookie(w, r)
  http.Redirect(w, r, "/", http.StatusSeeOther)
}

func get_normal_url(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  //w.WriteHeader(http.StatusOK)
  value, flag := global_data[vars["id"]]
  if flag {
    http.Redirect(w, r, value, http.StatusSeeOther)
  } else {
    http.Redirect(w, r, "/", http.StatusSeeOther)
  }
}

func short_url(url string, u *URL) {
  rand := randomString(6)
  _, err := global_data[rand]
  for err {
    rand = randomString(6)
    _, err = global_data[rand]
  }
  global_data[rand] = url

  u.Short_url = fmt.Sprintf("http://localhost:8080/ref/%v", rand)
}

func pageHeaders() {
  http.Handle("/styles/",
             http.StripPrefix("/styles/",
                              http.FileServer(http.Dir("./styles/"))))

  rtr := mux.NewRouter()
  rtr.HandleFunc("/", home_page).Methods("GET")
  rtr.HandleFunc("/find/", redirect).Methods("POST")
  rtr.HandleFunc("/ref/{id}", get_normal_url).Methods("GET", "POST")

  http.Handle("/", rtr)  // перенаправление на роутер
  http.ListenAndServe(":8080", nil)
}

func main() {
  log.Println("==== Start ====")
  //global_data["vk"] = "https://vk.com/"
  pageHeaders()
}
