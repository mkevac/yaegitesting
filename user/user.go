package user

import "math/rand"

const (
	GenderMale Gender = iota
	GenderFemale
	GenderUnknown
	GenderMax
)

type Gender uint8

func GetRandomGender() Gender {
	return Gender(rand.Intn(int(GenderMax)))
}

type User struct {
	UserID           uint32
	Gender           Gender
	Age              uint16
	CustomParameters []CustomParameter
}

func NewRandomUser() *User {
	return &User{
		UserID:           rand.Uint32(),
		Gender:           GetRandomGender(),
		Age:              uint16(rand.Int31n(50) + 18),
		CustomParameters: GetRandomCustomParameterList(uint(rand.Intn(20))),
	}
}

type CustomParameter struct {
	id    uint32
	value int32
}

func GetRandomCustomParameter() CustomParameter {
	return CustomParameter{
		id:    rand.Uint32(),
		value: rand.Int31(),
	}
}

func GetRandomCustomParameterList(count uint) []CustomParameter {
	list := make([]CustomParameter, 0, count)
	for i := 0; i < rand.Intn(20); i++ {
		list = append(list, GetRandomCustomParameter())
	}
	return list
}
