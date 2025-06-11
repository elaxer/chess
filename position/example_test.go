package position_test

import (
	"fmt"

	"github.com/elaxer/chess/position"
)

func ExampleNew() {
	pos := position.New(position.FileA, position.Rank1)
	fmt.Println(pos)

	pos = position.New(position.FileH, position.Rank8)
	fmt.Println(pos)

	pos = position.New(position.FileD, position.Rank12)
	fmt.Println(pos)

	partialPos := position.New(position.FileNull, position.Rank1)
	fmt.Println(partialPos)

	partialPos = position.New(position.FileA, position.RankNull)
	fmt.Println(partialPos)

	emptyPos := position.New(position.FileNull, position.RankNull)
	fmt.Println(emptyPos)

	invalidPos := position.New(position.File(-3), position.Rank4)
	fmt.Println(invalidPos)

	invalidPos = position.New(position.FileM, position.RankMax+1)
	fmt.Println(invalidPos)

	// Output:
	// a1
	// h8
	// d12
	// 1
	// a
	//
	// 4
	// m
}

func ExampleFromString() {
	pos := position.FromString("b5")
	fmt.Printf("File: %s, Rank: %d\n", pos.File, pos.Rank)

	pos = position.FromString("n14")
	fmt.Printf("File: %s, Rank: %d\n", pos.File, pos.Rank)

	partialPos := position.FromString("f")
	fmt.Printf("File: %s, Rank: %d\n", partialPos.File, partialPos.Rank)

	partialPos = position.FromString("9")
	fmt.Printf("File: %s, Rank: %d\n", partialPos.File, partialPos.Rank)

	emptyPos := position.FromString("")
	fmt.Printf("File: %s, Rank: %d\n", emptyPos.File, emptyPos.Rank)

	// Output:
	// File: b, Rank: 5
	// File: n, Rank: 14
	// File: f, Rank: 0
	// File: , Rank: 9
	// File: , Rank: 0
}
