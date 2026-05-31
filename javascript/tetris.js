const canvas = document.getElementById('tetrisCanvas');
const ctx = canvas.getContext('2d');
const nextCanvas = document.getElementById('nextCanvas');
const nextCtx = nextCanvas.getContext('2d');

const COLS = 10;
const ROWS = 20;
const CELL_SIZE = 30;

let field = Array(ROWS).fill().map(() => Array(COLS).fill(0));
let score = 0;
let activePiece = null;
let pieceX, pieceY;
let gameLoop = null;

const SHAPES = [
    [[1,1,1,1]],
    [[1,1],[1,1]],
    [[0,1,0],[1,1,1]],
    [[1,0,0],[1,1,1]],
    [[0,0,1],[1,1,1]],
    [[0,1,1],[1,1,0]],
    [[1,1,0],[0,1,1]]
];

function randomPiece() {
    const shape = SHAPES[Math.floor(Math.random() * SHAPES.length)];
    return shape.map(row => [...row]);
}

function spawnNewPiece() {
    activePiece = randomPiece();
    pieceX = Math.floor((COLS - activePiece[0].length) / 2);
    pieceY = 0;
    if (collision()) {
        gameOver();
    }
}

function collision() {
    for (let y = 0; y < activePiece.length; y++) {
        for (let x = 0; x < activePiece[y].length; x++) {
            if (activePiece[y][x] !== 0) {
                const fieldX = pieceX + x;
                const fieldY = pieceY + y;
                if (fieldX < 0 || fieldX >= COLS || fieldY >= ROWS || (fieldY >= 0 && field[fieldY][fieldX] !== 0)) {
                    return true;
                }
            }
        }
    }
    return false;
}

function merge() {
    for (let y = 0; y < activePiece.length; y++) {
        for (let x = 0; x < activePiece[y].length; x++) {
            if (activePiece[y][x] !== 0) {
                field[pieceY + y][pieceX + x] = 1;
            }
        }
    }
    clearLines();
    spawnNewPiece();
}

function clearLines() {
    let linesCleared = 0;
    for (let y = ROWS - 1; y >= 0; y--) {
        if (field[y].every(cell => cell !== 0)) {
            field.splice(y, 1);
            field.unshift(Array(COLS).fill(0));
            linesCleared++;
            y++;
        }
    }
    score += linesCleared * 100;
    document.getElementById('score').innerText = score;
}

function move(dx, dy) {
    pieceX += dx;
    pieceY += dy;
    if (collision()) {
        pieceX -= dx;
        pieceY -= dy;
        if (dy === 1) {
            merge();
        }
        return false;
    }
    return true;
}

function rotate() {
    const rotated = activePiece[0].map((_, idx) => activePiece.map(row => row[idx]).reverse());
    const oldPiece = activePiece;
    activePiece = rotated;
    if (collision()) {
        activePiece = oldPiece;
    }
}

function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // поле
    for (let y = 0; y < ROWS; y++) {
        for (let x = 0; x < COLS; x++) {
            if (field[y][x]) {
                ctx.fillStyle = '#FF4B4B';
                ctx.fillRect(x * CELL_SIZE, y * CELL_SIZE, CELL_SIZE - 1, CELL_SIZE - 1);
            } else {
                ctx.strokeStyle = '#333';
                ctx.strokeRect(x * CELL_SIZE, y * CELL_SIZE, CELL_SIZE - 1, CELL_SIZE - 1);
            }
        }
    }
    // активная фигура
    for (let y = 0; y < activePiece.length; y++) {
        for (let x = 0; x < activePiece[y].length; x++) {
            if (activePiece[y][x]) {
                ctx.fillStyle = '#FF4B4B';
                ctx.fillRect((pieceX + x) * CELL_SIZE, (pieceY + y) * CELL_SIZE, CELL_SIZE - 1, CELL_SIZE - 1);
            }
        }
    }
}

function gameOver() {
    clearInterval(gameLoop);
    alert('Game Over! Score: ' + score);
}

function step() {
    move(0, 1);
    draw();
}

function init() {
    spawnNewPiece();
    draw();
    gameLoop = setInterval(step, 500);
    window.addEventListener('keydown', (e) => {
        switch(e.key) {
            case 'ArrowLeft': move(-1, 0); draw(); break;
            case 'ArrowRight': move(1, 0); draw(); break;
            case 'ArrowDown': move(0, 1); draw(); break;
            case 'ArrowUp': rotate(); draw(); break;
            case ' ': move(0, 1); while(move(0,1)) {}; draw(); break;
        }
    });
}

init();
