package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomDescription() string {
	return RandomString(int(RandomInt(2, 5))) + " " + RandomString(int(RandomInt(2, 5))) + " " + RandomString(int(RandomInt(2, 5))) + " " + RandomString(int(RandomInt(2, 5)))
}

func RandomType() int32 {
	types := []int32{1, 3}
	n := len(types)
	return types[rand.Intn(n)]
}

func RandomEmail() string {
	return RandomString(6) + fmt.Sprint(RandomInt(1000, 9999)) + "@gmail.com"
}

func RandomGender() int32 {
	genders := []int32{0, 1}
	n := len(genders)
	return genders[rand.Intn(n)]
}

func RandomAvatar(gender int32) string {
	if gender == 0 {
		women := []string{"upload/avatar/avatar_women_1.jpg", "upload/avatar/avatar_women_2.jpg", "upload/avatar/avatar_women_3.jpg", "upload/avatar/avatar_women_4.jpg"}
		n := len(women)
		return women[rand.Intn(n)]
	} else {
		men := []string{"upload/avatar/avatar_men_1.jpg", "upload/avatar/avatar_men_2.jpg", "upload/avatar/avatar_men_3.jpg", "upload/avatar/avatar_men_4.jpg"}
		n := len(men)
		return men[rand.Intn(n)]
	}
}

func RandomURL() string {
	url := []string{"upload/avatar/avatar_men_1.jpg", "upload/avatar/avatar_men_2.jpg", "upload/avatar/avatar_men_3.jpg", "upload/avatar/avatar_men_4.jpg", "upload/avatar/avatar_women_1.jpg", "upload/avatar/avatar_women_2.jpg", "upload/avatar/avatar_women_3.jpg", "upload/avatar/avatar_women_4.jpg"}
	n := len(url)
	return url[rand.Intn(n)]
}
