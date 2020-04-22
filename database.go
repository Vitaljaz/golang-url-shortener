package main

import "github.com/recoilme/slowpoke"

func SaveToDB(hash string, longURL string) {
	key := []byte(hash)
	val := []byte(longURL)

	if !CheckKey(key) {
		slowpoke.Set(file, key, val)
	}
}

func CheckKey(key []byte) bool {
	_, err := slowpoke.Get(file, key)

	if err != nil {
		return false
	} else {
		return true
	}
}

func GetFromDB(hash string) (string, error) {
	key := []byte(hash)
	longURL, err := slowpoke.Get(file, key)
	return string(longURL), err
}
