package main

<<<<<<< HEAD
import (
	"encoding/json"
)
=======
import "encoding/json"
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2

// Dari contoh yang telah diberikan, cobalah untuk melakukan encode struct menjadi json.
// Lengkapi function EncodeToJson agar dapat mengembalikan nilai byte hasil dari encode objek Leaderboard.
// Modifikasi struct UserRank sehingga field Name menjadi name ,field Rank menjadi rank, dan field Email tidak ikut untuk diencode.

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	// TODO: answer here
<<<<<<< HEAD
	Name  string `json:"name"`
	Email string
	Rank  int `json:"rank"`
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

type Leaderboard struct {
	Users []*UserRank
}

func EncodeToJson(leaderboard Leaderboard) ([]byte, error) {
	// TODO: answer here
<<<<<<< HEAD

	test, err := json.Marshal(leaderboard)

	if err != nil {
		return test, err
	} else {
		return test, nil
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
