#offsets = [0,3,0,1,-3]
offsets=[int(line) for line in open('input.txt').readlines()]
index = 0
steps = 0

while len(offsets) > index >= 0:
    jump = offsets[index]
    if jump >= 3:
        offsets[index] -= 1
    else:
        offsets[index] += 1
    index = index + jump
    steps += 1

print(steps)
