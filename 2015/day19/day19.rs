use std::{
    collections::{HashMap, HashSet},
    fs,
};

fn main() {
    test();
    part1();
    part2();
}

fn part1() {
    let data = fs::read_to_string("./day19.txt").expect("File not found");
    let mut replacements = Vec::new();
    let mut molecule = "";

    for line in data.lines() {
        if line.contains("=>") {
            let r = line.split("=>").map(|s| s.trim()).collect::<Vec<&str>>();
            replacements.push((r[0], r[1]));
        } else if !line.is_empty() {
            molecule = line.clone();
        }
    }

    let molecules = get_molecules(molecule, &replacements);
    println!("{}", molecules.len());
}

fn test() {
    let replacements = vec![("H", "HO"), ("H", "OH"), ("O", "HH")];
    let test_string = "HOH";

    println!("{:?}", get_molecules(test_string, &replacements));
}

fn get_molecules(molecule: &str, replacements: &Vec<(&str, &str)>) -> HashSet<String> {
    let mut molecules = HashSet::new();

    for (i, c) in molecule.chars().enumerate() {
        let current_str = molecule[i..].to_string();
        for (rc, rs) in replacements {
            if current_str.starts_with(rc) {
                let mut new_str = molecule.to_string();
                let from = i;
                let to = i + rc.len() - 1;

                if to < new_str.len() {
                    new_str.replace_range(from..=to, rs);
                    molecules.insert(new_str);
                }
            }
        }
    }

    return molecules;
}

fn part2() {
    let data = fs::read_to_string("./day19.txt").expect("File not found");
    let mut replacements = Vec::new();
    let mut molecule = "".to_string();

    for line in data.lines() {
        if line.contains("=>") {
            let r = line.split("=>").map(|s| s.trim()).collect::<Vec<&str>>();
            replacements.push((r[0], r[1]));
        } else if !line.is_empty() {
            molecule = line.to_string();
        }
    }

    // Sort the replacements from longuest molecule to the shortest. This way, when testing each replacement,
    // the longuest molecule will be tested first maximizing the reduction in length.
    replacements.sort_by(|a, b| b.1.len().cmp(&a.1.len()));

    // Instead of going from "e" to the molecule, we will instead solve this backwards going from
    // the molecule to "e". Taking a long string and minimizing the length seems a better approach than
    // trying to inflate a string and hope to land on the given molecule.
    let mut total = 0;
    let mut iterations = 0;
    while molecule != "e" {
        for (rc, rs) in &replacements {
            if let Some(i) = molecule.find(rs) {
                let from = i;
                let to = i + rs.len() - 1;
                molecule.replace_range(from..=to, rc);
                total += 1;
            }
        }

        iterations += 1;

        if iterations > 1000 {
            println!("Max iteration reached ! ({})", iterations);
            break;
        }
    }

    // Returning the number of replacements performed. Given our algorithm, we cannot be sure
    // that this is the absolute minimum. It is assumed here that there is only a single
    // path possible between e and the molecule.
    println!("{} {}", molecule, total);
}
