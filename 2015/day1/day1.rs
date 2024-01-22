use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    part1();
    part2();
}

fn part1() {
    let mut floor = 0;
    if let Ok(lines) = read_lines("./day1.txt") {
        for line in lines.flatten() {
            for char in line.chars() {
                floor += match char {
                    '(' => 1,
                    ')' => -1,
                    _ => 0,
                }
            }
        }
        println!("{}", floor);
    }
}

fn part2() {
    let mut floor = 0;
    let mut position = 0;
    if let Ok(lines) = read_lines("./day1.txt") {
        for line in lines.flatten() {
            for char in line.chars() {
                position += 1;
                floor += match char {
                    '(' => 1,
                    ')' => -1,
                    _ => 0,
                };
                if floor == -1 {
                    println!("{}", position);
                    return
                }
            }
        }
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
