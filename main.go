package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
 "regexp"
)

func playerInfoUrlFetcher(serverUrl string) []string {
  playerInfoUrlsRegex, _ := regexp.Compile(`http://xxlgamers.gameme.com/playerinfo/\d+`)
  res := playerInfoUrlsRegex.FindAllString(urlFetcher(serverUrl), -1)
  return uniq(res)
}

func serverUrlFetecher() []string {
  serverUrlRegex, _ := regexp.Compile(`http://xxlgamers.gameme.com/overview/\d+`)
  clanUrl := "http://xxlgamers.gameme.com/tf"

  // fetch page and parse
  res := serverUrlRegex.FindAllString(urlFetcher(clanUrl), -1)
  // unique results and save sever urls
  return uniq(res)
}

func uniq(s []string) []string {
  var seen bool
  uniqSlice := []string{}

  for i := range s {
    seen = false
    for j := range uniqSlice {
      if s[i] == uniqSlice[j] {
        seen = true
      }
    }
    if seen == false {
      uniqSlice = append(uniqSlice, s[i])
    }
  }
  return uniqSlice
}

func urlFetcher(url string) string {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  // if err yada yada???
  return string(body)
}

//channels
func sendPlayerInfoUrls(playerInfoUrls []string, cs chan string) {
  for i := range playerInfoUrls {
    cs <- playerInfoUrls[i]
  }
}

func recievePlayerInfoUrl(cs chan string) {

}

func main() {
  //severUrls := serverUrlFetecher()
  playerInfoUrls := playerInfoUrlFetcher("http://xxlgamers.gameme.com/overview/18")
  /* playerInfoIdChannel := make(chan string) */
  /* go sendPlayerInfoUrls(playerInfoUrls, playerInfoIdChannel) */
  //steamIdChannel := make(chan string)
  fmt.Printf("%v", playerInfoUrls)
}
