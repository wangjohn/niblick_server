package main

import (
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/zenazn/goji"
)

func main() {
  goji.Post("/v0/rounds", PostRounds)
  goji.Serve()
}

func PostRounds(w http.ResponseWriter, r *http.Request) {
  var postRoundRequest PostRoundRequest
  decoder := json.NewDecoder(r.Body)
  decodeErr := decoder.Decode(&postRoundRequest)
  if (decodeErr != nil) {
    // TODO(wangjohn): Write a wrapper that returns better validation errors
    http.Error(w, decodeErr.Error(), http.StatusBadRequest)
    return
  }

  roundLength := len(postRoundRequest.Holes)
  if (roundLength != 9 && roundLength != 18) {
    invalidRoundLengthMsg := fmt.Sprintf(
      "Invalid number of holes in a round. Must have either 9 or 18 holes, not %v",
      roundLength)
    http.Error(w, invalidRoundLengthMsg, http.StatusBadRequest)
    return
  }

  responseInterface := PostRoundResponse{"1"}
  responseData, encodeErr := json.Marshal(responseInterface)
  if (encodeErr != nil) {
    http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
    return
  }

  w.Write(responseData)
}

type PostRoundResponse struct {
  ID string
}

type PostRoundRequest struct {
  Holes []Hole
}

type Hole struct {
  Score int
  Par int
  Yardage int
  Handicap int
  GreenInReg bool
  Fairway bool
  Putts int
}
