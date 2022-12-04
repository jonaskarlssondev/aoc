fn main() {
    let input = include_str!("./day1.input");

    let lines = input.split("\n\n");
   
    let mut inventories: Vec<u32> = lines
        .map(|line| line.split("\n")
             .flat_map(|n| n.parse::<u32>())
             .sum())
        .collect();

    inventories.sort_by(|a, b| b.cmp(a));

    println!("1: {:?}", inventories[0]);
    println!("2: {:?}", inventories.iter().take(3).sum::<u32>());
}
