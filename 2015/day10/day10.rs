fn main() {
    part1();
}

fn part1() {
    let mut input = "1321131112 ".to_string();

    for _ in 0..50 {
        input = run_process(input);
    }

    println!("{:?}", input.len() - 1);
}

fn run_process(input: String) -> String {
    let mut output: String = String::new();
    let mut count = 0;
    let mut last_ch: char = input.chars().next().unwrap();

    for ch in input.chars() {
        if ch != last_ch {
            output += &format!("{}{}", count, last_ch);
            count = 1;
        } else {
            count += 1;
        }
        last_ch = ch;
    }

    output + " "
}
