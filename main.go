package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	srcFile string
	level   LevelOption
	order   Order
)

func init() {
	flag.StringVar(&srcFile, "src", "./data/en.json", "Source file that contains irregular verbs")
	flag.Var(&level, "level", fmt.Sprintf("Language level. Allowed options: %s", strings.Join(levelOptions, ", ")))
	flag.Var(&order, "order", fmt.Sprintf("Quiz order. Allowed options: %s", strings.Join(orderOptions, ", ")))
	flag.Parse()
}

type Verb struct {
	Infinitiv   string      `json:"infinitiv"`
	Prateritum  string      `json:"prateritum"`
	Perfect     string      `json:"perfect"`
	Level       LevelOption `json:"level"`
	Translation string      `json:"translation"`
}

func (v Verb) Quiz(scanner *bufio.Scanner) {
	fmt.Printf("Verb: %s\n", v.Translation)

	entries := [][2]string{
		{v.Infinitiv, "Infinitiv"},
		{v.Prateritum, "Prateritum"},
		{v.Perfect, "Perfect"},
	}

	for _, entry := range entries {
		var line string
		for !strings.EqualFold(line, entry[0]) {
			fmt.Printf("%s: ", entry[1])
			if !scanner.Scan() {
				break
			}
			line = strings.TrimSpace(scanner.Text())
		}
	}
}

type IrregularVerbs []Verb

func NewIrregularVerbs(path string) (*IrregularVerbs, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var verbs IrregularVerbs
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&verbs); err != nil {
		return nil, err
	}

	return &verbs, nil
}

func (v *IrregularVerbs) VerbIndexes(l LevelOption, o Order) []int {
	selectedVerbIdxs := []int{}

	for idx, verb := range *v {
		if verb.Level == l {
			selectedVerbIdxs = append(selectedVerbIdxs, idx)
		}
	}

	if o == Random {
		rand.Shuffle(len(selectedVerbIdxs), func(i, j int) {
			selectedVerbIdxs[i], selectedVerbIdxs[j] = selectedVerbIdxs[j], selectedVerbIdxs[i]
		})
	}
	return selectedVerbIdxs
}

func (v *IrregularVerbs) Play(l LevelOption, o Order) {
	indexes := v.VerbIndexes(l, o)

	scanner := bufio.NewScanner(os.Stdin)
	for _, verbIndex := range indexes {
		verb := (*v)[verbIndex]
		verb.Quiz(scanner)
		fmt.Println()
	}
}

func main() {
	verbs, _ := NewIrregularVerbs(srcFile)
	verbs.Play(level, order)
}
