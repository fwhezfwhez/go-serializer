package main

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/go-serializer"
	"time"
)

type User struct {
	serializer.SerializerI
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Salary    float64   `json:"salary"`
	CreatedAt time.Time `json:"created_at"`
}

func (u User) ToRepresentation(f func(serializer.Block) (func(b serializer.Block) serializer.Block, []string)) ([]byte, error) {
	return u.Serialize(u, f)
}

func main() {
	u := User{
		SerializerI: serializer.JsonSerializer{},
		Name:      "ft2",
		Age:       9,
		Salary:    1000,
		CreatedAt: time.Now(),
	}

	buf, er := u.ToRepresentation(func(m serializer.Block) (func(b serializer.Block) serializer.Block, []string) {
		// abandon the 'Salary' field,use 'pick:= []string(nil)' to stand for all picked
		pick := []string{"age", "created", "name"}
		handler := func(b serializer.Block) serializer.Block {
			// update
			b.Update("question", "why?")
			b.Update("location", "china")
			// pop
			b.Pop("location")
			// rewrite
			b.Update("created_at", u.CreatedAt.Format("2006-01-02 15:04:05"))
			return b
		}
		return handler, pick
	})
	if er != nil {
		fmt.Println(er.Error())
		return
	}
	// the result of buf is a json []byte
	fmt.Println(string(buf))
	// if you want a further operator to the json map,get it by:
	var m map[string]interface{}
	json.Unmarshal(buf, &m)
	fmt.Println(m)
}
