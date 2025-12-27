A small, well-tested Go library that implements core chess primitives: board and square management, pieces and moves, position handling, simple metrics and a textual visualizer.

# Requirements

- Go 1.23 or newer

Install with:

```bash
go get github.com/elaxer/chess
```

# About

## What this library is NOT

This library is not:
- A ready-to-play chess game with UI
- A full-featured chess engine with AI, evaluation, or search algorithms
- A beginner-friendly “learn chess by coding” toolkit

There is no GUI, no bots, no minimax, no magic. This package focuses strictly on core chess primitives and rules infrastructure.

## Who this library is for

This library is designed for developers who:
- Want to build their own chess engine or experiment with custom rules
- Need a clean, testable, low-level chess model (board, moves, states)
- Care about explicit state management and invariants
- Are comfortable working with interfaces and composing their own implementations

## Engine implementations

This package is intentionally **engine-agnostic**. It defines rules primitives, board mechanics, move validation, states, and metrics, but it does not enforce any specific chess rule set.

A reference implementation of **standard chess rules** built on top of this library is provided here:

**github.com/elaxer/standardchess**

That package contains:
- standard chess pieces and their movement logic
- turn handling, castling, en passant, promotion, and check/checkmate rules

Think of this repository as the *core*, and `standardchess` as a *canonical engine* that demonstrates how the core is meant to be composed.

If you want to implement:
- chess variants
- custom board sizes
- alternative rule sets

this library stays the same, while the engine implementation changes.

# Documentation
## Board information

You can get the information the board stores. See examples below:

Get the current turn:
```go
var board chess.Board
var turn chess.Side = board.Turn()
```
Get the current board state for the current turn:
```go
var state chess.State = board.State(board.Turn())

// ... or get the board state for the opposite side:
var oppositeState chess.State = board.State(!board.Turn())
```
Get a list of executed moves on the board:
```go
var moveHistory []chess.MoveResult = board.MoveHistory()
```

## Moves

You can easily get available or potential moves on the board:
```go
var availableMoves chess.PositionSet = board.Moves(board.Turn())
var potentialMoves chess.PositionSet = board.Moves(!board.Turn())
```
> `board.Moves` method returns the set of positions to which the pieces of a given side can move. Each piece has method `PseudoMoves`, so the `board.Moves` method returns just a set of filtered moves of the pieces.

You can also get a filtered set of moves for a specific piece:
```go
var piece chess.Piece = board.Squares().FindByPosition(chess.PositionFromString("e2"))
var pieceLegalMoves chess.PositionSet = board.LegalMoves(piece)
```

Here the question arises, what do **legal** and **pseudo** moves mean?

Legal moves are moves that can be made without breaking the rules of the board. Pseudo moves are moves that include moves that would be considered illegal. For example, according to the rules of standard chess, illegal moves are those after which the opponent gets the opportunity to capture the main piece - the king. `board.Moves` and `board.LegalMoves` both return legal moves.

## Moves making

You can make moves:
```go
var move chess.Move = chess.StringMove("Bc5")
var moveResult chess.MoveResult

moveResult, err := board.MakeMove(move)
```
The result of the method execution is an interface `chess.MoveResult`. This value stores the given input move `MoveResult.Move`, also the value has methods which returns a captured piece (if it has, otherwise `nil`) `MoveResult.CapturedPiece`, the side of the input move executor `MoveResult.Side` and the new board state after move execution `MoveResult.BoardNewState`. Also the method returns an `error` if it exists, for example if the input move is incorrect or impossible etc.

You can also undo the last move:
```go
var lastMoveResult chess.MoveResult

lastMoveResult, err := board.UndoMove()
```

## Working with squares

Work with chess squares and the arrangement of pieces on the board are encapsulated in a separate structure `chess.Square`. Your board contains it:
```go
var squares *chess.Square = board.Squares()
```
... or you can create your own:
```go
edgePosition := chess.NewPosition(chess.FileH, chess.Rank8)
squares = chess.NewSquares(edgePosition)
```

> "Edge position" means the most extreme square on the board, also it means the size of the board. In our example we created a 8x8 field

**NOTE**: It is impossible to create squares larger than the value specified in `chess.MaxSupportedPosition`

You can also create squares with placed pieces:
```go
// Here should be your implementation of the piece
var piece chess.Piece
squares, err := SquaresFromPlacement(edgePosition, map[chess.Position]Piece{
    chess.PositionFromString("g3"): piece,
})
```

... or you can place a piece on the existed field:
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
... or you can find them using different ways:
```go
var pieceNotation = "K"
var pieceSide = chess.SideWhite

// Also it finds the position:
piece, position := squares.FindPiece(pieceNotation, pieceSide)
if piece != nil {
    // The piece is found
}

// ... or you can find several pieces with the same notation and the same side:
var pieces chess.Piece[] = squares.GetPieces(pieceNotation, pieceSide)
```

... or you can get all the pieces for the specified side:
```go
pieces = squares.GetAllPieces(pieceSide)
```

And vice versa you can get the position by the piece placed on it:
```go
pos := squares.GetByPiece(piece)
if !pos.IsNull() {
    // Position is found
}

```

### Pieces moving

