export enum File {
    A = 1,
    B = 2,
    C = 3,
    D = 4,
    E = 5,
    F = 6,
    G = 7,
    H = 8,
}

export function fileToString(file: File): string {
    return 'abcdefgh'[file - 1]
}

export function fileFromChar(char: string): File {
    return 'abcdefgh'.indexOf(char) as File + 1
}
