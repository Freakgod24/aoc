from enum import StrEnum
from collections import namedtuple
import hashlib

class Direction(StrEnum):
    Up = 'U'
    Down = 'D'
    Left = 'L'
    Right = 'R'

Position = namedtuple('Position', ['x', 'y'])

get_md5_hash = lambda data: hashlib.md5(data.encode()).hexdigest()[:4]
get_new_passcode = lambda passcode, dir: passcode + dir

def get_new_position(pos, dir):
    x, y = pos
    match dir:
        case Direction.Up:
            y -= 1
        case Direction.Down:
            y += 1
        case Direction.Left:
            x -= 1
        case Direction.Right:
            x += 1
    return Position(x,y)

def is_door_open(pos: Position, hash: str):
    if len(hash) < 4:
        return ValueError('Invalid hash, expecting hash length of 4 or more !')
    up, down, left, right = hash[:4]
    return (
        (Direction.Up, up in 'bcdef' and pos.y > 0),
        (Direction.Down, down in 'bcdef' and pos.y < 3),
        (Direction.Left, left in 'bcdef' and pos.x > 0),
        (Direction.Right, right in 'bcdef' and pos.x < 3),
    )

def solve_maze(pos, passcode):
    global min_path_length, max_path_length

    if pos == (3,3):
        min_path_length = min(len(passcode), min_path_length)
        if len(passcode) > max_path_length: 
            max_path_length = len(passcode)
            print(pos, len(passcode), passcode)
        return passcode

    # if len(passcode) > min_path_length:
    #     return

    hash = get_md5_hash(passcode)
    for dir, opened in is_door_open(pos, hash):
        if opened:
            new_passcode = get_new_passcode(passcode, dir)
            new_position = get_new_position(pos, dir)
            solve_maze(new_position, new_passcode)

input_passcode = "gdjjyniy" #"kglvqrro" #"ihgpwlah" #"hijkl"
pos = Position(0, 0)
min_path_length = float('Inf')
max_path_length = 0

solve_maze(pos, input_passcode)
