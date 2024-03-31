use regex::Regex;
use serde_json::{Map, Value};
use std::fs;

fn main() {
    part1();
    part2();
}

fn part1() {
    let data = fs::read_to_string("./day12.txt").expect("File not found");
    let regex = Regex::new(r"-?\d+").unwrap();

    let result: i32 = regex
        .find_iter(data.as_str())
        .map(|m| -> i32 { m.as_str().parse().unwrap() })
        .sum();

    println!("{}", result);
}

fn part2() {
    let data = fs::read_to_string("./day12.txt").expect("File not found");
    let json: Value = serde_json::from_str(data.as_str()).expect("Invalid JSON format");

    println!("{}", parse_value(&json));
}

fn parse_value(json: &Value) -> i64 {
    let mut sum = 0;

    if json.is_object() {
        let object: Map<String, Value> = serde_json::from_value(json.clone()).unwrap();

        let contains_red = object
            .values()
            .filter(|v| match v.as_str() {
                Some("red") => true,
                _ => false,
            })
            .count()
            > 0;

        if !contains_red {
            sum += object.values().map(|v| parse_value(v)).sum::<i64>();
        }
    } else if json.is_array() {
        let array: Vec<Value> = serde_json::from_value(json.clone()).unwrap();
        sum += array.into_iter().map(|v| parse_value(&v)).sum::<i64>();
    } else if json.is_number() {
        sum += json.as_i64().unwrap();
    }

    return sum;
}
