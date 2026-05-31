
### **2.2. Код Тетриса на 4 языках**

#### **Python (`python/tetris.py`) — консольная версия без внешних библиотек**

```python
import random
import os
import time

# Размеры поля
WIDTH = 10
HEIGHT = 20

# Фигурки тетрамино
SHAPES = [
    [[1, 1, 1, 1]],  # I
    [[1, 1], [1, 1]],  # O
    [[0, 1, 0], [1, 1, 1]],  # T
    [[1, 0, 0], [1, 1, 1]],  # L
    [[0, 0, 1], [1, 1, 1]],  # J
    [[0, 1, 1], [1, 1, 0]],  # S
    [[1, 1, 0], [0, 1, 1]]   # Z
]

class Tetris:
    def __init__(self):
        self.field = [[0] * WIDTH for _ in range(HEIGHT)]
        self.score = 0
        self.game_over = False
        self.new_piece()

    def new_piece(self):
        self.piece = random.choice(SHAPES)
        self.piece_x = WIDTH // 2 - len(self.piece[0]) // 2
        self.piece_y = 0
        if self.collision():
            self.game_over = True

    def collision(self):
        for y, row in enumerate(self.piece):
            for x, cell in enumerate(row):
                if cell:
                    field_x = self.piece_x + x
                    field_y = self.piece_y + y
                    if (field_x < 0 or field_x >= WIDTH or
                        field_y >= HEIGHT or
                        (field_y >= 0 and self.field[field_y][field_x])):
                        return True
        return False

    def merge(self):
        for y, row in enumerate(self.piece):
            for x, cell in enumerate(row):
                if cell:
                    self.field[self.piece_y + y][self.piece_x + x] = 1
        self.clear_lines()
        self.new_piece()

    def clear_lines(self):
        lines_cleared = 0
        y = HEIGHT - 1
        while y >= 0:
            if all(self.field[y]):
                del self.field[y]
                self.field.insert(0, [0] * WIDTH)
                lines_cleared += 1
            else:
                y -= 1
        self.score += lines_cleared * 100

    def move(self, dx, dy):
        self.piece_x += dx
        self.piece_y += dy
        if self.collision():
            self.piece_x -= dx
            self.piece_y -= dy
            if dy == 1:
                self.merge()
            return False
        return True

    def rotate(self):
        rotated = list(zip(*self.piece[::-1]))
        old_piece = self.piece
        self.piece = [list(row) for row in rotated]
        if self.collision():
            self.piece = old_piece

    def draw(self):
        os.system('clear')
        # Верхняя граница
        print('+' + '-' * (WIDTH * 2) + '+')
        for y in range(HEIGHT):
            line = '|'
            for x in range(WIDTH):
                if (self.piece_y <= y < self.piece_y + len(self.piece) and
                    self.piece_x <= x < self.piece_x + len(self.piece[0]) and
                    self.piece[y - self.piece_y][x - self.piece_x]):
                    line += '[]'
                elif self.field[y][x]:
                    line += '[]'
                else:
                    line += '  '
            line += '|'
            print(line)
        print('+' + '-' * (WIDTH * 2) + '+')
        print(f'Score: {self.score}')
        if self.game_over:
            print('GAME OVER')

def main():
    game = Tetris()
    while not game.game_over:
        game.draw()
        cmd = input('> ')
        if cmd == 'a':
            game.move(-1, 0)
        elif cmd == 'd':
            game.move(1, 0)
        elif cmd == 's':
            game.move(0, 1)
        elif cmd == 'w':
            game.rotate()
        elif cmd == 'q':
            break
        else:
            game.move(0, 1)

if __name__ == '__main__':
    main()
