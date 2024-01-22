use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    part1();
    part2();
}

fn part1() {
    let mut total_paper = 0;
    if let Ok(lines) = read_lines("./day2.txt") {
        for line in lines.flatten() {
            let dimensions: Vec<&str> = line.split('x').collect();

            let l: i32 = dimensions[0].parse().unwrap();
            let w: i32 = dimensions[1].parse().unwrap();
            let h: i32 = dimensions[2].parse().unwrap();

            let surface_area = 2 * l * w + 2 * w * h + 2 * h * l;
            let sides_area = [l * w, w * h, h * l];
            let slack = sides_area.iter().min().unwrap();

            total_paper += surface_area + slack;
        }
        println!("{}", total_paper);
    }
}

fn part2() {
    let mut total_ribbon = 0;
    if let Ok(lines) = read_lines("./day2.txt") {
        for line in lines.flatten() {
            let dimensions: Vec<&str> = line.split('x').collect();

            let l: i32 = dimensions[0].parse().unwrap();
            let w: i32 = dimensions[1].parse().unwrap();
            let h: i32 = dimensions[2].parse().unwrap();

            let mut sides = vec![l, w, h];
            let largest_side = sides.iter().max().unwrap();
            let index = sides.iter().position(|&x| x == *largest_side).unwrap();
            sides.remove(index);

            let length_wrap = 2 * sides[0] + 2 * sides[1];
            let length_bow = l * w * h;
            
            total_ribbon += length_wrap + length_bow;
        }
        println!("{}", total_ribbon);
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
