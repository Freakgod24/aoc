use regex::Regex;
use std::fs;

#[derive(Debug)]
struct Reindeer {
    pub name: String,
    pub speed: i64,
    pub fly_time: i64,
    pub rest_time: i64,
    pub points: i64,

    current_fly_time: i64,
    current_rest_time: i64,
    travelled_distance: i64,
    resting: bool,
}

impl From<regex::Captures<'_>> for Reindeer {
    fn from(item: regex::Captures<'_>) -> Self {
        Reindeer {
            name: item[1].to_string(),
            speed: item[2].parse().unwrap(),
            fly_time: item[3].parse().unwrap(),
            rest_time: item[4].parse::<i64>().unwrap(),
            points: 0,
            current_fly_time: 0,
            current_rest_time: 0,
            travelled_distance: 0,
            resting: false,
        }
    }
}

impl Reindeer {
    pub fn get_distance_by_time(&mut self, time: i32) -> i64 {
        // println!("Start flying at 0 seconds");
        for t in 0..time {
            self.next();
        }
        return self.travelled_distance;
    }

    pub fn next(&mut self) -> () {
        if self.current_fly_time < self.fly_time && self.resting == false {
            self.travelled_distance += self.speed;
            self.current_fly_time += 1;
            // println!("Flying at {} seconds", t);
        } else if self.current_fly_time == self.fly_time && self.resting == false {
            self.resting = true;
            self.current_fly_time = 0;
            self.current_rest_time += 1;
            // println!("Resting at {} seconds", t);
        } else if self.current_rest_time < self.rest_time && self.resting == true {
            self.current_rest_time += 1;
            // println!("Resting at {} seconds", t);
        } else if self.current_rest_time == self.rest_time && self.resting == true {
            self.resting = false;
            self.current_rest_time = 0;
            self.current_fly_time += 1;
            self.travelled_distance += self.speed;
            // println!("Flying at {} seconds", t);
        }
        // println!("{},{},{}", t, travelled_distance, resting);
    }
}

fn main() {
    part1();
    part2();
}

fn part1() {
    let data = fs::read_to_string("./day14.txt").expect("File not found");
    let regex = Regex::new(
        r"(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds",
    )
    .unwrap();

    let max_distance: Vec<i64> = data
        .lines()
        .map(|line| -> Reindeer { regex.captures(line).unwrap().into() })
        .map(|mut reindeer| reindeer.get_distance_by_time(2503))
        .collect();

    println!("{:?}", max_distance);
    println!("{}", max_distance.into_iter().max().unwrap());
}

fn part2() {
    let data = fs::read_to_string("./day14.txt").expect("File not found");
    let regex = Regex::new(
        r"(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds",
    )
    .unwrap();

    let mut reindeers: Vec<Reindeer> = data
        .lines()
        .map(|line| -> Reindeer { regex.captures(line).unwrap().into() })
        .collect();

    for _ in 0..2503 {
        // Increment each reindeer racing status.
        reindeers.iter_mut().for_each(|r| r.next());

        // Look-up the current largest distance travelled.
        let current_max_distance = reindeers
            .iter_mut()
            .max_by_key(|r| r.travelled_distance)
            .unwrap()
            .travelled_distance;

        // Give points to leading reindeers. Ties are possible.
        reindeers
            .iter_mut()
            .filter(|r| r.travelled_distance == current_max_distance)
            .for_each(|r| r.points += 1);
    }

    let total_points: Vec<i64> = reindeers.iter_mut().map(|r| r.points).collect();
    println!("{:?}", total_points);
}
