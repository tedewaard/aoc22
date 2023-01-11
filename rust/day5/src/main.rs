use std::collections::HashMap;


fn main() {
    //let input = std::fs::read_to_string("test").expect("Failed to read in file");
    let input = std::fs::read_to_string("input").expect("Failed to read in file");

    let mut moves: Vec<Vec<&str>> = Vec::new();
    let mut cargo = Vec::new();
    let mut map: HashMap<usize, Vec<char>> = HashMap::new();

    //Parsing input into different vecs
    for lines in input.lines() {
        //Parsing my moves
        if lines.starts_with("move") {
            let split_lines: Vec<&str> = lines.split(" ").collect();
            let count = split_lines[1];
            let from = split_lines[3];
            let to = split_lines[5];
            let mut s = Vec::new();
            s.push(count);
            s.push(from);
            s.push(to);
            moves.push(s);
        } else if !lines.is_empty() {
            //Parsing my cargo/crates
            let newline = lines.replace(&['[', ']'][..], " ");
            cargo.push(newline);
        }
    }

    //Grabbing the base stacks and adding index to map
    for c in cargo.last().unwrap().chars() {
        if c != ' ' {
            let b = usize::try_from(c.to_digit(10).unwrap()).unwrap();
            map.insert(b, Vec::new());
        }
    }

    let n = cargo.len()-2;

    //Pushes the crates into the map
    for c in (0..=n).rev() {
        for (a, b) in cargo[c].char_indices() {
            if b != ' ' {
                let d = &cargo[cargo.len()-1];
                let x = d.chars().nth(a).unwrap();
                let y = usize::try_from(x.to_digit(10).unwrap()).unwrap();
                map.get_mut(&y).unwrap().push(b);
            }
        }
    }

    //Iterates through the moves and calls the move_those_crates function
    for i in &moves {
        let nums: Vec<usize> = i.iter().map(|s| s.parse::<usize>().unwrap()).collect();
        move_those_crates(&mut map, nums[0], nums[1], nums[2]);
    }

    for (key, value) in map {
        println!("{}: {:?}", key, value.last());
    }
}


fn move_those_crates(hm: &mut HashMap<usize, Vec<char>>, amount: usize, from: usize, to: usize) {
    let mut from_vec = hm.get(&from).cloned().unwrap();
    let mut to_vec = hm.get(&to).cloned().unwrap();
    
    /*PART 1
    for _i in 0..amount {
        let add = from_vec.pop();
        match add {
            Some(v) => to_vec.push(v),
            None => break,
        }
    }
    */

    //Part 2
    let mut new_vec = Vec::new();

    for _i in 0..amount {
        let add = from_vec.pop();
        match add {
            Some(v) => new_vec.push(v),
            None => break,
        }
    }
    new_vec.reverse();
    to_vec.append(&mut new_vec);
    hm.insert(from, from_vec);
    hm.insert(to, to_vec);
}
