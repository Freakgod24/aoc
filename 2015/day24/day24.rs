use itertools::Itertools;

fn main() {
    part1();
    part2();
}

fn part1() {
    let packages = vec![
        1, 2, 3, 7, 11, 13, 17, 19, 23, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
        101, 103, 107, 109, 113,
    ];

    let total_weight: i64 = packages.iter().sum();
    let target_weight = total_weight / 3;

    let qe = packages
        .into_iter()
        .combinations(6)
        .filter(|c| c.iter().sum::<i64>() == target_weight)
        .map(|g| g.iter().product::<i64>())
        .min()
        .unwrap();

    println!("{}", qe);
}

fn part2() {
    let packages = vec![
        1, 2, 3, 7, 11, 13, 17, 19, 23, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
        101, 103, 107, 109, 113,
    ];

    let total_weight: i64 = packages.iter().sum();
    let target_weight = total_weight / 4;

    let qe = packages
        .into_iter()
        .combinations(4)
        .filter(|c| c.iter().sum::<i64>() == target_weight)
        .map(|g| g.iter().product::<i64>())
        .min()
        .unwrap();

    println!("{}", qe);
}
