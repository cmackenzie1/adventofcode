use pico_args::Arguments;

mod day1;

struct Args {
    day: u8,
    part: Option<u8>,
}

fn main() {
    let mut args = Arguments::from_env();
    let args = Args {
        day: args
            .value_from_str(["-d", "--day"])
            .expect("day is required"),
        part: args
            .opt_value_from_str(["-p", "--part"])
            .unwrap_or_default(),
    };

    match args.day {
        1 => day1::run(args.part),
        _ => panic!("day not implemented"),
    }
}
