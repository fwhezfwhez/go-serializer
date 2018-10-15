package serializer

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	Serializer
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	Salary     float64   `json:"salary"`
	CreatedAt  time.Time `json:"created_at"`
}
func (u User) ToRepresentation(f func(Block)(func(b Block)Block,[]string))([]byte, error){
	return u.Serialize(u,f)
}
func TestUser_ToRepresentation(t *testing.T) {
	u := User{
		Name:      "ft2",
		Age:       9,
		Salary:    1000,
		CreatedAt: time.Now(),
	}

	buf, er := u.ToRepresentation(func(m Block)(func(b Block)Block,[]string){
		//pick :=[]string(nil)
		pick  := []string{"name","age"}
		handler := func(b Block)Block{
			b.Update("question","why?")
			b.Update("location","china")
			b.Update("created_at", u.CreatedAt.Format("2006-01-02 15:04:05"))
			b.Pop("location")
			return b
		}
		return handler,pick
	})
	if er!=nil {
		fmt.Println(er.Error())
		return
	}
	fmt.Println(string(buf))
}