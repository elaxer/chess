import Position from "./position";

export default class Square {
    public readonly piece: string | null;
    public readonly position: Position;

    constructor(piece: string | null, position: Position) {
        this.piece = piece;
        this.position = position;
    }
}