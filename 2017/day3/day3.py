def neighbor_sum(grid, x, y):
    value = ( try_get_neighbor(grid, x-1, y) +
            try_get_neighbor(grid, x, y-1) +
            try_get_neighbor(grid, x+1, y) +
            try_get_neighbor(grid, x, y+1) +
            try_get_neighbor(grid, x-1, y-1) +
            try_get_neighbor(grid, x+1, y-1) +
            try_get_neighbor(grid, x-1, y+1) +
            try_get_neighbor(grid, x+1, y+1) )
    print(value)
    return value

def try_get_neighbor(grid, x, y):
    try:
        return grid[y][x]
    except:
        return 0

def spiral_matrix(n):
    # Create an n×n grid initialized with zeros
    grid = [[0] * n for _ in range(n)]

    # Start in the center
    x = y = n // 2
    grid[y][x] = 1

    # Directions: right, up, left, down
    directions = [(1, 0), (0, -1), (-1, 0), (0, 1)]

    num = 2
    step_size = 1

    while num <= n*n:
        for i, (dx, dy) in enumerate(directions):
            for _ in range(step_size):
                if num > n * n:
                    break
                x += dx
                y += dy
                grid[y][x] = neighbor_sum(grid, x, y)
                num += 1

            # Every two directions, increase step size
            if i % 2 == 1:
                step_size += 1

    return grid


# Example usage
n = 11
matrix = spiral_matrix(n)

# for row in matrix:
#     print(" ".join(f"{v:2}" for v in row))
