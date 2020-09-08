
struct Arg {
    1: i32 a
    2: i32 b
}

service Calculate {
    i32 add(1: i32 a, 2: i32 b)
}