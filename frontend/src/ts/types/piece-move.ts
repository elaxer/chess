import { Piece } from "./piece"
import Position from "./position"
import { Side } from "./side"

export default class PieceMove {
    public readonly piece: Piece
    public readonly side: Side
    public readonly from: Position
    public readonly to: Position

    constructor(piece: Piece, side: Side, from: Position, to: Position) {
        this.piece = piece
        this.side = side
        this.from = from
        this.to = to
    }
}
