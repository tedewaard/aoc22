use std::collections::HashMap;


fn main() {
    //let input = std::fs::read_to_string("test").expect("Failed to read in file");
    let input = std::fs::read_to_string("input").expect("Failed to read in file");
    /*
     * lost -> 0
     * tied -> 3
     * win -> 6
     *
     * Rock -> 1
     * Paper -> 2
     * Scissors -> 3
     */

    let mut plays = HashMap::new();
    //Part 1
    /*
    plays.insert("AX", 4);
    plays.insert("AY", 8);
    plays.insert("AZ", 3);
    plays.insert("BX", 1);
    plays.insert("BY", 5);
    plays.insert("BZ", 9);
    plays.insert("CX", 7);
    plays.insert("CY", 2);
    plays.insert("CZ", 6);
    */ 
    //Part 2
    plays.insert("AX", 3);
    plays.insert("AY", 4);
    plays.insert("AZ", 8);
    plays.insert("BX", 1);
    plays.insert("BY", 5);
    plays.insert("BZ", 9);
    plays.insert("CX", 2);
    plays.insert("CY", 6);
    plays.insert("CZ", 7);

    //let mut scores: Vec<u8> = Vec::new();
    let mut sum = 0;

    for line in input.lines() {
        let moves = line.replace(" ", "");
        sum += *plays.get(&*moves).unwrap();
        //scores.push(*plays.get(&*moves).unwrap());



        //println!("{:?}", moves);
    }
    println!("{}", sum);
}
