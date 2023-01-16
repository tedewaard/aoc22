use std::collections::HashSet;

fn main() {
    //let input = std::fs::read_to_string("test").expect("Failed to read in file");
    //let input = std::fs::read_to_string("test2").expect("Failed to read in file");
    let input = std::fs::read_to_string("input").expect("Failed to read in file");

    let mut rope_moves: Vec<(&str, i32)> = Vec::new();
    let mut knots = vec![
        (0, 0),
        (0, 0),
        (0, 0),
        (0, 0),
        (0, 0),
        (0, 0),
        (0, 0),
        (0, 0),
        (0, 0),
        (0, 0),
    ];
    //let mut head = (0, 0);
    //let mut tail: (i32, i32) = (0, 0);

    for line in input.lines() {
        let line_split: Vec<&str> = line.split(" ").collect();
        let moves = (line_split[0], line_split[1].parse::<i32>().unwrap());
        rope_moves.push(moves);
    }

    let mut tail_visits: HashSet<(i32, i32)>  = HashSet::new();

    for moves in rope_moves {
        println!("{:?}", knots);
        println!("{:?}", moves);
        for _i in 0..moves.1 {
            move_head(&mut knots[0], moves.0);
            for i in 1..10 {
                let last_knot = knots[i-1];
                move_tail(&mut knots[i], last_knot);
            }
            tail_visits.insert(knots[9]);
        }
    }

    tail_visits.insert(knots[9]);
    //println!("{:?}", tail_visits);
    println!("{:?}", tail_visits.len());
}


fn move_head(head: &mut (i32, i32), direction: &str) {
    match direction {
        "L" => head.0 -= 1,
        "R" => head.0 += 1,
        "U" => head.1 += 1,
        "D" => head.1 -= 1,
        _ => panic!("Uh oh error in input"),
    };
}

fn move_tail(tail: &mut (i32, i32), head: (i32, i32)) {
    let x_diff = head.0 - tail.0;
    let y_diff = head.1 - tail.1;

    //Up
    if x_diff == 0 && y_diff > 1 {
        tail.1 += 1;
    }
    //Down
    if x_diff == 0 && y_diff < -1 {
        tail.1 -= 1;
    }
    //Right
    if x_diff > 1 && y_diff == 0 {
        tail.0 += 1;
    }
    //Left
    if x_diff < -1 && y_diff == 0 {
        tail.0 -= 1;
    }
    //Diagnal up and to right
    if x_diff == 1 && y_diff == 2 || x_diff == 2 && y_diff == 1 || x_diff == 2 && y_diff == 2 {
        tail.0 += 1;
        tail.1 += 1;
    }
    //Diagnal up and to left
    if x_diff == -1 && y_diff == 2 || x_diff == -2 && y_diff == 1 || x_diff == -2 && y_diff == 2{
        tail.0 -= 1;
        tail.1 += 1;
    }
    //Diagnal down and to right
    if x_diff == 1 && y_diff == -2 || x_diff == 2 && y_diff == -1 || x_diff == 2 && y_diff == -2{
        tail.0 += 1;
        tail.1 -= 1;
    }
    //Diagnal down and to left
    if x_diff == -1 && y_diff == -2 || x_diff == -2 && y_diff == -1 || x_diff == -2 && y_diff == -2{
        tail.0 -= 1;
        tail.1 -= 1;
    }
}
