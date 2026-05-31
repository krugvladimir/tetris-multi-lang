#include <ncurses.h>
#include <cstdlib>
#include <ctime>
#include <vector>

const int WIDTH = 10;
const int HEIGHT = 20;

std::vector<std::vector<int>> field(HEIGHT, std::vector<int>(WIDTH, 0));
int score = 0;
bool gameOver = false;

std::vector<std::vector<std::vector<int>>> shapes = {
    {{1,1,1,1}},
    {{1,1},{1,1}},
    {{0,1,0},{1,1,1}},
    {{1,0,0},{1,1,1}},
    {{0,0,1},{1,1,1}},
    {{0,1,1},{1,1,0}},
    {{1,1,0},{0,1,1}}
};

std::vector<std::vector<int>> activePiece;
int pieceX, pieceY;

std::vector<std::vector<int>> getRandomPiece() {
    int idx = rand() % shapes.size();
    return shapes[idx];
}

bool collision() {
    for (size_t y = 0; y < activePiece.size(); y++) {
        for (size_t x = 0; x < activePiece[y].size(); x++) {
            if (activePiece[y][x]) {
                int fx = pieceX + x;
                int fy = pieceY + y;
                if (fx < 0 || fx >= WIDTH || fy >= HEIGHT || (fy >= 0 && field[fy][fx]))
                    return true;
            }
        }
    }
    return false;
}

void merge() {
    for (size_t y = 0; y < activePiece.size(); y++) {
        for (size_t x = 0; x < activePiece[y].size(); x++) {
            if (activePiece[y][x])
                field[pieceY + y][pieceX + x] = 1;
        }
    }
    // очистка линий
    for (int y = HEIGHT-1; y >= 0; ) {
        bool full = true;
        for (int x = 0; x < WIDTH; x++)
            if (!field[y][x]) { full = false; break; }
        if (full) {
            for (int i = y; i > 0; i--)
                field[i] = field[i-1];
            field[0] = std::vector<int>(WIDTH, 0);
            score += 100;
        } else y--;
    }
    // новая фигура
    activePiece = getRandomPiece();
    pieceX = WIDTH/2 - activePiece[0].size()/2;
    pieceY = 0;
    if (collision()) gameOver = true;
}

void move(int dx, int dy) {
    pieceX += dx;
    pieceY += dy;
    if (collision()) {
        pieceX -= dx;
        pieceY -= dy;
        if (dy == 1) merge();
    }
}

void rotate() {
    std::vector<std::vector<int>> rotated(activePiece[0].size(), std::vector<int>(activePiece.size(), 0));
    for (size_t i = 0; i < activePiece.size(); i++)
        for (size_t j = 0; j < activePiece[i].size(); j++)
            rotated[j][activePiece.size()-1-i] = activePiece[i][j];
    auto old = activePiece;
    activePiece = rotated;
    if (collision()) activePiece = old;
}

void draw() {
    clear();
    printw("Score: %d\n", score);
    for (int y = 0; y < HEIGHT; y++) {
        for (int x = 0; x < WIDTH; x++) {
            bool cell = field[y][x];
            if (!cell && pieceY <= y && y < pieceY + (int)activePiece.size() &&
                pieceX <= x && x < pieceX + (int)activePiece[0].size() &&
                activePiece[y-pieceY][x-pieceX])
                cell = true;
            printw(cell ? "[]" : "  ");
        }
        printw("\n");
    }
    refresh();
}

int main() {
    srand(time(0));
    initscr();
    cbreak();
    noecho();
    keypad(stdscr, TRUE);
    nodelay(stdscr, TRUE);
    curs_set(0);

    activePiece = getRandomPiece();
    pieceX = WIDTH/2 - activePiece[0].size()/2;
    pieceY = 0;

    while (!gameOver) {
        draw();
        int ch = getch();
        switch(ch) {
            case KEY_LEFT:  move(-1,0); break;
            case KEY_RIGHT: move(1,0); break;
            case KEY_DOWN:  move(0,1); break;
            case KEY_UP:    rotate(); break;
            case 'q': gameOver = true; break;
        }
        move(0,1);
        napms(300);
    }
    clear();
    printw("GAME OVER! Score: %d\n", score);
    refresh();
    napms(3000);
    endwin();
    return 0;
}
