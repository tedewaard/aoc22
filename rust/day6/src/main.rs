

fn main() {
    //let input = std::fs::read_to_string("test").expect("Failed to read in file");
    let input = std::fs::read_to_string("input").expect("Failed to read in file");

    let character_bytes = input.into_bytes();


    //let mut bytes_iter = character_bytes.windows(4);

    /*
    for byte in bytes_iter {
        let mut check = 0;
        'outer: for i in 0..4 {
             for j in 0..4 {
                if i != j && byte[i] == byte[j] {
                    check = 1;
                    break 'outer 
                }
            }
        }
        if check == 0 {
            let next = bytes_iter.next().unwrap();
            let first = next[0];
        }
    }
    */


    let mut next_check = 0;
    let mut count = 3;
    for byte in character_bytes.windows(4) {
        count += 1;
        println!("{}", count);
        println!("{:?}", byte);
        if next_check != 0 {
            break;    
        }
        let mut check = 0;
        'outer: for i in 0..4 {
             for j in 0..4 {
                if i != j && byte[i] == byte[j] {
                    check = 1;
                    println!("{} is equal to {}", byte[i], byte[j]);
                    break 'outer 
                }
            }
        }
        if check == 0 {
            next_check = 1;
        }
    }


    println!("{}", count-1)
}
