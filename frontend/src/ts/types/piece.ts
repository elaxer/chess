export enum Piece {
    Pawn = '',
    Rook = 'R',
    Knight = 'N',
    Bishop = 'B',
    Queen = 'Q',
    King = 'K',
}

export function pieceFromChar(char: string): Piece {
    switch (char) {
        case '':
            return Piece.Pawn
        case 'R':
            return Piece.Rook
        case 'N':
            return Piece.Knight
        case 'B':
            return Piece.Bishop
        case 'Q':
            return Piece.Queen
        case 'K':
            return Piece.King
        default:
            throw new Error(`Invalid piece character: ${char}`)
    }
}
