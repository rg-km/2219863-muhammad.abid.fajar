package main

<<<<<<< HEAD
import (
	"encoding/json"
)
=======
import "encoding/json"
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2

// Dari contoh yang telah diberikan, cobalah untuk melakukan decode dari json menjadi objek dari struct.
// Mengambil kasus pada encode, lengkapi function DecodeToLeaderboard agar json dapat di decode menjadi objek Leaderboard

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	// TODO: answer here
<<<<<<< HEAD
	Name  string `json:"name"`
	Email string `json:"-"`
	Rank  int    `json:"rank"`
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

type Leaderboard struct {
	Users []*UserRank
}

func DecodeToLeaderboard(jsonData []byte) (Leaderboard, error) {
	// TODO: answer here
<<<<<<< HEAD
	dumy := Leaderboard{}
	err := json.Unmarshal(jsonData, &dumy)
	if err != nil {
		return dumy, err
	} else {
		return dumy, nil
	}

=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
