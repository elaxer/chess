import { File, fileFromChar } from "../types/file"
import { Piece, pieceFromChar } from "../types/piece"
import PieceMove from "../types/piece-move"
import Position from "../types/position"
import { Side } from "../types/side"

export default class PieceMoveFactory {
    public static createFromNotation(notation: string, side: Side): PieceMove[] {
        try {
            return [PieceMoveFactory.createMoveFromNotation(notation, side)]
        } catch (error) {
            return PieceMoveFactory.createCastlingMoveFromNotation(notation, side)
        }
    }

    private static createMoveFromNotation(notation: string, side: Side): PieceMove {
        const regex = /^(?<piece>[KQBNR])?(?<file_from>[a-h])(?<rank_from>[1-8])(?<is_capture>x)?(?<file>[a-h])(?<rank>[1-8])(?<checkmate>[+#])?$/
        const match = notation.trim().match(regex)

        if (!match || !match.groups) {
            throw new Error(`Invalid move notation: "${notation}"`)
        }

        const pieceChar = !match.groups.piece ? '' : match.groups.piece.toUpperCase()
        const piece = pieceFromChar(pieceChar)

        return new PieceMove(
            piece,
            side,
            new Position(fileFromChar(match.groups.file_from),parseInt(match.groups.rank_from, 10)),
            new Position(fileFromChar(match.groups.file),parseInt(match.groups.rank, 10)),
        )
    }

    private static createCastlingMoveFromNotation(notation: string, side: Side): PieceMove[] {
        const regex = /^0-0(?<long_castling>-0)?(?<checkmate>[+#])?$/
        const match = notation.match(regex)
        
        if (!match || !match.groups) {
            throw new Error(`It is not a castling move: "${notation}"`)
        }

        const rank = side === Side.White ? 1 : 8

        const newKingFile = !match.groups.long_castling ? File.G : File.C
        const newRookFile = !match.groups.long_castling ? File.F : File.D
        const rookFile = !match.groups.long_castling ? File.H : File.A

        return [
            new PieceMove(Piece.King, side, new Position(File.E, rank), new Position(newKingFile, rank)),
            new PieceMove(Piece.Rook, side, new Position(rookFile, rank), new Position(newRookFile, rank)),
        ]
    }
}