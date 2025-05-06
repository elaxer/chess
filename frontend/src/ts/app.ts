import * as bootstrap from 'bootstrap'
import "../scss/style.scss";

import Board from "./types/board"
import BoardFactory from "./factory/board"
import MoveType from "./types/move-type"
import { File, fileToString } from "./types/file"
import PieceMoves from "./types/piece-moves"
import { Side, sideToString } from "./types/side"
import Position from "./types/position"
import PositionFactory from "./factory/position"
import WebsocketMessage from "./types/websocket/message"
import WebsocketMessageType from "./types/websocket/message-type"
import { MoveNormal, MoveCastling } from "./types/move-normal";
import Game from './view/game';

let board: Board | null = null

const socket = new WebSocket("http://127.0.0.1:8080/ws")

socket.onopen = () => {
    console.log('[open] Соединение установлено');
    socket.send('hello');
}

socket.onmessage = (event: MessageEvent) => {
    board = BoardFactory.createFromObject(JSON.parse(event.data))
    displayBoard(board)
    
    if (board.isGameOver()) end(board)
}

socket.onclose = (event: CloseEvent) => {
    if (!event.wasClean) return console.log('[close] Соединение прервано')

    console.log(`[close] Соединение закрыто чисто, код=${event.code} причина=${event.reason}`)
}

socket.onerror = (error: Event) => {
    console.log(error);
}

document.addEventListener('DOMContentLoaded', () => {
    document.querySelectorAll('.square').forEach((squareElement: Element) => {
        squareElement.addEventListener('click', (event: Event) => {
            if (!board) return

            let castling: MoveType | null = null;
            if (squareElement.querySelector('#castlingShort')) {
                castling = MoveType.CastlingShort
            } else if (squareElement.querySelector('#castlingLong')) {
                castling = MoveType.CastlingLong
            } else return

            animate(getActiveSquareElement()!, squareElement)
            animate(
                getRookSquareElement(board, castling),
                getNewRookSquareElement(board, castling), 
                () => socket.send(JSON.stringify(new WebsocketMessage(WebsocketMessageType.Move, new MoveCastling(castling)))),
            )

            event.stopImmediatePropagation()
        })

        squareElement.addEventListener('click', (event: Event) => {
            if (!squareElement.querySelector('.square-move-marker')) return

            const activeSquareElement = getActiveSquareElement()!

            animate(activeSquareElement, squareElement, () => {
                socket.send(JSON.stringify(new WebsocketMessage(
                    WebsocketMessageType.Move,
                    new MoveNormal(
                        PositionFactory.createFromString(activeSquareElement.id),
                        PositionFactory.createFromString(squareElement.id),
                    ),
                )))
            })

            event.stopImmediatePropagation()
        })

        squareElement.addEventListener('click', () => {
            if (board === null) return

            clearBoard()

            const pieceElement: HTMLImageElement | null = squareElement.querySelector('.piece')
            if (pieceElement === null) return

            const [sideStr, pieceStr] = pieceElement.src.split('/').at(-1)!.split('.')[0].split('')

            if (sideStr !== sideToString(board.turn)) return

            squareElement.classList.add('square-active')
            
            const moves = board.moves.find((move: PieceMoves) => move.from.toString() === squareElement.id)
            moves!.to.forEach((position: Position) => document.querySelector('#' + position.toString())!.appendChild(newSquareMoveMarker()))

            if (pieceStr !== 'k') return

            const kingSquare = PositionFactory.createFromString(squareElement.id)

            if (board.castlings.includes(MoveType.CastlingShort)) {
                const castlingPosition = new Position((kingSquare.file + 2) as File, kingSquare.rank)
                document.getElementById(castlingPosition.toString())!.appendChild(newSquareMoveMarker('castlingShort'))
            }
            if (board.castlings.includes(MoveType.CastlingLong)) {
                const castlingPosition = new Position((kingSquare.file - 2) as File, kingSquare.rank)
                document.getElementById(castlingPosition.toString())!.appendChild(newSquareMoveMarker('castlingLong'))
            }
        })
    })
})

