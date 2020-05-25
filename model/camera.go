package model

// A compact 3D camera with x, y, z position and direction
type Camera struct {
  Position  float64[] `json:"p"`
  Direction  float64[] `json:"d"`
}
