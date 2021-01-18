from collections import namedtuple, Counter, defaultdict, deque
from copy import deepcopy
from itertools import combinations, permutations, product, chain, count
from functools import reduce
import operator
import math
import re

cat = " ".join
flatten = chain.from_iterable

def data(day, parser=int, sep="\n"):
    with open(f"day{day}.txt") as f:
        sections = f.read().rstrip().split(sep)
        return map(parser, sections)

def do(day):
    g = globals()
    got = []
    for part in (1, 2):
        if (fname := f"day{day}_{part}") in g:
            res = g[fname](g[f"input_{day}"])
            print(f"{fname} -> {res}")
            got.append(res)
    if (aname := f"answers_{day}") in g:
        assert g[aname] == got
    return got

def mapt(*args, **kwargs):
    return tuple(map(*args, **kwargs))

def prod(iterable):
    return reduce(operator.mul, iterable)

def quantify(iterable, pred=bool):
    return sum(1 for i in iterable if pred(i))

input_1 = set(data(1, int))
def day1_1(nums):
    return next(x * y for x in nums for y in nums & {2020 - x} if x != y)

def day1_2(nums):
    return next(x * y * z for x, y in combinations(nums, 2)
                          for z in nums & {2020 - x - y}
                          if x != z != y)
answers_1 = [969024, 230057040]

def parse_2(line):
    policy, letter, password = line.split()
    mi, ma = map(int, policy.split("-"))
    letter = letter.rstrip(":")
    return (mi, ma), letter, password

input_2 = list(data(2, parse_2))
def day2_1(strs):
    def valid_pw(x):
        (mi, ma), letter, password = x
        c = password.count(letter)
        return mi <= c and c <= ma
    return quantify(strs, valid_pw)

def day2_2(strs):
    def valid_pw(x):
        (i1, i2), letter, password = x
        return (password[i1 - 1] == letter) ^ (password[i2 -1] == letter)
    return quantify(strs, valid_pw)


answers_2 = [622, 263]

input_3 = list(data(3, str))
def day3_1(trees, dx=3, dy=1):
    return quantify(line[i*dx % len(line)] == "#" for i, line in enumerate(trees[::dy]))

def day3_2(trees):
    return prod(day3_1(trees, *d) for d in ((1, 1), (3, 1), (5, 1), (7, 1), (1, 2)))

answers_3 = [232, 3952291680]

def parse_4(line):
    return dict(re.findall(r"([a-z]+):([^\s]+)", cat(line.split("\n"))))
input_4 = list(data(4, parse_4, "\n\n"))

