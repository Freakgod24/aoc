use fancy_regex::Regex;
use itertools::Itertools;
use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    part1();
}

fn part1() {
    let file = File::open("./day9.txt").unwrap();
    let reader = BufReader::new(file);

    let mut cities: Vec<String> = Vec::new();
    let mut leg_distances: HashMap<String, i32> = HashMap::new();

    let re1 = Regex::new(r"(\S+) to (\S+) = (\d+)").unwrap();

    // Since the dataset is small, I will use a HashMap instead of a Graph.
    // The keys will be composed of the name of the cities.
    for line in reader.lines().flatten() {
        let captures = re1.captures(&line).unwrap().unwrap();

        let city1 = captures[1].to_string();
        let city2 = captures[2].to_string();
        let leg_distance = captures[3].parse().unwrap_or(0);

        if !cities.contains(&city1) {
            cities.push(city1.clone());
        }

        if !cities.contains(&city2) {
            cities.push(city2.clone());
        }

        let forward = format!("{}{}", city1, city2);
        let reverse = format!("{}{}", city2, city1);
        leg_distances.entry(forward).or_insert(leg_distance);
        leg_distances.entry(reverse).or_insert(leg_distance);
    }

    // Again, considering that the dataset is relatively small, I will bruteforce the solution.
    let k = cities.len();
    let total_distances = cities.into_iter().permutations(k).map(|route| -> i32 {
        route
            .windows(2)
            .map(|city| {
                leg_distances
                    .get(&format!("{}{}", city[0], city[1]))
                    .unwrap()
            })
            .sum()
    });

    let min_distance = total_distances.clone().min();
    let max_distance = total_distances.clone().max();

    println!("{:?}", min_distance);
    println!("{:?}", max_distance);
}
