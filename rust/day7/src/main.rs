

fn main() {
    //let input = std::fs::read_to_string("test").expect("Failed to read in file");
    let input = std::fs::read_to_string("input").expect("Failed to read in file");

    let mut folders = Vec::new();    
    let mut all_folders = Vec::new();    
    let mut small_folders = Vec::new();

    //folders.push(("/", 0));

    for lines in input.lines() {
        if lines.starts_with("$ cd") && lines != "$ cd .."{
            let dir = lines.split_at(4).1;
            folders.push((dir, 0));
        }
        if !lines.starts_with("$") && !lines.starts_with("dir") {
            let size = lines.split(" ").nth(0).unwrap().parse::<i32>().unwrap();
            //We need to iterate over folders and add the size to everything
            for (_d, i) in &mut folders {
                *i += size;
            }
        } 

        if lines == "$ cd .." {
            let n = folders.pop().unwrap();
            all_folders.push(n);
            if n.1 <= 100000 {
                small_folders.push(n);
            }
        }
    }

    let n = folders.pop().unwrap();
    all_folders.push(n);
    if n.1 <= 100000 {
        small_folders.push(n);
    }

    println!("Answer 1: {:?}", small_folders.iter().map(|x| x.1).sum::<i32>());


    let free = 70000000 - folders[0].1;

    let remainder = 30000000 - free;

    let mut viable = Vec::new();
    for (_d, i) in &all_folders {
        if *i >= remainder {
            viable.push(*i);                 
        }
    }
    viable.sort();
    println!("Answer 2: {:?}", viable.first().unwrap()) 
}