required_fields = {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
def day4_1(passports):
    return quantify(passports, required_fields.issubset)

def day4_2(passports):
    return quantify(passports, validate_passport)

constraints = {
    "byr": lambda x: 1920 <= int(x) <= 2002,
    "iyr": lambda x: 2010 <= int(x) <= 2020,
    "eyr": lambda x: 2020 <= int(x) <= 2030,
    "hgt": lambda x: (x.endswith("cm") and 150 <= int(x[:-2]) <= 193) or
                     (x.endswith("in") and 59 <= int(x[:-2]) <= 76),
    "hcl": lambda x: re.search(r"^#[a-f0-9]{6}$", x),
    "ecl": lambda x: x in {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"},
    "pid": lambda x: re.search(r"^[0-9]{9}$", x)
}
def validate_passport(passport):
    return all(field in passport and constraints[field](passport[field]) for field in required_fields)


answers_4 = [204, 179]

def seat_id(seats, table=str.maketrans("FLBR", "0011")):
    return int(seats.translate(table), base=2)
input_5 = list(data(5, seat_id))
# else unnecessary function wrapping
day5_1 = max
def day5_2(data):
    seats = set(range(min(data), max(data)))
    [missing] = seats.difference(set(data))
    return missing

answers_5 =[861, 633]

input_6 = list(data(6, lambda x: x.split("\n"), "\n\n"))
def day6_1(groups):
    return sum(len(set("".join(group))) for group in groups)
def day6_2(groups):
    # return sum(len(reduce(lambda x, y: x.intersection(y), map(set, group))) for group in groups)
    return sum(len(set.intersection(*map(set, group))) for group in groups)

answers_6 = [6437, 3229]

color_map = {}
def parse_7():
    filename = "day7.txt"
    adj = []
    with open(filename) as f:
        for y, line in enumerate(f):
            adj.append([])
            color = re.search(r"^\w+\s\w+", line).group().strip()
            color_map[color] = y
    with open(filename) as f:
        for y, line in enumerate(f):
            _, edges = line.split("contain")
            edges = re.findall(r"\d+\s\w+\s\w+", edges)
            for edge in edges:
                num, *name = edge.split()
                name = cat(name).strip()
                # adj[color_map[name]].append((y, int(num)))
                adj[y].append((color_map[name], int(num)))
        return adj
input_7 = parse_7()

def day7_1(adj):
    start = color_map["shiny gold"]
    n = len(adj)
    visited = [False for i in adj]
    visited[start] = True
    def dfs(x):
        v, w = x
        visited[v] = True
        for n in adj[v]:
            dfs(n)
    dfs((start, 0))
    return quantify(visited) - 1

def day7_2(adj):
    def sum_bags(bag):
        return sum(c + c * sum_bags(b) for (b, c) in adj[bag] if c > 0)
    return sum_bags(color_map["shiny gold"])

def parse_8(line):
    instr, num = line.split()
    return instr, int(num)
input_8 = list(data(8, parse_8))

def day8_1(program):
    acc = 0
    visited = [False for line in program]
    def step(pointer, acc):
        if visited[pointer]:
            return acc
        visited[pointer] = True
        instr, num = program[pointer]
        if instr == "acc":
            return step(pointer + 1, acc + num)
        elif instr == "jmp": 
            return step(pointer + num, acc)
        else: # NOP
            return step(pointer + 1, acc)
    return step(0, 0)

# TODO: concise, efficient
def day8_2(program):
    acc = 0
    history = []
    def step(pointer, acc, visited, hist=False):
        if pointer == len(program):
            return True, acc
        if visited[pointer]:
            return None
        if hist:
            history.append((pointer, program[pointer]))
        visited[pointer] = True
        instr, num = program[pointer]
        if instr == "acc":
            return step(pointer + 1, acc + num, visited, hist)
        elif instr == "jmp": 
            return step(pointer + num, acc, visited, hist)
        else: # NOP
            return step(pointer + 1, acc, visited, hist)
    visited = [False for line in program]
    step(0, 0, visited[::], True)
    while history:
        i, (instr, num) = history.pop()
        if instr == "acc":
            continue
        elif instr == "nop":
            program[i] = "jmp", num
            if (res := step(0, 0, visited[::])) is not None:
                print(res)
                return res
            program[i] = "nop", num
        elif instr == "jmp":
            program[i] = "nop", num
            if (res := step(0, 0, visited[::])) is not None:
                print(res)
                return res
            program[i] = "jmp", num
    return False

# TODO Insert from other

input_10 = list(data(10))
def day10_2(adapters):
    adapters = list(sorted(adapters + [0]))
    dp = [0 for _ in adapters]
    dp[0] = 1
    for i, a in enumerate(adapters):
        for j in (1, 2, 3):
            if i - j >= 0 and adapters[i - j] + 3 >= a:
                dp[i] += dp[i - j]
    return dp[-1]

answers_10 = [347250213298688]

input_11 = list(data(11, list))
def day11_1(state):
    """
    If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
    If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
    Otherwise, the seat's state does not change.
    """
    new_state = apply_rules(state)
    while state != new_state:
        state = new_state
        new_state = apply_rules(new_state)
    return sum(line.count("#") for line in state)

def occupied_neighbors(state, x, y):
    occupied = 0
    for dy in (-1, 0, 1):
        for dx in (-1, 0, 1):
            if (dy == dx) and dx == 0:
                continue
            try:
                occupied += state[y + dy][x + dx] == "#"
            except:
                pass
    return occupied

def apply_rules(state):
    new_state = deepcopy(state)
    for y, line in enumerate(state):
        for x, c in enumerate(line):
            neighbors = occupied_neighbors(state, x, y)
            if c == "L" and neighbors == 0:
                new_state[y][x] = "#"
            elif c == "#" and neighbors >= 4:
                new_state[y][x] = "L"
    return new_state

do(11)
