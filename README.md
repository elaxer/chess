# chess - core chess primitives

A small, well-tested Go library that implements core chess primitives: board and square management, pieces and moves, position handling, simple metrics and a textual visualizer.

## Requirements

- Go 1.23 or newer

Install with:

```bash
go get github.com/elaxer/chess
```

## About

### What this library is NOT

This library is not:
- A ready-to-play chess game with UI
- A full-featured chess engine with AI, evaluation, or search algorithms
- A beginner-friendly “learn chess by coding” toolkit

There is no GUI, no bots, no minimax, no magic. This package focuses strictly on core chess primitives and rules infrastructure.

### Who this library is for

This library is designed for developers who:
- want to build their own chess engine or experiment with custom rules
- need a clean, testable, low-level chess model (board, moves, states)
- care about explicit state management and invariants
- are comfortable working with interfaces and composing their own implementations

### Engine implementations

This package is intentionally engine-agnostic. It defines chess primitives such as board representation, positions, directions, move semantics, states, and metrics, without enforcing any concrete rule set or starting position.

Concrete rule implementations are provided as separate engines built on top of this core:
- [github.com/elaxer/standardchess](https://github.com/elaxer/standardchess)
- [github.com/elaxer/fischerchess](https://github.com/elaxer/fischerchess)

Think of this repository as the core mechanics layer, while standardchess and fischerchess are engine-level compositions that define concrete rules and constraints.

If you want to implement any of the following:
- chess variants
- alternative starting positions
- custom rule sets
- non-standard boards

this library remains unchanged, while the engine implementation varies.

## Documentation
### Board information

You can retrieve information stored in the board. See examples below:

Get the current turn:
```go
var board chess.Board
var turn chess.Color = board.Turn()
```
Get the current board state:
```go
var state chess.State = board.State()

```
Get a list of executed moves on the board:
```go
var moveHistory []chess.MoveResult = board.MoveHistory()
```

### Moves

You can easily get available moves on the board:
```go
var availableMoves []chess.Position = board.Moves()
```
> The `board.Moves` method returns the set of positions to which the pieces can move. Each piece has a method `PseudoMoves`, so `board.Moves` returns a filtered set of the pieces' moves.

You can also get a filtered set of moves for a specific piece:
```go
var piece chess.Piece = board.Squares().FindByPosition(chess.PositionFromString("e2"))
var pieceLegalMoves []chess.Position = board.LegalMoves(piece)
```

Here the question arises: what do **legal** and **pseudo** moves mean?

Legal moves are moves that can be made without breaking the board's rules. Pseudo moves include moves that would be considered illegal. For example, in standard chess, an illegal move is one after which the opponent can capture the king. `board.Moves` and `board.LegalMoves` both return legal moves.

### Making moves

You can make moves:
```go
var move chess.Move = chess.StringMove("Bc5")
var moveResult chess.MoveResult

moveResult, err := board.MakeMove(move)
```
The method returns a `chess.MoveResult` and an `error`. `MoveResult` contains the input move (`MoveResult.Move`) and provides methods to get the captured piece (if any) via `MoveResult.CapturedPiece`, the side that made the move (`MoveResult.Side`), and the new board state after the move (`MoveResult.BoardNewState`). The `error` is non‑nil if the move was incorrect or impossible.

You can also undo the last move:
```go
var lastMoveResult chess.MoveResult

lastMoveResult, err := board.UndoMove()
```

### Working with squares

Operations on chess squares and piece arrangement are encapsulated in the `chess.Square` structure. Your board contains it:
```go
var squares *chess.Square = board.Squares()
```
... or you can create your own:
```go
edgePosition := chess.NewPosition(chess.FileH, chess.Rank8)
squares = chess.NewSquares(edgePosition)
```

> "Edge position" refers to the most extreme square on the board and defines the board's size. In our example we created an 8x8 field.

**NOTE**: It is impossible to create squares larger than the value specified by `chess.MaxSupportedPosition`.

You can also create squares with placed pieces:
```go
// Here should be your implementation of the piece
var piece chess.Piece
squares, err := SquaresFromPlacement(edgePosition, map[chess.Position]Piece{
    chess.PositionFromString("g3"): piece,
})
```

... or you can place a piece on an existing field:
```go
err := squares.PlacePiece(piece, chess.PositionFromString("g3"))
```

### Finding pieces, positions

You can find your placed pieces:
```go
piece, err := squares.FindByPosition(chess.PositionFromString("g3"))
if piece != nil {
    // The piece is found
}
```
... or you can find them in different ways:
```go
var pieceNotation = "K"
var pieceColor = chess.ColorWhite

// Also it finds the position:
piece, position := squares.FindPiece(pieceNotation, pieceColor)
if piece != nil {
    // The piece is found
}

// ... or you can find several pieces with the same notation and the same side:
// todo:
var pieces chess.Piece[] = squares.GetPieces(pieceNotation, pieceColor)
```

... or you can get all the pieces of a certain color:
```go
pieces = squares.GetAllPieces(pieceColor)
```

Conversely, you can get the position of a piece placed on it:
```go
pos := squares.GetByPiece(piece)
if !pos.IsNull() {
    // Position is found
}

```

### Moving pieces

You can move a piece from one position to another:
```go
capturedPiece, err := squares.MovePiece(chess.PositionFromString("c3"), chess.PositionFromString("h8"))
if capturedPiece != nil {
    // If true, then there was a piece on the "h8" square.
}
```

... or you can move the piece, call a callback, and then return the board to its original position:
```go
err := squares.MovePiece(chess.PositionFromString("c3"), chess.PositionFromString("h8"), func () {
    // Do things within this new temporary position
})
```

### Iteration over squares

There are different ways to iterate over squares. All these methods are built on the Go standard package `iter`. Here is one of them, which goes through all the squares starting from the very first one and ending with the **edge square**:
```go
for pos, piece := range square.Iter() {
    if piece != nil {
        // There is a piece at that position
    }
}
```

... or you can iterate over rows:
```go
// Switch the iteration direction
backwards = false
for rank, row := range square.Iter(backwards) {
    for file, piece := range row {
        if piece != nil {
            // There is a piece
        }

        pos := chess.NewPosition(file, rank)
    }
}
```

... or iterate over squares in a given direction:
```go
// Traverse squares on the same file and the next rank
dir1 := chess.DirectionTop
// Traverse squares on the previous file and the same rank
dir2 := chess.DirectionBottom
// Traverse diagonally down and to the right
dir3 := chess.DirectionTopRight
// Traverse diagonally up and to the left
dir4 := chess.DirectionBottomLeft
// Use a custom direction
dir5 := position.New(chess.File(2), chess.Rank(1))
// and so on...

fromPos := chess.PositionFromString("d4")
for pos, piece := range square.IterByDirection(fromPos, dir1) {
    if piece != nil {
        // There is a piece on the position
    }
}
```

### Positions, files, ranks

The `chess.Position` structure is used for working with positions. It contains two fields: `File` and `Rank`.
The fields correspond to types `chess.File` and `chess.Rank` respectively.
A file can range between `chess.FileNull` and `chess.FileMax`.
A rank can range between `chess.RankNull` and `chess.RankMax`.
The engine doesn't expect you to use values outside these ranges.

#### Files

There is a special type for file representation:
```go
type File int8
```

There are several built-in file constants that you should use: `chess.FileNull`, `chess.FileMin` which equals `chess.FileA`, `chess.FileB` ..., `chess.FileP` which equals `chess.FileMax`.
`chess.FileNull` is the zero value and does not necessarily mean it is invalid.

You can check if the file has zero value:
```go
var file chess.File
if file.IsNull() {
    // ...
}
```

or validate the file:
```go
if err := file.Validate(); err != nil {
    // ...
}
```

You can also get the text representation of the file:
```go
fileStr := file.String()
```
> An empty string for zero value or invalid files, and letters for other files (`chess.FileA.String() == "A"`, `chess.FileG.String() == "G"` etc.)

#### Ranks

There is a special type for rank representation:
```go
type Rank int8
```

There are several built-in rank constants that you should use: `chess.RankNull`, `chess.RankMin` which equals `chess.Rank1`, `chess.Rank2` ..., `chess.Rank16` which equals `chess.RankMax`.
`chess.RankNull` is the zero value and does not necessarily mean it is invalid.

You can check if the rank has zero value:
```go
var rank chess.Rank
if rank.IsNull() {
    // ...
}
```

or validate the rank:
```go
if err := rank.Validate(); err != nil {
    // ...
}
```

You can also get the text representation of the rank:
```go
rankStr := rank.String()
```
> An empty string for zero value or invalid ranks, and numbers for other ranks (`chess.Rank1.String() == "1"`, `chess.Rank10.String() == "10"` etc.)

#### Positions

Create a position in any way convenient for you:
```go
// Create empty position: 
pos := chess.Position{}
pos = chess.NewPositionEmpty()
pos = chess.NewPosition(chess.FileNull, chess.RankNull)
pos = chess.PositionFromString("")

// ... or half-filled position:
pos = chess.Position{File: chess.FileA}
pos = chess.NewPosition(chess.FileNull, chess.Rank8)
pos = chess.PositionFromString("g")
pos = chess.PositionFromString("7")

// ... or full filled positions:
pos = chess.Position{File: chess.FileA, Rank: chess.Rank8}
pos = chess.NewPosition(chess.FileD, chess.Rank2)
pos = chess.PositionFromString("j3")
```

Get the string representation of the position:
```go
chess.NewPosition(chess.FileNull, chess.RankNull) == ""
chess.NewPosition(chess.FileG, chess.RankNull) == "g"
chess.NewPosition(chess.FileNull, chess.Rank7) == "7"
chess.NewPosition(chess.FileJ, chess.Rank3) == "j3"
```

A position may have several states:
```go
chess.PositionFromString("j3").IsFull() == true
chess.NewPositionEmpty().IsEmpty() == true
chess.PositionFromString("g").IsValid() == true // It is not empty and full but still valid
chess.PositionFromString("z22").IsValid() == false // Invalid
```

### Pieces

You can get a piece's side:
```go
var piece chess.Piece
side := piece.Color()
```

... or its notation:
```go
var notation string = piece.Notation()
```

... or its weight, which evaluates the piece's value on the board:
```go
var pieceWeight uint8 = piece.Weight()
```

... or check if the piece has moved:
```go
if piece.IsMoved() {
    // ...
}
```

... or mark it as moved:
```go
piece.MarkMoved()
piece.IsMove() == true 
```

... or get pseudo moves which the piece generates:
```go
piecePosition := squares.GetByPiece(piece)

var pseudoMoves []chess.Position = piece.PseudoMoves(piecePosition, squares)
```

### States, state types

There are three board state types: `chess.StateTypeClear`, `chess.StateTypeThreat`, `chess.StateTypeTerminal`.

`chess.StateTypeClear` indicates that the chess board is in a clear state, meaning there are no threats or special conditions affecting the game. This is the default state of the board.

`chess.StateTypeThreat` indicates that there is a threat on the chess board, which is useful for indicating check or other conditions where a piece is under threat.

`chess.StateTypeTerminal` indicates that the game has reached a terminal state, such as checkmate or stalemate,
where no further moves can be made. This state is used to signify the end of the game.

There are **states** which are built on **state types**. The `chess.State` interface provides a string representation and includes a state type.
The difference between a **state** and a **state type** is that a state type is more abstract, while a state represents a concrete case of the board state.

You can create your own states for various cases:
```go
var (
    StateCheck = chess.NewState("check", chess.StateTypeThreat)
    StateCheckmate = chess.NewState("checkmate", chess.StateTypeTerminal)
    // etc.
)
```

Also, there is a single built-in engine state, `chess.StateClear`, with the state type `chess.StateTypeClear`.

Note that boards contain states:
```go
var board chess.Board
var state chess.State = board.State()
if !state.Type().IsTerminal() {
    // The game is over
}
```

... which can be changed during the process of working with the board

### Metrics

**Metrics** show meta information about the board. There are **metric functions** which return **metrics**.

Use these built-in metric functions to get metrics:
```go
// The number of half-moves made in the game
metr = metric.HalfmoveCounter(board)
// The number of full moves made in the game
metr = metric.FullmoveCounter(board)
// The last move made, or nil if no moves exist
metr = metric.LastMove(board)
// Material values for White and Black as a slice [white, black]
metr = metric.Material(board)
// The material advantage (white - black)
metr = metric.MaterialDifference(board)

fmt.Printf("Metric \"%s\" shows: %v\n", metr.Name(), metr.Value())
```

... or create your own:
```go
func TurnMetric(board chess.Board) Metric {
    return metric.New("Turn", board.Turn())
}
```
> Note that a metric function should implement the `metric.MetricFunc` type

### Visualizer

Use the `visualizer` package for displaying your board in the ascii format.
It's very useful for debugging your code.
Here is a quick example:
```go

// Thus the white side will be at the bottom and the black side at the top
var orientation visualizer.OptionOrientation = visualizer.OptionOrientationDefault
// The black side will be at the bottom and the white side at the top
orientation = visualizer.OptionOrientationReversed
// The white side will be at the bottom if the current turn is White,
// otherwise the black side will be at the bottom
orientation = visualizer.OptionOrientationByTurn

var vis visualizer.Visualizer{
    Options: visualizer.Options{
        Orientation: orientation,
        // If it is true then the visualizer will show ranks at the left and files at the bottom
        DisplayPositions: true,
        // Metric funcs for displaying the board metrics 
        MetricFuncs: [
            metric.HalfmoveCounterFunc,
            metric.LastMove,
            metric.MaterialDifference,
        ],
    }
}

var board chess.Board

// It will show the board in the terminal
vis.Fprintln(board, os.Stdout)
```

## Contributing

Bug reports and contributions are welcome. Please open issues or pull requests against this repository. Keep changes small and add tests for new behavior.

## License

The GNU General Public License