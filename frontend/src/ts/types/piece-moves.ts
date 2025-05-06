import Position from "./position";
import { Piece } from "./piece";
import { Side } from "./side";

export default class PieceMoves {
    public piece: Piece
    public side: Side
    public from: Position
    public to: Position[]

    constructor(piece: Piece, side: Side, from: Position, to: Position[]) {
        this.piece = piece.toLowerCase() as Piece
        this.side = side
        this.from = from
        this.to = to
    }
}
