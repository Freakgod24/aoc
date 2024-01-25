use fancy_regex::Regex;
use std::collections::HashSet;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    part1();
    part2();
}

fn part1() {
    let mut naughty = 0;
    let mut nice = 0;

    let re1 = Regex::new(r"[aeiou]").unwrap();
    let re2 = Regex::new(r"(\w)\1").unwrap();
    let re3 = Regex::new(r"ab|cd|pq|xy").unwrap();

    if let Ok(lines) = read_lines("./day5.txt") {
        for line in lines.flatten() {
            if re3.is_match(&line).unwrap() {
                naughty += 1;
                continue;
            }

            if re1.find_iter(&line).count() < 3 || re2.is_match(&line).unwrap() == false {
                naughty += 1;
                continue;
            }

            nice += 1;
        }
        println!("naughty: {}, nice: {}", naughty, nice);
    }
}

fn part2() {
    let mut naughty = 0;
    let mut nice = 0;

    let re1 = Regex::new(r"(\w{2}).*\1").unwrap();
    let re2 = Regex::new(r"(\w).\1").unwrap();

    if let Ok(lines) = read_lines("./day5.txt") {
        for line in lines.flatten() {
            if re1.is_match(&line).unwrap() && re2.is_match(&line).unwrap() == true {
                nice += 1;
                continue;
            }

            naughty += 1;
        }
        println!("naughty: {}, nice: {}", naughty, nice);
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
