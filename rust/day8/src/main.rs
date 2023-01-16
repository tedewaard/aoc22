

fn main() {
    //let input = std::fs::read_to_string("test").expect("Failed to read in file");
    let input = std::fs::read_to_string("input").expect("Failed to read in file");

    let mut forest: Vec<Vec<u32>> = Vec::new();

    for line in input.lines() {
       //let char_slice: Vec<char> = line.chars().collect();
       let row: Vec<u32> = line.chars().map(|x| x.to_digit(10).unwrap()).collect();

       forest.push(row);
    }

    let perimeter = (forest.len() * 2) + ((forest[0].len()-2) * 2);
    println!("{}", perimeter);

    let mut _visible_count = perimeter;
    let mut scenic_score = Vec::new();
    for i in 1..forest.len()-1 {
        for j in 1..forest[0].len()-1 {
            let horizontal = left_right(&forest[i], j);
            let vertical = up_down(&forest, j, i);
            if horizontal.0 || vertical.0 {
                _visible_count += 1;
            }
            scenic_score.push(horizontal.1 * horizontal.2 * vertical.1 * vertical.2); 
            println!("Coord: {},{} horz: {:?} and vert: {:?}", j, i, horizontal, vertical)
        }
    }
    scenic_score.sort();
    //println!("Part 1: {}", visible_count);
    println!("Part 2: {}", scenic_score.last().unwrap());
    //println!("Part 2: {:?}", scenic_score);
   
}

fn left_right(row: &Vec<u32>, i: usize) -> (bool, u32, u32) {
    let tree_size = row[i];
    let mut left_count = 0;
    let mut right_count = 0;
    let mut left = true;
    let mut right = true;
    for j in (0..i).rev() {
        if row[j] >= tree_size {
            left = false;
            left_count += 1;
            break;
        } else {
            left_count += 1;
        }
    }
    for j in i+1..row.len() {
        if row[j] >= tree_size {
            right = false;
            right_count += 1;
            break;
        } else {
            right_count += 1;
        }
    }
    if left || right {
        return (true, left_count, right_count);
    } else {
        return (false, left_count, right_count);
    }
}


fn up_down(forest: &Vec<Vec<u32>>, x: usize, y: usize) -> (bool, u32, u32) {
    let tree_size = forest[y][x];
    let mut up_count = 0;
    let mut down_count = 0;
    let mut up = true;
    let mut down = true;
    //check up
    for i in (0..y).rev() {
        if forest[i][x] >= tree_size {
            up = false;
            up_count += 1;
            break;
        } else {
            up_count += 1;
        }
    }
    //check down
    for i in y+1..forest.len(){
        if forest[i][x] >= tree_size {
            down = false;
            down_count += 1;
            break;
        } else {
            down_count += 1;
        }
    }

    if up || down {
        return (true, up_count, down_count);
    } else {
        return (false, up_count, down_count);
    }
}
