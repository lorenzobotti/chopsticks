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
        document.getElementById('one-one'),
        document.getElementById('one-two'),
    ],
    [
        document.getElementById('two-one'),
        document.getElementById('two-two'),
    ],
]

const splitButtons = [
    document.getElementById('split')
]

function positionToString(pos: Position): string {
    return `${pos.turn} -> ${pos[0][0]}-${pos[0][1]} | ${pos[1][0]}-${pos[1][1]}`
}

function parsePosition(s: string): Position {
    const [turn, rest] = s.split(' -> ')
    const [one, two] = rest.split(' | ')

    const [oneOne, oneTwo] = one.split('-')
    const [twoOne, twoTwo] = one.split('-')

    return {
        turn: parseInt(turn),
        players: [
            [parseInt(oneOne), parseInt(oneTwo)],
            [parseInt(twoOne), parseInt(twoTwo)],
        ]
    }
}