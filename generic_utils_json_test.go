package general_goutils

import "testing"
import "github.com/google/uuid"

type TestData struct {
	Id    uuid.UUID `json:"id" gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	Name  string
	Email string
}

func TestGetData(t *testing.T) {
	s := getData(TestData{
		Id:    uuid.New(),
		Name:  "Daniel",
		Email: "daniel@mail.com",
	})
	println(s)
}
