use std::collections::HashMap;
use std::collections::HashSet;


fn main() {
    //let input = std::fs::read_to_string("test").expect("Failed to read in file");
    let input = std::fs::read_to_string("input").expect("Failed to read in file");


    let mut priority = HashMap::new();

    priority.insert('a', 1);
    priority.insert('b', 2);
    priority.insert('c', 3);
    priority.insert('d', 4);
    priority.insert('e', 5);
    priority.insert('f', 6);
    priority.insert('g', 7);
    priority.insert('h', 8);
    priority.insert('i', 9);
    priority.insert('j', 10);
    priority.insert('k', 11);
    priority.insert('l', 12);
    priority.insert('m', 13);
    priority.insert('n', 14);
    priority.insert('o', 15);
    priority.insert('p', 16);
    priority.insert('q', 17);
    priority.insert('r', 18);
    priority.insert('s', 19);
    priority.insert('t', 20);
    priority.insert('u', 21);
    priority.insert('v', 22);
    priority.insert('w', 23);
    priority.insert('x', 24);
    priority.insert('y', 25);
    priority.insert('z', 26);
    priority.insert('A', 27);
    priority.insert('B', 28);
    priority.insert('C', 29);
    priority.insert('D', 30);
    priority.insert('E', 31);
    priority.insert('F', 32);
    priority.insert('G', 33);
    priority.insert('H', 34);
    priority.insert('I', 35);
    priority.insert('J', 36);
    priority.insert('K', 37);
    priority.insert('L', 38);
    priority.insert('M', 39);
    priority.insert('N', 40);
    priority.insert('O', 41);
    priority.insert('P', 42);
    priority.insert('Q', 43);
    priority.insert('R', 44);
    priority.insert('S', 45);
    priority.insert('T', 46);
    priority.insert('U', 47);
    priority.insert('V', 48);
    priority.insert('W', 49);
    priority.insert('X', 50);
    priority.insert('Y', 51);
    priority.insert('Z', 52);

    let mut lines = Vec::new();



    for line in input.lines() {
        let l1: HashSet<char> = line.chars().collect();

        lines.push(l1);
    }

    let mut sum = 0;

    //println!("{:?}", lines);

    //We want to iterate over the lines vec and grab three at a time
    for line in lines.chunks(3) {
        let a = &line[0];
        let b = &line[1];
        let c = &line[2];
        //Now I have the first three hashsets so I need to figure out common letter between first
        //two
        let intersection1: HashSet<char> = a.intersection(&b).cloned().collect();
        let intersection2: HashSet<char> = intersection1.intersection(&c).cloned().collect();

        for v in intersection2 {
            sum += priority.get(&v).unwrap();
        }
    }



    /* Part 1
    for line in input.lines() {
        let half = (line.chars().count())/2;
        //println!("Length is {} and half is {}", len, half)

        let first = &line[..half];
        let second = &line[half..];

        'outer: for a in first.chars() {
            for b in second.chars() {
                if a == b {
                    sum += priority.get(&a).unwrap();
                    println!("{}, {a}", priority.get(&a).unwrap());
                    break 'outer;
                }
            }
        }

    }
    */
    println!("{}", sum);
}
