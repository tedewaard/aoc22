use std::collections::HashMap;

fn main() {
    let input = std::fs::read_to_string("test").expect("Failed to read in file");
    //let input = std::fs::read_to_string("input").expect("Failed to read in file");

    let mut instructions = Vec::new();

    for line in input.lines() {
        let instuction: Vec<&str> = line.split(" ").collect();
        instructions.push(instuction);
    }

    let mut cycle_count = 0;
    let mut x_register = 1;
    let mut results: HashMap<i32, i32> = HashMap::new(); 

    //Part 2 CRT vecj
    let mut crt = vec![
        vec![".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","."],
        vec![".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","."],
        vec![".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","."],
        vec![".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","."],
        vec![".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","."],
        vec![".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","."],
    ];

    for line in instructions {
        match line[0] {
            "noop" => {
                cycle_count += 1;
                results.insert(cycle_count, x_register);
                if cycle_count >= x_register-1 && cycle_count <= x_register+1 || cycle_count <= x_register+1 && cycle_count >= x_register-1 {
                    update_crt(&mut crt, cycle_count)
                }
                }
            "addx" => {
                cycle_count += 1;
                results.insert(cycle_count, x_register);
                if cycle_count >= x_register-1 && cycle_count <= x_register+1 || cycle_count <= x_register+1 && cycle_count >= x_register-1 {
                    println!("Cycle count: {}, X register: {}", cycle_count, x_register);
                    update_crt(&mut crt, cycle_count)
                }
                cycle_count += 1;
                results.insert(cycle_count, x_register);
                if cycle_count >= x_register-1 && cycle_count <= x_register+1 || cycle_count <= x_register+1 && cycle_count >= x_register-1 {
                    println!("Cycle count: {}, X register: {}", cycle_count, x_register);
                    update_crt(&mut crt, cycle_count)
                }
            }
            _  => panic!("Uh oh"),
        };
        if line.len() == 2 {
            let x = line[1].parse::<i32>().unwrap();
            x_register += x;
        }
    }
    //println!("{:?}", results.get(&100).unwrap());

    let mut answers = Vec::new();
    let mut i = 20;
    while i <= 220 {
        let val = results.get(&i).unwrap(); 
        let ans = val * i;
        answers.push(ans);
        i += 40
    }
    
    //println!("{}", answers.iter().sum::<i32>());
    println!("{:?}", crt);
}

fn update_crt(map: &mut Vec<Vec<&str>>, cycle: i32) {
    let mut y = 0;
    let mut x = cycle-1;
    if cycle <= 40 {
        map[y as usize][x as usize] = "#"; 
    } else if cycle <= 80 {
        x -= 40;
        y = 1;
        map[y as usize][x as usize] = "#"; 
    } else if cycle <= 120 {
        x -= 80;
        y = 2;
        map[y as usize][x as usize] = "#"; 
    } else if cycle <= 160 {
        x -= 120;
        y = 3;
        map[y as usize][x as usize] = "#"; 
    } else if cycle <= 200 {
        x -= 160;
        y = 4;
        map[y as usize][x as usize] = "#"; 
    } else if cycle <= 240 {
        x -= 200;
        y = 5;
        map[y as usize][x as usize] = "#"; 
    }




}
