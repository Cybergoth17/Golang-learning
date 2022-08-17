package main

import (
  "fmt"
)

//встраивание экомпозиций типов
type Adress struct {
  Street string
  City   string
}

// вместо классов структуры
type User struct {
  FirstName string
  LastName  string
  Email     string
  Adress    Adress
}

// в го не принято создавать объекты в потоке исполнения программы а принято создавать функции конструкторы которые создает объект
func NewUser(firstName, secondName, email string, adress Adress) *User {
  return &User{
    FirstName: firstName,
    LastName:  secondName,
    Email:     email,
    Adress:    adress,
  }
}

// эксперимент для создания функции, чтобы влиять на оригинальные поля а не работать с копиями
func UpdateName(u *User) {
  u.FirstName = "some new name"
}

// мы можем добавлять действия объектов с помощью методов объектов
func (u *User) Show() {
  fmt.Printf("%s - %s adress: %v\n", u.FirstName, u.LastName, u.Adress)
}
func (u *User) SetFirstName(name string) {
  u.FirstName = name
}

//функция встраивания
type Admin struct {
  IsAdmin bool
  User
}

func (a *Admin) Show() {
  fmt.Printf("admin: %s %s\n", a.FirstName, a.LastName)
}
func main() {
  adress := Adress{City: " Some City", Street: "Some street"}
  user := User{"Sergey", "Vasilev", "some@email.com", adress}
  user1 := NewUser("Sergey", "Second", "email@email.com", adress)
  fmt.Printf("user = %+v\n", user)
  fmt.Printf("user1 = %+v\n", user1)
  UpdateName(user1)
  fmt.Printf("updated user1 = %+v\n", user1)
  user1.Show()
  user1.SetFirstName("Obama")
  user1.Show()
  admin := Admin{IsAdmin: true, User: *user1}
  admin.Show()
}
