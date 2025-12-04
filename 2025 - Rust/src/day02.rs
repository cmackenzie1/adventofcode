use std::io::{self, BufRead};

fn parse(input: &str) -> Vec<(u64, u64)> {
    input
        .split(",")
        .filter(|l| !l.trim().is_empty())
        .map(|l| {
            let parts = l
                .split("-")
                .filter(|l| !l.trim().is_empty())
                .map(|v| v.parse::<u64>().unwrap())
                .take(2)
                .collect::<Vec<u64>>();

            assert!(parts.len() == 2);
            (parts[0], parts[1])
        })
        .collect()
}

fn solve_part1(ranges: &[(u64, u64)]) -> u64 {
    let mut invalid: Vec<u64> = vec![];

    for range in ranges {
        for i in range.0..=range.1 {
            let s = i.to_string();
            let mid = s.len() / 2;
            if &s[..mid] == &s[mid..] {
                invalid.push(i);
            }
        }
    }

    invalid.iter().sum()
}

fn solve_part2(ranges: &[(u64, u64)]) -> u64 {
    let mut invalid: Vec<u64> = vec![];

    for range in ranges {
        for i in range.0..=range.1 {
            let s = i.to_string();
            let mid = s.len() / 2;

            for j in 1..=mid {
                let needle = &s[0..j];
                println!("s={s} mid={mid} needle={needle}");
                if s.matches(&needle).count() * j == s.len() {
                    println!("invalid={s}");
                    invalid.push(i);
                    break;
                }
            }
        }
    }

    invalid.iter().sum()
}

fn main() {
    let stdin = io::stdin();

    // Process line into ranges
    for line in stdin.lock().lines() {
        let ranges = parse(&line.unwrap());

        let part1 = solve_part1(ranges.as_slice());
        let part2 = solve_part2(ranges.as_slice());

        println!("par1={part1} part2={part2}");
    }
}
