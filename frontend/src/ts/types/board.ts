import PieceMoves from "./piece-moves";
import MoveType from "./move-type";
import { Side } from "./side";

export default class Board {
    public turn: Side
    public moves: PieceMoves[]
    public castlings: MoveType[]
    public isCheck: boolean
    public isMate: boolean
    public isStalemate: boolean

    constructor(turn: Side, moves: PieceMoves[], castlings: MoveType[], isCheck: boolean, isMate: boolean, isStalemate: boolean) {
        this.turn = turn
        this.moves = moves
        this.castlings = castlings
        this.isCheck = isCheck
        this.isMate = isMate
        this.isStalemate = isStalemate
    }

    isGameOver(): boolean {
        return this.isMate || this.isStalemate
    }
}
