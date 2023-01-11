fn main() {
    let input = std::fs::read_to_string("input").expect("Should have been able to read the file");
    //    println!("{}", input);

    let mut sum = 0;
    let mut totals = Vec::new();
    for line in input.lines() {
        if line == "" {
            totals.push(sum);
            sum = 0;
        } else {
            sum += line.parse::<usize>().unwrap();
        }
    }
    totals.push(sum);
    totals.sort();
    totals.reverse();

    println!("{:?}", &totals[..3].iter().sum::<usize>());
}
