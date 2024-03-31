fn main() {
    part1();
    part2();
}

fn part1() {
    let mut highest_score: i32 = 0;

    for i in 0..100 {
        for j in 0..100 - i {
            for k in 0..100 - i - j {
                let h = 100 - i - j - k;
                let capacity: i32 = (4 * i - k).max(0);
                let durability: i32 = (-2 * i + 5 * j).max(0);
                let flavour: i32 = (-j + 5 * k - 2 * h).max(0);
                let texture: i32 = (2 * h).max(0);
                let score = capacity * durability * flavour * texture;
                if score > highest_score {
                    highest_score = score;
                }
            }
        }
    }

    println!("{}", highest_score);
}

fn part2() {
    let mut highest_score: i32 = 0;

    for i in 0..100 {
        for j in 0..100 - i {
            for k in 0..100 - i - j {
                let h = 100 - i - j - k;
                let capacity: i32 = (4 * i - k).max(0);
                let durability: i32 = (-2 * i + 5 * j).max(0);
                let flavour: i32 = (-j + 5 * k - 2 * h).max(0);
                let texture: i32 = (2 * h).max(0);
                let calories: i32 = 5 * i + 8 * j + 6 * k + h;
                let score = capacity * durability * flavour * texture;
                if calories == 500 && score > highest_score {
                    highest_score = score;
                }
            }
        }
    }

    println!("{}", highest_score);
}
