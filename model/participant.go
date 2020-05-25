package model

type Participant struct {
  Uid   int `json:"uid"`
  Pid   int `json:"pid"`
  Cid   int `json:"cid"`
  Name  string `json:"name"`
}
