// send message post
const url_mask = "http://localhost:8080/api?";

function cleen_cookie(line) {
  document.cookie = document.cookie.replace(`|${line}`, "");
}

function elim_init(key, str) {
  data = document.getElementById(key)
  if (data !== null) {
    data.innerHTML = str;
  } else {
    // add 2 map && init in body
    console.log(`Add ${key}`);
    elim = document.createElement('p');
    elim.setAttribute("id", key);
    elim.innerHTML = str;
    document.body.append(elim);
  }
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

function json_work(js) {
  //console.log(js);
  if (js.answ !== null) {
    // меняем ttl
    for (i = 0; i < js.answ.length; i++) {
      var str = `TTL: ${js.answ[i][2]}<br>`;
      elim_init(js.answ[i][3], str);
    }
  }
  if (js.notGood !== null) {
    // удаляем ненужные элементы
    for (i = 0; i < js.notGood.length; i++) {
      console.log("DIV"+js.notGood[i]);
      document.getElementById("DIV"+js.notGood[i]).remove();
      cleen_cookie(js.notGood[i]);
    }
  }
}

function write() {
  //alert("JS starting!!")
  var url = url_mask + document.cookie
  fetch(url)
    .then(response => response.json())
    .then(json => json_work(json))
}
