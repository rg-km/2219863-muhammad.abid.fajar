package spellchecker

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/ruangkata/downloader"
)

const url string = "https://gist.githubusercontent.com/fikriauliya/25c107ae057d3dc87abdb5dfb2f330b6/raw/639d54126b27c57651ac42ef8ece3c5c92a3d76b/en"
const filePath string = "./english-words.csv"

type spellchecker struct {
	words map[string]bool
}

func parseFile() (map[string]bool, error) {
	words := make(map[string]bool)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		words[word] = true
	}

	return words, nil
}

func NewEnglishSpellChecker() (SpellChecker, error) {
	err := downloader.Download(url, filePath)
	if err != nil {
		return nil, err
	}

	words, err := parseFile()
	if err != nil {
		return nil, err
	}

	return &spellchecker{words}, nil
}

func (s *spellchecker) CheckWord(word string) bool {
	var d []string
	c, _ := parseFile()
	for v, _ := range c {
		d = append(d, v)
	}
	for _, v := range d {
		if v == strings.ToLower(word) {
			return true
		}
	}
	return false // TODO: replace this
}

func (s *spellchecker) CheckSentence(sentence string) ([]string, []string) {

	dummy := strings.Split(sentence, " ")
	fmt.Println(dummy)
	var d []string
	validWords := make([]string, 0, 20)
	invalidWords := make([]string, 0, 20)
	c, _ := parseFile()
	for v, _ := range c {
		d = append(d, v)
	}
	for _, v := range dummy {
		for _, j := range d {
			if v == j {
				validWords = append(validWords, v)
			}
		}
	}

	for i, v := range dummy {
		for o, j := range validWords {
			if v != j && dummy[i] == dummy[o+(len(dummy)-len(validWords))] {
				invalidWords = append(invalidWords, v)
			}
		}
	}

	return validWords, invalidWords // TODO: replace this
}
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
