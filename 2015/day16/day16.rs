use regex::Regex;
use std::{collections::HashMap, fs};

#[derive(Debug, Clone)]
struct Sue {
    number: i32,
    compounds: HashMap<String, i32>,
}

impl Sue {
    fn new(
        number: i32,
        compound_1: String,
        qty_1: i32,
        compound_2: String,
        qty_2: i32,
        compound_3: String,
        qty_3: i32,
    ) -> Self {
        let mut compounds = HashMap::new();
        compounds.insert(compound_1, qty_1);
        compounds.insert(compound_2, qty_2);
        compounds.insert(compound_3, qty_3);

        Self { number, compounds }
    }

    fn is_matching_part1(&self, other_compounds: &HashMap<&str, i32>) -> bool {
        for (compound, qty) in &self.compounds {
            if let Some(other_qty) = other_compounds.get(compound.as_str()) {
                if other_qty != qty {
                    return false;
                }
            }
        }
        return true;
    }

    fn is_matching_part2(&self, other_compounds: &HashMap<&str, i32>) -> bool {
        for (compound, qty) in &self.compounds {
            if let Some(other_qty) = other_compounds.get(compound.as_str()) {
                let res = match compound.as_str() {
                    "cats" => qty <= other_qty,
                    "trees" => qty <= other_qty,
                    "pomerians" => qty >= other_qty,
                    "goldfish" => qty >= other_qty,
                    _ => qty != other_qty,
                };
                if res == true {
                    return false;
                }
            }
        }
        return true;
    }
}

impl From<regex::Captures<'_>> for Sue {
    fn from(item: regex::Captures<'_>) -> Self {
        Sue::new(
            item[1].parse().unwrap(),
            item[2].to_string(),
            item[3].parse().unwrap(),
            item[4].to_string(),
            item[5].parse().unwrap(),
            item[6].to_string(),
            item[7].parse().unwrap(),
        )
    }
}

fn main() {
    let data = fs::read_to_string("./day16.txt").expect("File not found");
    let mut compounds: HashMap<&str, i32> = HashMap::new();
    compounds.insert("children", 3);
    compounds.insert("cats", 7);
    compounds.insert("samoyeds", 2);
    compounds.insert("pomeranians", 3);
    compounds.insert("akitas", 0);
    compounds.insert("vizslas", 0);
    compounds.insert("goldfish", 5);
    compounds.insert("trees", 3);
    compounds.insert("cars", 2);
    compounds.insert("perfumes", 1);

    part1(&data, &compounds);
    part2(&data, &compounds);
}

fn part1(data: &String, compounds: &HashMap<&str, i32>) {
    let regex = Regex::new(r"Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)").unwrap();

    data.lines()
        .map(|line| -> Sue { regex.captures(line).expect("Invalid data").into() })
        .filter(|sue| sue.is_matching_part1(&compounds))
        .for_each(|sue| println!("{:?}", sue.number));
}

fn part2(data: &String, compounds: &HashMap<&str, i32>) {
    let regex = Regex::new(r"Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)").unwrap();

    data.lines()
        .map(|line| -> Sue { regex.captures(line).expect("Invalid data").into() })
        .filter(|sue| sue.is_matching_part2(&compounds))
        .for_each(|sue| println!("{:?}", sue.number));
}
