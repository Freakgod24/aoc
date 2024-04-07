use std::fs;

fn main() {
    test();
    part1();
    part2();
}

fn part1() {
    let data = fs::read_to_string("./day18.txt").expect("File not found");
    let grid: Vec<Vec<char>> = data
        .lines()
        .map(|line| -> Vec<char> { line.chars().collect() })
        .collect();

    let mut final_grid = grid.clone();

    for _ in 0..100 {
        final_grid = get_next_grid(&final_grid);
    }

    let total = final_grid
        .iter()
        .flatten()
        .copied()
        .filter(|&c| c == '#')
        .count();

    println!("{:?}", total);
}

fn part2() {
    let data = fs::read_to_string("./day18.txt").expect("File not found");
    let grid: Vec<Vec<char>> = data
        .lines()
        .map(|line| -> Vec<char> { line.chars().collect() })
        .collect();

    let mut final_grid = grid.clone();

    for _ in 0..100 {
        final_grid = get_next_grid(&final_grid);

        // Force the four corners to stay lit.
        // Hardcoded because why not ;)
        final_grid[0][0] = '#';
        final_grid[0][99] = '#';
        final_grid[99][0] = '#';
        final_grid[99][99] = '#';
    }

    let total = final_grid
        .iter()
        .flatten()
        .copied()
        .filter(|&c| c == '#')
        .count();

    println!("{:?}", total);
}

fn test() {
    let mut test_grid = [
        ['.', '#', '.', '#', '.', '#'].to_vec(),
        ['.', '.', '.', '#', '#', '.'].to_vec(),
        ['#', '.', '.', '.', '.', '#'].to_vec(),
        ['.', '.', '#', '.', '.', '.'].to_vec(),
        ['#', '.', '#', '.', '.', '#'].to_vec(),
        ['#', '#', '#', '#', '.', '.'].to_vec(),
    ]
    .to_vec();

    for _ in 0..4 {
        test_grid = get_next_grid(&test_grid);
        test_grid.iter().for_each(|row| println!("{:?}", row));
        println!("");
    }

    let total = test_grid
        .iter()
        .flatten()
        .copied()
        .filter(|&c| c == '#')
        .count();

    println!("{}", total);
}

fn get_next_grid(grid: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    let mut next_grid = grid.clone();

    for i in 0..grid.len() {
        for j in 0..grid.len() {
            let total = get_neighbors_on(&grid, i, j);
            next_grid[i][j] = match (grid[i][j], total) {
                ('#', 2) => '#',
                ('#', 3) => '#',
                ('.', 3) => '#',
                _ => '.',
            }
        }
    }

    return next_grid;
}

fn get_neighbors_on(grid: &Vec<Vec<char>>, i: usize, j: usize) -> i32 {
    let grid_size = grid.len() as isize;
    let mut total_neighbors = 0;

    // Navigate all the neighbors including diagonals.
    for xoff in -1..=1 {
        for yoff in -1..=1 {
            let x = i as isize + xoff;
            let y = j as isize + yoff;

            // Do not count ourselves
            if x == i as isize && y == j as isize {
                continue;
            }

            // Stay within the grid
            if x >= 0 && x < grid_size && y >= 0 && y < grid_size {
                if grid[x as usize][y as usize] == '#' {
                    total_neighbors += 1;
                }
            }
        }
    }

    return total_neighbors;
}
