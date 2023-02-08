package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Data struct {
	User User `json:"data"`
}

type UserCreated struct {
	Name      string    `json:"name"`
	Job       string    `json:"job"`
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	DeleteAt  time.Time `json:"deletedAt,omitempty"`
}

const (
	base = "https://reqres.in/api"
)

func main() {
	fmt.Println("Que sucede?")
	r, e := GetExample(fmt.Sprintf("%s/users/2", base))
	if e != nil {
		fmt.Println("Error")
	}
	fmt.Println(string(r))
	b, e := Get(fmt.Sprintf("%s/users/2", base))
	if e != nil {
		fmt.Println("Error")
	}
	fmt.Println(string(b))

	us, er := GetUser("3")
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println(us)
	fmt.Println(us)
	fmt.Println("Id: ", us.Id)
	fmt.Println("FirstName: ", us.FirstName)
	fmt.Println("LastName: ", us.LastName)
	fmt.Println("Email: ", us.Email)
	fmt.Println()

	u, err := CreateUser("Dorys", "Prosesora")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	fmt.Println(u.Id)
	fmt.Println(u.Name)
	fmt.Println(u.Job)
	fmt.Println(u.CreatedAt)
}

func GetExample(url string) ([]byte, error) {
	response, err := http.Get(url) //response,err
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll((response.Body))
	if err != nil {
		return nil, err
	}

	return body, nil

}

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil) // metodo, url y body
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf(" stattus code %d ", resp.StatusCode)
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func GetUser(userId string) (*User, error) {
	b, e := Get(fmt.Sprintf("%s/users/%s", base, userId))
	if e != nil {
		return nil, e
	}
	var dataresponse Data
	if err := json.Unmarshal(b, &dataresponse); err != nil {
		return nil, err
	}
	return &dataresponse.User, nil
}

func Post(url string, data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application-json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf(" status code %d  ", resp.StatusCode)
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

func CreateUser(name, job string) (*UserCreated, error) {
	user := &UserCreated{Name: name, Job: job}
	req, err := Post(fmt.Sprintf("%s/users", base), user)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(req, user); err != nil {
		return nil, err
	}

	return user, nil
}
