use std::fs;

#[derive(Debug)]
struct Computer {
    register_a: i32,
    register_b: i32,
    program_position: i32,
    program_length: i32,
}

fn main() {
    // test();
    part1();
    part2();
}

fn part1() {
    let data = fs::read_to_string("day23.txt").expect("file not found");
    let program: Vec<&str> = data.lines().collect();
    let mut computer = Computer {
        register_a: 0,
        register_b: 0,
        program_position: 0,
        program_length: program.len() as i32,
    };

    while computer.program_position < computer.program_length {
        let instruction = program
            .get(computer.program_position as usize)
            .expect("invalid program position in main")
            .to_owned();

        println!("{}", instruction);

        match instruction {
            i if i.starts_with("hlf") => hlf(i, &mut computer),
            i if i.starts_with("tpl") => tpl(i, &mut computer),
            i if i.starts_with("inc") => inc(i, &mut computer),
            i if i.starts_with("jmp") => jmp(i, &mut computer),
            i if i.starts_with("jie") => jie(i, &mut computer),
            i if i.starts_with("jio") => jio(i, &mut computer),
            _ => panic!("invalid instrution in main"),
        }

        computer.program_position += 1;
    }

    println!("{:?}", computer);
}

fn part2() {
    let data = fs::read_to_string("day23.txt").expect("file not found");
    let program: Vec<&str> = data.lines().collect();
    let mut computer = Computer {
        register_a: 1,
        register_b: 0,
        program_position: 0,
        program_length: program.len() as i32,
    };

    while computer.program_position < computer.program_length {
        let instruction = program
            .get(computer.program_position as usize)
            .expect("invalid program position in main")
            .to_owned();

        println!("{}", instruction);

        match instruction {
            i if i.starts_with("hlf") => hlf(i, &mut computer),
            i if i.starts_with("tpl") => tpl(i, &mut computer),
            i if i.starts_with("inc") => inc(i, &mut computer),
            i if i.starts_with("jmp") => jmp(i, &mut computer),
            i if i.starts_with("jie") => jie(i, &mut computer),
            i if i.starts_with("jio") => jio(i, &mut computer),
            _ => panic!("invalid instrution in main"),
        }

        computer.program_position += 1;
    }

    println!("{:?}", computer);
}

fn hlf(instruction: &str, computer: &mut Computer) {
    let operands: Vec<&str> = instruction.split(" ").collect();
    let register = operands.get(1).expect("invalid hlf instruction").to_owned();
    match register {
        "a" => computer.register_a = computer.register_a / 2,
        "b" => computer.register_b = computer.register_b / 2,
        _ => panic!("invalid register in hlf instruction"),
    }
}

fn tpl(instruction: &str, computer: &mut Computer) {
    let operands: Vec<&str> = instruction.split(" ").collect();
    let register = operands.get(1).expect("invalid tpl instruction").to_owned();
    match register {
        "a" => computer.register_a = computer.register_a * 3,
        "b" => computer.register_b = computer.register_b * 3,
        _ => panic!("invalid register in tpl instruction"),
    }
}

fn inc(instruction: &str, computer: &mut Computer) {
    let operands: Vec<&str> = instruction.split(" ").collect();
    let register = operands.get(1).expect("invalid inc instruction").to_owned();
    match register {
        "a" => computer.register_a = computer.register_a + 1,
        "b" => computer.register_b = computer.register_b + 1,
        _ => panic!("invalid register in inc instruction"),
    }
}

fn jmp(instruction: &str, computer: &mut Computer) {
    let operands: Vec<&str> = instruction.split(" ").collect();
    let offset: i32 = operands
        .get(1)
        .unwrap()
        .parse()
        .expect("invalid jmp instruction");

    let program_position = computer.program_position + offset;

    if program_position < 0 || program_position > computer.program_length {
        panic!("invalid program position in jmp instruction");
    } else {
        computer.program_position += offset - 1;
    }
}

fn jie(instruction: &str, computer: &mut Computer) {
    let instruction_sanitized = instruction.replace(",", "");
    let operands: Vec<&str> = instruction_sanitized.split(" ").collect();

    let register = operands
        .get(1)
        .expect("invalig register in jie instruction")
        .to_owned();

    let offset: i32 = operands
        .get(2)
        .unwrap()
        .parse()
        .expect("invalid offset parameter in jie instruction");

    let program_position = computer.program_position + offset;

    if program_position < 0 || program_position > computer.program_length {
        panic!("invalid program position in jie instruction");
    }

    match register {
        "a" => {
            if computer.register_a % 2 == 0 {
                computer.program_position += offset - 1;
            }
        }
        "b" => {
            if computer.register_b % 2 == 0 {
                computer.program_position += offset - 1;
            }
        }
        _ => panic!("invalid register in jie instruction"),
    }
}

fn jio(instruction: &str, computer: &mut Computer) {
    let instruction_sanitized = instruction.replace(",", "");
    let operands: Vec<&str> = instruction_sanitized.split(" ").collect();

    let register = operands
        .get(1)
        .expect("invalig register in jio instruction")
        .to_owned();

    let offset: i32 = operands
        .get(2)
        .unwrap()
        .parse()
        .expect("invalid offset parameter in jio instruction");

    let program_position = computer.program_position + offset;

    if program_position < 0 || program_position > computer.program_length {
        panic!("invalid program position in jio instruction");
    }

    match register {
        "a" => {
            if computer.register_a == 1 {
                computer.program_position += offset - 1;
            }
        }
        "b" => {
            if computer.register_b == 1 {
                computer.program_position += offset - 1;
            }
        }
        _ => panic!("invalid register in jio instruction"),
    }
}
