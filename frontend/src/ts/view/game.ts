import PieceMove from "../types/piece-move"
import Position from "../types/position"
import { Side } from "../types/side"
import PieceMoveFactory from "../factory/piece-moves"
import PieceMoves from "../types/piece-moves"
import Board from "./board"

export default class Game extends Board {
    private moves: PieceMove[][] = []
    private currentMove: number = 0

    constructor() {
        super()

        const movesJsonEl = document.getElementById('moves-json')
        if (!movesJsonEl || !movesJsonEl.textContent) return

        const movesNotation = JSON.parse(movesJsonEl.textContent)
        
        let moves: PieceMove[][] = []
        for (let i = 0; i < movesNotation.length; i++) {
            moves.push(PieceMoveFactory.createFromNotation(movesNotation[i], i % 2 === 0 ? Side.White : Side.Black))
        }

        this.moves = moves
        this.currentMove = moves.length

        this.initBoard()
    }

    public displayBoardFromMoves() {
        const moves = this.moves.slice(0, this.currentMove).flat()

        const moveBadges = document.querySelectorAll('.move')
        for (let i = 0; i < moveBadges.length; i++) {
            if (i === this.currentMove - 1) {
                moveBadges[i].classList.remove('bg-light')
                moveBadges[i].classList.add('bg-primary')

                continue
            }

            moveBadges[i].classList.remove('bg-primary')
            moveBadges[i].classList.add('bg-light')
        }

        for (let i = 0; i < moves.length; i++) {
            const move = moves[i]
            const toSquareEl = document.getElementById(move.to.toString())!
    
            toSquareEl.querySelector('.piece')?.remove()
            toSquareEl.appendChild(this.newPieceImage(new PieceMoves(move.piece, move.side, new Position(1, 1), [])))
    
            document.getElementById(move.from.toString())!.querySelector('.piece')?.remove()
        }
    }

    public firstMove() {
        this.currentMove = 0
        this.initBoard()
        this.displayBoardFromMoves()
    }

    public prevMove() {
        this.currentMove = Math.max(this.currentMove - 1, 0)
        this.initBoard()
        this.displayBoardFromMoves()
    }

    public nextMove() {
        this.currentMove = Math.min(this.currentMove + 1, this.moves.length)
        this.initBoard()
        this.displayBoardFromMoves()
    }

    public lastMove() {
        this.currentMove = this.moves.length
        this.initBoard()
        this.displayBoardFromMoves()
    }
}


