

fn main() {
    //let input = std::fs::read_to_string("test").expect("Failed to read in file");
    let input = std::fs::read_to_string("input").expect("Failed to read in file");

    let mut lines: Vec<&str> = Vec::new();
    for line in input.lines() {
       lines.push(line); 
    }
     
    let mut sum = 0;

    for line in &lines {
        let split: Vec<&str> = line.split(",").collect();
        
        let nums1: Vec<&str> = split[0].split("-").collect();
        let nums2: Vec<&str> = split[1].split("-").collect();

        let p1_1 = nums1[0].parse::<usize>().unwrap();
        let p1_2 = nums1[1].parse::<usize>().unwrap();
        let p2_1 = nums2[0].parse::<usize>().unwrap();
        let p2_2 = nums2[1].parse::<usize>().unwrap();

        /* PART 1
        if p1_1 >= p2_1 && p1_2 <= p2_2 || p1_1 <= p2_1 && p1_2 >= p2_2 {
           sum += 1; 
        }
        */

        //2-4, 3-5
        //3-5, 2-4
        // PART 2
        if p1_2 >= p2_1 && p1_2 <= p2_2 || p2_2 >= p1_1 && p2_2 <= p1_2  {
           sum += 1; 
        }


        //println!("{:?}", sum);
    }



    println!("{:?}", sum);

}
