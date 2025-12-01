
def part1():
    with open('day1.txt') as f:
        dial = 50
        password = 0
        for line in f:
            dir = line[0]
            steps = int(line[1:])
            match dir:
                case 'L': dial = (dial - steps) % 100
                case 'R': dial = (dial + steps) % 100

            if dial == 0:
                password += 1

    print(password)

def part2():
    with open('day1.txt') as f:
        dial = 50
        password = 0
        for line in f:
            dir = line[0]
            steps = int(line[1:])
            while steps > 0:
                match dir:
                    case 'L': dial = (dial - 1) % 100
                    case 'R': dial = (dial + 1) % 100

                if dial == 0:
                    password += 1

                steps -= 1

    print(password)


part1()
part2()
