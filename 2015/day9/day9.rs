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

    let mut distances: HashMap<String, HashMap<String, i32>> = HashMap::new();

    // Since the dataset is small, I will use a HashMap instead of a Graph.
    // The keys will be composed of the name of the cities.
    for line in reader.lines().flatten() {
        let (src, _, dst, _, value) = line.split_whitespace().collect_tuple().unwrap();
        let distance = value.parse().unwrap_or(0);

        distances
            .entry(src.to_string())
            .or_insert_with(HashMap::new)
            .insert(dst.to_string(), distance);

        distances
            .entry(dst.to_string())
            .or_insert_with(HashMap::new)
            .insert(src.to_string(), distance);
    }

    // Again, considering that the dataset is relatively small, I will bruteforce the solution.
    let k = distances.keys().count();
    let total_distances = distances
        .keys()
        .into_iter()
        .permutations(k)
        .map(|route| -> i32 {
            route
                .windows(2)
                .map(|city| distances[city[0]][city[1]])
                .sum()
        });

    let min_distance = total_distances.clone().min();
    let max_distance = total_distances.clone().max();

    println!("{:?}", min_distance);
    println!("{:?}", max_distance);
}
