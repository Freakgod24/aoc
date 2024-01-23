use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use md5;

fn main() {
    part1();
    part2();
}

fn part1() {
    let secret_key = "ckczppom";

    for i in 0..9999999 {
        let data = format!("{}{}", secret_key, i);
        let digest = format!("{:x}", md5::compute(data));
        if digest.starts_with("00000"){
            println!("{}", i);
            return;
        }
    }
}

fn part2() {
    let secret_key = "ckczppom";

    for i in 0..9999999 {
        let data = format!("{}{}", secret_key, i);
        let digest = format!("{:x}", md5::compute(data));
        if digest.starts_with("000000"){
            println!("{}", i);
            return;
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
