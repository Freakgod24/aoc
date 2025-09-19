def read_file(filepath: str):
    with open(filepath, 'r') as file:
        for line in file:
            start, stop = line.strip().split('-')
            yield int(start), int(stop) 

value = sorted(read_file('input.txt'))
ip = 0
count = 0
for start, end in value:
    print(start, end, ip, count)
    if start > ip:
        #break
        count +=  start - ip
        ip = end+1
    elif end > ip:
        ip = end+1
