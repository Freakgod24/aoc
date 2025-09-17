from itertools import islice

disk_size = 35651584  # 272
input = "11101000110010100"


def batch(iterable, n):
    iterator = iter(iterable)
    while batch := tuple(islice(iterator, n)):
        yield batch


def fill_disk(input: str, size: int) -> str:
    a = input[::]
    while len(a) < size:
        b = "".join("0" if c == "1" else "1" for c in reversed(a))
        a = "0".join([a, b])

    return a[:size]


def calculate_checksum(disk_data: str) -> str:
    checksum = disk_data[::]
    while len(checksum) % 2 == 0:
        checksum = ["0" if int(a) ^ int(b) else "1" for a, b in batch(checksum, 2)]

    return "".join(checksum)


disk = fill_disk(input, disk_size)
print(calculate_checksum(disk))
