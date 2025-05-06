import { File, fileFromChar } from "../types/file";
import Position from "../types/position";

export default class PositionFactory {
    public static createFromString(square: string): Position {
        return new Position(fileFromChar(square[0]), +square[1])
    }
}