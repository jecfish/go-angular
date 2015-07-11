package main

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
  "net/http"
  "strings"
  "github.com/rs/cors"
  "github.com/zenazn/goji"
  "github.com/zenazn/goji/web"
)

func addFriend(friend string) {
  fileName := "data.txt"
  contents,_ := ioutil.ReadFile(fileName)
  result := string(contents) + "\n" + friend
  //println(string(contents))
  ioutil.WriteFile(fileName, []byte(result), 0x777)
}

func getFriends() ([]string)  {
  fileName := "data.txt"
  contents,_ := ioutil.ReadFile(fileName)
  result := strings.Split(string(contents),"\n")
  //fmt.Println(string(contents))
  return result
}

func removeFriend(joke string){
  jokes := getFriends();
  //fmt.Printf("%q\n", jokes)
  index := findFriend(joke, jokes)

  if index != 999 {
    jokes := append(jokes[:index],jokes[index+1:]...)
    //fmt.Printf("%q\n", jokes)

    fileName := "data.txt"
    ioutil.WriteFile(fileName, []byte(strings.Join(jokes,"\n")), 0x777)
  }
  //fmt.Printf("%v\n", index)

}

func findFriend(a string, list []string) int {
  for key, val := range list {
    if val == a {
      return key
    }
  }
  return 999
}

func post(c web.C, w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Add, %s!", c.URLParams["name"])
  addFriend(c.URLParams["name"])
}

func get(c web.C, w http.ResponseWriter, r *http.Request) {
  encoder := json.NewEncoder(w)
  encoder.Encode(getFriends())
}

func delete(c web.C, w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Remove, %s!", c.URLParams["name"])
  removeFriend(c.URLParams["name"])
}

func main() {
  //cross browser origin
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET","POST","DELETE",},
  })

  //routing
  goji.Use(c.Handler)
  goji.Get("/friend/", get)
  goji.Post("/friend/:name", post)
  goji.Delete("/friend/:name", delete)
  goji.Serve()
}
