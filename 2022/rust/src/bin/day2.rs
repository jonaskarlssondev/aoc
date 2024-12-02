fn main() {

}

fn to_char_value(s: &str) -> u32 {
    return match s {
        "A" => 0,
        "B" => 2,
        "C" => 1,
        "X" => 1,
        "Y" => 2,
        "Z" => 2
    }
}

fn to_result_value(s: &str) -> u32 {
    return 0
}

