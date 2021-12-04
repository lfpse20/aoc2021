use std::fs;

fn main() {
    let contents = fs::read_to_string("./input.txt").expect("failed to read file");

    let len = contents.find('\n').unwrap_or_default();

    //initialize vector with initial binary line length
    let mut v = vec![0; len];

    //loop over content and update counter of vector
    let mut i = 0;
    for c in contents.chars() {
        if c == '\n' {
            i = 0;
            continue;
        } else if c == '1' {
            v[i] += 1;
        } else {
            v[i] -= 1;
        }
        i += 1;
    }

    // convert vector to string for gamma value
    let mut gamma = String::from("");
    let mut epsilon = String::from("");
    for bit in v {
        if bit > 0 {
            gamma.push('1');
            epsilon.push('0');
        } else {
            gamma.push('0');
            epsilon.push('1');
        }
    }

    // get decimal value of gamma and epsilon
    let gamma_decimal = isize::from_str_radix(&gamma, 2).unwrap();
    let epsilon_decimal = isize::from_str_radix(&epsilon, 2).unwrap();

    // multiply the 2
    println!("{}",gamma_decimal * epsilon_decimal)

}
