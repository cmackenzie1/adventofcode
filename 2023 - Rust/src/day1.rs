pub fn run(part: Option<u8>) {
    let input = include_str!("../input/day01.txt");
    let input: Vec<String> = input
        .lines()
        .filter(|line| !line.is_empty())
        .map(|line| line.to_string())
        .collect::<Vec<_>>();

    match part {
        Some(1) => part1(&input),
        Some(2) => part2(&input),
        _ => {
            part1(&input);
            part2(&input)
        }
    }
}

fn part1(input: &[String]) {
    let n = input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let l = line
                .to_string()
                .chars()
                .filter(|c| c.is_ascii_digit())
                .collect::<String>();

            let n = l
                .chars()
                .take(1)
                .chain(l.chars().rev().take(1))
                .collect::<String>();

            n.parse::<u32>().unwrap()
        })
        .collect::<Vec<_>>();

    let sum = n.iter().sum::<u32>();
    println!("part 1: {}", sum);
}

fn part2(input: &[String]) {
    let replacements = vec![
        ("sixteen", "#"),
        ("seventeen", "#"),
        ("eighteen", "#"),
        ("nineteen", "#"),
        ("eight", "8"),
        ("one", "1"),
        ("two", "2"),
        ("three", "3"),
        ("four", "4"),
        ("five", "5"),
        ("six", "6"),
        ("seven", "7"),
        ("nine", "9"),
        ("zero", "0"),
    ];

    let n = input
        .iter()
        .map(|line| {
            let s = replacements
                .iter()
                .fold(line.to_string(), |acc, (from, to)| acc.replace(from, to));

            println!("before: {}, after: {}", line, s);
            s
        })
        .map(|line| {
            let l = line
                .to_string()
                .chars()
                .filter(|c| c.is_ascii_digit())
                .collect::<String>();

            let n = l
                .chars()
                .take(1)
                .chain(l.chars().rev().take(1))
                .collect::<String>();

            n.parse::<u32>().unwrap()
        })
        .collect::<Vec<_>>();

    let sum = n.iter().sum::<u32>();
    println!("part 2: {}", sum);
}
