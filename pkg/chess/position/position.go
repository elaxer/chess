package position

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Position представляет позицию на шахматной доске.
// Он состоит из вертикали (File) и горизонтали (Rank).
// Например, позиция "e4" соответствует FileE и Rank4.
type Position struct {
	File File `json:"file"`
	Rank Rank `json:"rank"`
}

func New(file File, rank Rank) Position {
	return Position{file, rank}
}

// FromNotation создает новую позицию из шахматной нотации.
// Например, "e4" будет преобразовано в Position{FileE, Rank4}.
// todo:
// Нотация должна содержать ровно 2 символа: первый - буква от 'a' до 'h', второй - цифра от '1' до '8'.
func FromNotation(notation string) Position {
	if len(notation) == 0 {
		return Position{}
	}

	var rank int
	if len(notation) > 1 {
		rank, _ = strconv.Atoi(notation[1:2])
	}

	return Position{NewFile(notation[0:1]), Rank(rank)}
}

func (p Position) Validate() error {
	return validation.ValidateStruct(&p, validation.Field(&p.File), validation.Field(&p.Rank))
}

func (p Position) String() string {
	return fmt.Sprintf("%s%s", p.File, p.Rank)
}

func (p *Position) UnmarshalJSON(data []byte) error {
	position := make(map[string]any, 2)
	err := json.Unmarshal(data, &position)
	if err != nil {
		return err
	}

	file, ok := position["file"].(float64)
	if !ok {
		return errors.New("invalid file coordinates")
	}

	rank, ok := position["rank"].(float64)
	if !ok {
		return errors.New("invalid rank coordinates")
	}

	p.File = File(file)
	p.Rank = Rank(rank)

	return nil
}
