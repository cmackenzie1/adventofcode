use std::io::{self, BufRead};

const START_POSITION: i32 = 50;

fn main() {
    let mut position = START_POSITION;
    let mut hits = 0;
    let mut hits_p2 = 0;

    let stdin = io::stdin();
    println!("Dial starts at {position}");
    for line in stdin.lock().lines() {
        let line = line.unwrap();
        let (direction, amount) = line.split_at(1);
        let mut amount = amount.parse::<i32>().unwrap();
        match direction {
            "L" => {
                while amount > 0 {
                    position -= 1;
                    if position == 0 {
                        hits_p2 += 1;
                    }
                    if position == -1 {
                        position = 99; // wrap aronud
                    }
                    amount -= 1;
                }
            }
            "R" => {
                while amount > 0 {
                    position += 1;
                    if position == 100 {
                        // wrap around
                        position = 0;
                        hits_p2 += 1;
                    }
                    amount -= 1;
                }
            }
            _ => panic!("Invalid direction"),
        };
        println!("direction={line} result={position} hits_p2={hits_p2}");
        if position == 0 {
            hits += 1;
        }
    }

    println!("final_position={position} hits={hits} hits_p2={hits_p2}");
}
