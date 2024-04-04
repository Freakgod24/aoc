use itertools::Itertools;
use std::fs;

fn main() {
    test();
    part1();
    part2();
}

fn part1() {
    let data = fs::read_to_string("./day17.txt").expect("File not found");
    let containers: Vec<i32> = data.lines().map(|line| line.parse().unwrap()).collect();
    let mut total = 0;

    for k in 2..containers.len() {
        containers
            .iter()
            .combinations(k)
            .filter(|p| p.iter().copied().sum::<i32>() == 150)
            .for_each(|p| {
                total += 1;
                // println!("{:?}", p);
            });
    }

    println!("{}", total);
}

fn part2() {
    let data = fs::read_to_string("./day17.txt").expect("File not found");
    let containers: Vec<i32> = data.lines().map(|line| line.parse().unwrap()).collect();
    let mut total = 0;

    for k in 2..containers.len() {
        containers
            .iter()
            .combinations(k)
            .filter(|p| p.iter().copied().sum::<i32>() == 150)
            .for_each(|p| {
                total += 1;
                // println!("{:?}", p);
            });

        if total > 1 {
            break;
        }
    }

    println!("{}", total);
}

fn test() {
    let test_containers = [20, 15, 10, 5, 5].to_vec();
    let mut total = 0;

    for k in 2..test_containers.len() {
        test_containers
            .iter()
            .combinations(k)
            .filter(|p| p.iter().copied().sum::<i32>() == 25)
            .for_each(|p| {
                total += 1;
                // println!("{:?}", p);
            });
    }

    println!("{}", total);
}