You can move a piece from one position to another:
```go
capturedPiece, err := squares.MovePiece(chess.PositionFromString("c3"), chess.PositionFromString("h8"))
if capturedPiece != nil {
    // If it is true, then there was a piece on the "h8" square
}
```

... or you can move the piece, call the callback, then return to the original position of the board:
```go
err := squares.MovePiece(chess.PositionFromString("c3"), chess.PositionFromString("h8"), func () {
    // Do things within this new temporary position
})
```

### Iteration over squares

You have different ways to iterate over squares. All this methods are built on the go standard package `iter`. Here is one of them, which goes through all the squares starting from the very first one
and ending with the **edge square**:
```go
for pos, piece := range square.Iter() {
    if piece != nil {
        // There is a piece on the position
    }
}
```

... or you can iterate over rows :
```go
// Switch the iteration direction
backwards = false
for rank, row := range square.Iter(backwards) {
    for file, piece := range row {
        if piece != nil {
            // There is an existed piece
        }

        pos := chess.NewPosition(file, rank)
    }
}
```

... or iterate over squares in a given direction:
```go
// Go through squares on the same file and the next rank
dir1 := chess.NewPosition(chess.FileNull, chess.Rank1)
// Go through squares on the previous file and the same rank
dir2 := chess.NewPosition(-chess.FileA, chess.RankNull)
// Go through diagonally down and to the right of the squares
dir3 := chess.NewPosition(chess.FileA, chess.Rank1)
// Go through diagonally up and to the left of the squares
dir4 := chess.NewPosition(-chess.FileA, -chess.Rank1)
// and so on...

fromPos := chess.PositionFromString("d4")
for pos, piece := range square.IterByDirection(fromPos, dir1) {
    if piece != nil {
        // There is a piece on the position
    }
}
```

## Positions, files, ranks

There is the `chess.Position` structure for working with positions. It contains two fields: `File` and `Rank`.
The fields correspond to types `chess.File` and `chess.Rank` respectively.
A file can be in the range between the values ​​of `chess.FileNull` and `chess.FileMax`.
A rank can be in the range between the values ​​of `chess.RankNull` and `chess.RankMax`.
The engine doesn't expect you to use values out this ranges.

### Files

There is a special type for the files representation:
```go
type File int8
```

Also there are several built-in file constants which you should to use: `chess.FileNull`, `chess.FileMin` which equals to `chess.FileA`, `chess.FileB` ..., `chess.FileP` which equals to `chess.FileMax`.
`chess.FileNull` means zero value and doesn't mean it is invalid.

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

### Ranks

There is a special type for the ranks representation:
```go
type Rank int8
```

Also there are several built-in rank constants which you should to use: `chess.RankNull`, `chess.RankMin` which equals to `chess.Rank1`, `chess.Rank2` ..., `chess.Rank16` which equals to `chess.RankMax`.
`chess.RankNull` means zero value and doesn't mean it is invalid.

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

### Positions

Create position in any way convenient for you:
```go
// Create empty position: 
pos := chess.Position{}
pos = chess.NewEmptyPosition()
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

Position may has several states:
```go
chess.PositionFromString("j3").IsFull() == true
chess.NewEmptyPosition().IsEmpty() == true
chess.PositionFromString("g").IsValid() == true // It is not empty and full but still valid
chess.PositionFromString("z22").IsValid() == false // Invalid
```

## Pieces

You can get the piece side:
```go
var piece chess.Piece
side := piece.Side()
```

... or get it notation:
```go
var notation string = piece.Notation()
```

... or it weight which evaluates the piece's value on the board:
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

var pseudoMoves chess.PositionSet = piece.PseudoMoves(piecePosition, squares)
```

## States, state types

There are 3 board state types: `chess.StateTypeClear`, `chess.StateTypeThreat`, `chess.StateTypeTerminal`.

`chess.StateTypeClear` indicates that the chess board is in a clear state,
meaning there are no threats or special conditions affecting the game. This is the default state of the board.

`chess.StateTypeThreat` indicates that there is a threat on the chess board, which is useful for indicating check or other conditions where a piece is under threat. 

`chess.StateTypeTerminal` indicates that the game has reached a terminal state, such as checkmate or stalemate,
where no further moves can be made. This state is used to signify the end of the game.

There are **states** which built on **state types**. There is the `chess.State` interface which has string representation and contains a state type.
The difference between the **state** and the **state type** is **state type** is more abstract while **state** represents a concrete case of the board state.

You can create your own states for various cases:
```go
var (
    StateCheck = chess.NewState("check", chess.StateTypeThreat)
    StateCheckmate = chess.NewState("checkmate", chess.StateTypeTerminal)
    // etc.
)
```

Also there is the only one built in engine state `chess.StateClear` with state type `chess.StateTypeClear`.

Note that boards contain states:
```go
var board chess.Board
state := board.State(board.Turn())
```

... which can be changed during the process of working with the board

## Metrics

**Metrics** show meta information about the board. There are **metric functions** which returns **metrics**.

Use this built in metric functions for getting the metrics:
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
> Note that a **metric func** should implement the `metric.MetricFunc` type

## Visualizer

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
vis.Visualize(board, os.Stdout)
```

# License

The GNU General Public License