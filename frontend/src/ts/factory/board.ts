import Board from "../types/board";
import MoveType from "../types/move-type";
import MoveFactory from "./move";
import { Side } from "../types/side";

export default class BoardFactory {
    public static createFromObject(obj: any): Board {
        return new Board(
            obj['turn'] as Side,
            obj['moves'].map((move: any) => MoveFactory.createFromObject(move)),
            obj['castlings'] as MoveType[],
            obj['is_check'],
            obj['is_mate'],
            obj['is_stalemate'],
        )
    }
}