import { File, fileToString } from "./file";

export default class Position {
    public readonly file: File;
    public readonly rank: number;

    public constructor(file: File, rank: number) {
        this.file = file
        this.rank = rank
    }

    public toString() {
        return fileToString(this.file) + this.rank.toString();
    }
}
