package main

import (
  "fmt"
)

type University struct {
  Chosen1 string
  Chosen2 string
  Chosen3 string
}
type Abitur struct {
  StudentName string
  Score       int
  PassExam    bool
  DocsLength  int
  University  University
}
type Admin struct {
  IsAdmin bool
  Abitur
}

func NewAbitur(studentName string, score int, passExam bool, docsLength int, university University) *Abitur {
  return &Abitur{
    StudentName: studentName,
    Score:       score,
    PassExam:    passExam,
    DocsLength:  docsLength,
    University:  university,
  }
}
func UpdateName(a *Abitur, score int) {
  a.StudentName = "Almaz"
  a.Score = score
}
func (a *Admin) show() {
  fmt.Printf("admin: %s - %d\n", a.StudentName, a.Score)
}
func main() {
  address := University{"AITU", "ENU", "KarGUU"}
  user := NewAbitur("Anel", 106, true, 6, address)
  fmt.Println(user)
  UpdateName(user, 123)
  admin := Admin{true, *user}
  admin.show()
}
