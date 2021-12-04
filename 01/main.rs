use std::fs;

fn main() {
    let contents = fs::read_to_string("./input.txt").expect("failed to read file");
    let mut v: Vec<i32> = Vec::new();
    for s in contents.lines() {
        v.push(s.parse::<i32>().unwrap());
    }

    let mut sliding_sums: Vec<i32> = Vec::new();
    let mut sum = 0;
    let mut i = 0;
    let mut window_count = 0;
    while i < v.len() {
        sum += v[i];
        if window_count == 2 {
            sliding_sums.push(sum);
            window_count = 0;
            sum = 0;
            i -= 1;
            continue;
        }
        window_count += 1;
        i += 1;
    }

    let mut increased_counter = 0;
    for i in 1..sliding_sums.len() {
        if sliding_sums[i] > sliding_sums[i - 1] {
            increased_counter += 1;
        }
    }

    println!("{}", increased_counter);
}
