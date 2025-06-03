package position

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/elaxer/chess/pkg/rgx"
	validation "github.com/go-ozzo/ozzo-validation"

	mapset "github.com/deckarep/golang-set/v2"
)

var Regexp = regexp.MustCompile("^(?P<file>[a-p])?(?P<rank>1[0-6]|[1-9])?$")

type Set = mapset.Set[Position]

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

func NewEmpty() Position {
	return Position{}
}

// FromString создает новую позицию из шахматной нотации.
// Например, "e4" будет преобразовано в Position{FileE, Rank4}.
func FromString(str string) Position {
	data, err := rgx.Group(Regexp, str)
	if err != nil {
		return NewEmpty()
	}

	rank, _ := strconv.Atoi(data["rank"])

	return Position{NewFile(data["file"]), Rank(rank)}
}

func (p Position) IsBoundaries(position Position) bool {
	return p.File <= position.File && p.File >= FileMin && p.Rank <= position.Rank && p.Rank >= RankMin
}

func (p Position) IsFull() bool {
	return !p.File.IsNull() && !p.Rank.IsNull()
}

func (p Position) IsPartial() bool {
	return p.File.IsNull() || p.Rank.IsNull()
}

func (p Position) IsEmpty() bool {
	return p.File.IsNull() && p.Rank.IsNull()
}

func (p Position) Validate() error {
	return validation.ValidateStruct(&p, validation.Field(&p.File), validation.Field(&p.Rank))
}

func (p Position) String() string {
	return fmt.Sprintf("%s%s", p.File, p.Rank)
}

func (p *Position) UnmarshalJSON(data []byte) error {
	position := make(map[string]any, 2)
	if err := json.Unmarshal(data, &position); err != nil {
		return err
	}

	if file, ok := position["file"].(float64); ok {
		if File(file).Validate() == nil {
			p.File = File(file)
		}
	}

	if rank, ok := position["rank"].(float64); ok {
		if Rank(rank).Validate() == nil {
			p.Rank = Rank(rank)
		}
	}

	return nil
}
