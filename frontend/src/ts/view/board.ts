import { Piece } from "../types/piece"
import PieceMove from "../types/piece-move"
import Position from "../types/position"
import { Side, sideToString } from "../types/side"
import { File } from "../types/file"
import PieceMoves from "../types/piece-moves"

export default class Board {
    protected newPieceImage(move: PieceMoves): HTMLImageElement {
        let image = document.createElement('img')
        image.src = '/public/img/' + sideToString(move.side) + move.piece + '.svg'
        image.classList.add('piece')

        return image;
    }

    protected initBoard() {
        document.querySelectorAll('.piece').forEach(piece => piece.remove())
        
        this.getInitPieces().forEach((pieceMove: PieceMove) => {
            document.getElementById(pieceMove.to.toString())!
                .appendChild(this.newPieceImage(new PieceMoves(pieceMove.piece, pieceMove.side, new Position(1, 1), [])))
        })
    }

    private getInitPieces(): PieceMove[] {
        const pieces = [
            new PieceMove(Piece.Rook, Side.White, new Position(1, 1), new Position(File.A, 1)),
            new PieceMove(Piece.Knight, Side.White, new Position(1, 1), new Position(File.B, 1)),
            new PieceMove(Piece.Bishop, Side.White, new Position(1, 1), new Position(File.C, 1)),
            new PieceMove(Piece.Queen, Side.White, new Position(1, 1), new Position(File.D, 1)),
            new PieceMove(Piece.King, Side.White, new Position(1, 1), new Position(File.E, 1)),
            new PieceMove(Piece.Bishop, Side.White, new Position(1, 1), new Position(File.F, 1)),
            new PieceMove(Piece.Knight, Side.White, new Position(1, 1), new Position(File.G, 1)),
            new PieceMove(Piece.Rook, Side.White, new Position(1, 1), new Position(File.H, 1)),
    
            new PieceMove(Piece.Rook, Side.Black, new Position(1, 1), new Position(File.A, 8)),
            new PieceMove(Piece.Knight, Side.Black, new Position(1, 1), new Position(File.B, 8)),
            new PieceMove(Piece.Bishop, Side.Black, new Position(1, 1), new Position(File.C, 8)),
            new PieceMove(Piece.Queen, Side.Black, new Position(1, 1), new Position(File.D, 8)),
            new PieceMove(Piece.King, Side.Black, new Position(1, 1), new Position(File.E, 8)),
            new PieceMove(Piece.Bishop, Side.Black, new Position(1, 1), new Position(File.F, 8)),
            new PieceMove(Piece.Knight, Side.Black, new Position(1, 1), new Position(File.G, 8)),
            new PieceMove(Piece.Rook, Side.Black, new Position(1, 1), new Position(File.H, 8)),
        ]
    
        for (let file = 1; file <= 8; file++) {
            pieces.push(new PieceMove(Piece.Pawn, Side.White, new Position(1, 1), new Position(file as File, 2)))
            pieces.push(new PieceMove(Piece.Pawn, Side.Black, new Position(1, 1), new Position(file as File, 7)))
        }
    
        return pieces
    }
}
