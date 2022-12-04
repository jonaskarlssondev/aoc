fn main() {
    let input = include_str!("./day1.input");

    let lines = input.split("\n\n");
   
    let max: Option<u32> = lines
        .map(|line| line.split("\n")
             .flat_map(|n| n.parse::<u32>())
             .sum())
        .max();

    println!("1: {:?}", max)
}
