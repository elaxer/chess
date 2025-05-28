package position

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/elaxer/chess/pkg/rgx"
	validation "github.com/go-ozzo/ozzo-validation"

	mapset "github.com/deckarep/golang-set/v2"
)

var (
	RegexpFrom = fmt.Sprintf("(?P<from>%s?%s?)", RegexpFile, RegexpRank)
	RegexpTo   = fmt.Sprintf("(?P<to>%s%s)", RegexpFile, RegexpRank)
)

var notationRegexp = regexp.MustCompile("^(?P<file>[a-p])?(?P<rank>1[0-6]|[1-9])?$")

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

func NewNull() Position {
	return Position{}
}

// FromNotation создает новую позицию из шахматной нотации.
// Например, "e4" будет преобразовано в Position{FileE, Rank4}.
func FromNotation(notation string) Position {
	result, err := rgx.Group(notationRegexp, notation)
	if err != nil {
		return NewNull()
	}

	rank, _ := strconv.Atoi(result["rank"])

	return Position{NewFile(result["file"]), Rank(rank)}
}

func (p Position) IsInRange(position Position) bool {
	return p.File <= position.File && p.File >= FileA && p.Rank <= position.Rank && p.Rank >= Rank1
}

func (p Position) IsNull() bool {
	return p.File == 0 && p.Rank == 0
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
