package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
 "regexp"
 "strings"
)

func crawler(url string) string {
  regex, _ := regexDeterminer(url)
  if regex != nil {
    res := regex.FindAllString(urlFetcher(url), -1)
    urls := uniq(res)
    for i := range urls {
      crawler(urls[i])
    }
  }
  return url
}

func regexDeterminer(url string) (*regexp.Regexp, error) {
  if strings.Contains(url, "tf") {
    return regexp.Compile(`.gameme.com/overview/\d+`)
  } else if strings.Contains(url, "overview") {
    return regexp.Compile(`gameme.com/playerinfo/\d+`)
  } else if strings.Contains(url, "playerinfo") {
    return regexp.Compile(`http://steamcommunity.com/profiles/\d+`)
  }
  return nil, nil
}

func errorHandler(err error) {
  if err != nil {
    fmt.Println(err)
  }
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
  errorHandler(err)
  body, err := ioutil.ReadAll(resp.Body)
  errorHandler(err)
  resp.Body.Close()
  return string(body)
}

func main() {
  clans := []string {"xxlgamers", "db"}

  for i := range clans {
    clanUrl := "http://" + clans[i] + ".gameme.com/tf"
    fmt.Println(crawler(clanUrl))
  }
}
