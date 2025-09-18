tiles = 5
safe_symbol = '.'
trap_symbol = '^'

def row_to_bools(row):
    return [False if c == trap_symbol else True for c in row]

def get_tiles(row):
    for i in range(len(row)):
        if i == 0:
            yield True, row[0], row[1]
        elif i == len(row) - 1:
            yield row[-2], row[-1], True
        else:
            yield row[i-1], row[i], row[i+1]

def is_tile_safe(left, center, right):
    if not left and not center and right:
        return False
    
    if not center and not right and left:
        return False

    if not left and center and right:
        return False

    if not right and center and left:
        return False

    return True

starting_row = '.^^..^...^..^^.^^^.^^^.^^^^^^.^.^^^^.^^.^^^^^^.^...^......^...^^^..^^^.....^^^^^^^^^....^^...^^^^..^'
row = row_to_bools(starting_row)
rows = 0
total_safe = sum(row)
while (rows := rows + 1) < 400000:
    row = [is_tile_safe(left, center, right) for left, center, right in get_tiles(row)]
    total_safe += sum(row) 

print(total_safe)
