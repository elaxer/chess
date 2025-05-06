import Position from "./position";
import MoveType from "./move-type";

export class MoveNormal {
    public from: Position;
    public to: Position;
    public moveType: MoveType = MoveType.Normal;

    constructor(from: Position, to: Position) {
        this.from = from;
        this.to = to;
    }
}

export class MoveCastling {
    public moveType: MoveType;

    constructor(moveType: MoveType) {
        this.moveType = moveType;
    }
}