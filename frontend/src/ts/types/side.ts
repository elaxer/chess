export enum Side {
    Black = 0,
    White = 1,
}

export function sideToString(side: Side): string {
    return side == Side.White ? 'w' : 'b'
}
