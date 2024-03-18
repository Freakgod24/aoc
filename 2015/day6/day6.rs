use fancy_regex::Regex;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    part1();
    part2();
}

fn part1() {
    let width : usize = 1000;
    let height : usize  = 1000;
    let mut array = vec![vec![0; width]; height];
    
    let re1 = Regex::new(r"turn on (\d+),(\d+) through (\d+),(\d+)").unwrap();
    let re2 = Regex::new(r"turn off (\d+),(\d+) through (\d+),(\d+)").unwrap();
    let re3 = Regex::new(r"toggle (\d+),(\d+) through (\d+),(\d+)").unwrap();

    if let Ok(lines) = read_lines("./day6.txt") {
        for line in lines.flatten() {
            if let Some(coords) = get_coordinates(&re1, &line) {
                apply_turn_on_part1(&mut array, coords);
            } else if let Some(coords) = get_coordinates(&re2, &line) {
                apply_turn_off_part1(&mut array, coords);
            } else if let Some(coords) = get_coordinates(&re3, &line) {
                apply_turn_toggle_part1(&mut array, coords);
            }
        }
    }

    let total_lights: usize = array.iter().map(|row| row.iter().sum::<usize>()).sum();
    println!("{}", total_lights);
}

fn part2()
{
    let width : usize = 1000;
    let height : usize  = 1000;
    let mut array = vec![vec![0; width]; height];
    
    let re1 = Regex::new(r"turn on (\d+),(\d+) through (\d+),(\d+)").unwrap();
    let re2 = Regex::new(r"turn off (\d+),(\d+) through (\d+),(\d+)").unwrap();
    let re3 = Regex::new(r"toggle (\d+),(\d+) through (\d+),(\d+)").unwrap();

    if let Ok(lines) = read_lines("./day6.txt") {
        for line in lines.flatten() {
            
            if let Some(coords) = get_coordinates(&re1, &line) {
                apply_turn_on_part2(&mut array, coords);
            } else if let Some(coords) = get_coordinates(&re2, &line) {
                apply_turn_off_part2(&mut array, coords);
            } else if let Some(coords) = get_coordinates(&re3, &line) {
                apply_turn_toggle_part2(&mut array, coords);
            }
        }
    }
    let total_intensity: usize = array.iter().map(|row| row.iter().sum::<usize>()).sum();
    println!("{}", total_intensity);
}

fn get_coordinates(re: &Regex, line: &str) -> Option<(usize, usize, usize, usize)> {
    let captures = re.captures(&line).ok()??;
    let x1: usize = captures[1].parse().ok()?;
    let y1: usize = captures[2].parse().ok()?;
    let x2: usize = captures[3].parse().ok()?;
    let y2: usize = captures[4].parse().ok()?;
    Some((x1, y1, x2, y2))
}

fn apply_turn_on_part1(array : &mut Vec<Vec<usize>>, coords : (usize,usize,usize,usize)) {
    let (x1, y1, x2, y2) = coords;

    for row in &mut array[y1..=y2] {
        for cell in &mut row[x1..=x2] {
            *cell = 1;
        }
    }
}

fn apply_turn_on_part2(array : &mut Vec<Vec<usize>>, coords : (usize,usize,usize,usize)) {
    let (x1, y1, x2, y2) = coords;

    for row in &mut array[y1..=y2] {
        for cell in &mut row[x1..=x2] {
            *cell += 1;
        }
    }
}

fn apply_turn_off_part1(array : &mut Vec<Vec<usize>>, coords : (usize,usize,usize,usize)) {
    let (x1, y1, x2, y2) = coords;

    for row in &mut array[y1..=y2] {
        for cell in &mut row[x1..=x2] {
            *cell = 0;
        }
    }
}

fn apply_turn_off_part2(array : &mut Vec<Vec<usize>>, coords : (usize,usize,usize,usize)) {
    let (x1, y1, x2, y2) = coords;

    for row in &mut array[y1..=y2] {
        for cell in &mut row[x1..=x2] {
            if *cell > 0 {
                *cell -= 1;
            }
        }
    }
}

fn apply_turn_toggle_part1(array : &mut Vec<Vec<usize>>, coords : (usize,usize,usize,usize)) {
    let (x1, y1, x2, y2) = coords;

    for row in &mut array[y1..=y2] {
        for cell in &mut row[x1..=x2] {
            *cell ^= 1;
        }
    }
}

fn apply_turn_toggle_part2(array : &mut Vec<Vec<usize>>, coords : (usize,usize,usize,usize)) {
    let (x1, y1, x2, y2) = coords;

    for row in &mut array[y1..=y2] {
        for cell in &mut row[x1..=x2] {
            *cell += 2;
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