const game = new Game()
document.addEventListener('DOMContentLoaded', game.displayBoardFromMoves.bind(game))
document.getElementById('first-move')!.addEventListener('click', game.firstMove.bind(game))

document.getElementById('prev-move')!.addEventListener('click', game.prevMove.bind(game))
document.addEventListener('keydown', (e: KeyboardEvent) => e.code === 'ArrowLeft' && game.prevMove())

document.getElementById('next-move')!.addEventListener('click', game.nextMove.bind(game))
document.addEventListener('keydown', (e: KeyboardEvent) => e.code === 'ArrowRight' && game.nextMove())

document.getElementById('last-move')!.addEventListener('click', game.lastMove.bind(game))


function displayBoard(board: Board) {
    document.querySelectorAll('.piece').forEach(piece => piece.remove())
    document.querySelectorAll('.square').forEach((squareElement: Element) => (squareElement as HTMLElement)!.style.cursor = 'default')

    board.moves.forEach((move: PieceMoves) => {
        if (move.side === board.turn) document.getElementById(move.from.toString())!.style.cursor = 'pointer'
    })

    board.moves.forEach((move: PieceMoves) => document.getElementById(move.from.toString())!.appendChild(newPieceImage(move)))
}

function clearBoard() {
    getActiveSquareElement()?.classList.remove('square-active')

    document.querySelectorAll('.square-move-marker').forEach((element: Element) => element.remove())
}

function newPieceImage(move: PieceMoves): HTMLImageElement {
    let image = document.createElement('img')
    image.src = '/public/img/' + sideToString(move.side) + move.piece + '.svg'
    image.classList.add('piece')

    return image;
}

function newSquareMoveMarker(id: string | null = null): HTMLSpanElement {
    let squareMoveMarker = document.createElement('span')
    squareMoveMarker.classList.add('square-move-marker')
    squareMoveMarker.id = id ? id : ''

    return squareMoveMarker
}

function getActiveSquareElement(): Element | null {
    return document.querySelector('.square-active')
}

function animate(fromSquareElement: Element, toSquareElement: Element, callback: CallableFunction | null = null, duration: number = 200) {
    const pieceElement = fromSquareElement!.querySelector('.piece') as HTMLImageElement
    
    const fromRect = fromSquareElement!.getBoundingClientRect()
    const toRect = toSquareElement.getBoundingClientRect()

    const dx = toRect.left - fromRect.left
    const dy = toRect.top - fromRect.top

    clearBoard()

    pieceElement!.style.transform = `translate(${dx}px, ${dy}px) scale(0.85)`;
    pieceElement.style.zIndex = '102';

    setTimeout(() => {
        toSquareElement!.appendChild(pieceElement);
        pieceElement.style.transform = '';

        if (callback) callback()
    }, duration)
}

function getRookSquareElement(board: Board, castling: MoveType): Element {
    let file: File | null = null;
    if (castling === MoveType.CastlingShort) {
        file = File.H
    } else if (castling === MoveType.CastlingLong) {
        file = File.A
    } else throw new Error('Неверный тип castling')

    const rank = board.turn ? 1 : 8

    const rookSquareElement = document.getElementById(fileToString(file) + rank)
    if (!rookSquareElement) {
        throw new Error('Не удалось найти ладью для рокировки')
    }

    return rookSquareElement
}

function getNewRookSquareElement(board: Board, castling: MoveType): Element {
    let file: File | null = null;
    if (castling === MoveType.CastlingShort) {
        file = File.F
    } else if (castling === MoveType.CastlingLong) {
        file = File.D
    } else throw new Error('Неверный тип castling')

    const rank = board.turn ? 1 : 8

    return document.getElementById(fileToString(file) + rank)!
}

function end(board: Board) {
    const gameResultElement = new bootstrap.Modal(document.getElementById('game-result')!)
    let gameResultText = ''

    if (board.isMate) {
        const side = !board.turn ? 'белые' : 'чёрные'
        gameResultText = 'Объявлен мат, ' + side + ' победили!'
    } else if (board.isStalemate) {
        gameResultText = 'Объявлен пат, ничья!'
    }

    document.getElementById('game-result.text')!.textContent = gameResultText
    gameResultElement.show();
}
