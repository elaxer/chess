import PieceMoves from "../types/piece-moves";
import Position from "../types/position";
import { File } from "../types/file";
import { Side } from "../types/side";
import { Piece } from "../types/piece";

export default class MoveFactory {
    public static createFromObject(obj: any): PieceMoves {
        return new PieceMoves(
            obj['piece'] as Piece,
            obj['side'] as Side,
            new Position(obj['from']['file'] as File, obj['from']['rank']),
            obj['to'].map((square: any) => new Position(square['file'] as File, square['rank']))
        )
    }
}