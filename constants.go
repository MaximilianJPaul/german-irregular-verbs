package main

import (
	"fmt"
	"strings"
)

type LevelOption int

const (
	A1 = iota
	A2
	B1
	B2
	C1
	C2
	ALL
)

var levelOptions = []string{"A1", "A2", "B1", "B2", "C1", "C2", "ALL"}

func ParseLevel(s string) (LevelOption, error) {
	for i, name := range levelOptions {
		if strings.EqualFold(s, name) {
			return LevelOption(i), nil
		}
	}
	return 0, fmt.Errorf("invalid level %q (want %s)", s, strings.Join(levelOptions, ", "))
}

func (l *LevelOption) String() string {
	if *l < 0 || int(*l) >= len(levelOptions) {
		return fmt.Sprintf("Level(%d)", *l)
	}
	return levelOptions[*l]
}

func (l *LevelOption) Set(s string) error {
	for i, name := range levelOptions {
		if strings.EqualFold(s, name) {
			*l = LevelOption(i)
			return nil
		}
	}
	return fmt.Errorf("must be one of %s", strings.Join(levelOptions, ", "))
}

func (l *LevelOption) MarshalText() ([]byte, error) {
	if *l < 0 || int(*l) >= len(levelOptions)-1 {
		return nil, fmt.Errorf("invalid level %d", int(*l))
	}
	return []byte(levelOptions[*l]), nil
}

func (l *LevelOption) UnmarshalText(b []byte) error {
	v, err := ParseLevel(string(b))
	if err != nil {
		return err
	}
	*l = v
	return nil
}

type Order int

const (
	Random Order = iota
	Alphabetical
)

var orderOptions = []string{"random", "alphabetical"}

func (o *Order) String() string {
	if *o < 0 || int(*o) >= len(orderOptions) {
		return fmt.Sprintf("Order(%d)", *o)
	}
	return orderOptions[*o]
}

func (o *Order) Set(s string) error {
	for i, name := range orderOptions {
		if strings.EqualFold(s, name) {
			*o = Order(i)
			return nil
		}
	}
	return fmt.Errorf("must be one of %s", strings.Join(orderOptions, ", "))
}
