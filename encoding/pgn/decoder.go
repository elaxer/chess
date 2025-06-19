package pgn

import (
	"regexp"

	"github.com/elaxer/chess"
	"github.com/elaxer/rgx"
)

// Decoder decodes a PGN string into headers and moves.
// It uses regular expressions to extract headers and moves from the PGN string.
// The headersRegexp should match the PGN header format, and movesRegexp should match the move format.
type Decoder struct {
	headersRegexp, movesRegexp *regexp.Regexp
}

func NewDecoder(headersRegexp, movesRegexp *regexp.Regexp) *Decoder {
	return &Decoder{headersRegexp, movesRegexp}
}

// Decode decodes a PGN string into headers and moves.
// It returns a slice of Header structs and a slice of chess.Move structs.
// If there is an error during decoding, it returns an error.
// The PGN string should match the regular expressions defined in headersRegexp and movesRegexp.
func (d *Decoder) Decode(pgn string) ([]Header, []chess.Move, error) {
	headers, _ := d.decodeHeaders(pgn)
	moves, _ := d.decodeMoves(pgn)

	return headers, moves, nil
}

func (d *Decoder) decodeHeaders(pgn string) ([]Header, error) {
	headers := make([]Header, 0)

	data, err := rgx.Groups(d.headersRegexp, pgn)
	if err != nil {
		return nil, err
	}

	for _, match := range data {
		headers = append(headers, NewHeader(match["name"], match["value"]))
	}

	return headers, nil
}

func (d *Decoder) decodeMoves(pgn string) ([]chess.Move, error) {
	moves := make([]chess.Move, 0)
	data := d.movesRegexp.FindAllString(pgn, -1)

	for _, move := range data {
		moves = append(moves, chess.StringMove(move))
	}

	return moves, nil
}
