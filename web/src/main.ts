import * as positions from '../assets/moves.json'

interface Position {
    players: number[][],
    turn: number,
}

const startingPosition = {
    players: [[1,1], [1,1]],
    turn: 0,
}

const handsButtons = [
    [
        document.getElementById("one-one"),
        document.getElementById("one-two"),
    ],
    [
        document.getElementById("two-one"),
        document.getElementById("two-two"),
    ],
]

const splitButtons = [
    
]