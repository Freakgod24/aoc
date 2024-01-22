use std::collections::HashSet;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    part1();
    part2();
}

fn part1() {
    let mut houses = HashSet::new();
    let mut current_position = (0, 0);

    houses.insert(current_position);

    if let Ok(lines) = read_lines("./day3.txt") {
        for line in lines.flatten() {
            for char in line.chars() {
                match char {
                    '^' => current_position.1 += 1,
                    'v' => current_position.1 -= 1,
                    '>' => current_position.0 += 1,
                    '<' => current_position.0 -= 1,
                    _ => continue,
                };
                houses.insert(current_position);
            }
        }
        println!("{}", houses.len());
    }
}

fn part2() {
    let mut houses = HashSet::new();
    let mut santa_position = (0, 0);
    let mut robo_position = (0, 0);

    houses.insert((0,0));

    if let Ok(lines) = read_lines("./day3.txt") {
        for line in lines.flatten() {
            for (i, char) in line.chars().enumerate() {
                let current_position = if i % 2 == 0{
                    &mut santa_position
                } else {
                    &mut robo_position
                };
                match char {
                    '^' => current_position.1 += 1,
                    'v' => current_position.1 -= 1,
                    '>' => current_position.0 += 1,
                    '<' => current_position.0 -= 1,
                    _ => continue,
                };
                houses.insert(*current_position);
            }
        }
        println!("{}", houses.len());
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
