use std::{net::TcpStream, io::{Write, BufReader, BufRead}};
mod key;
use key::*;

fn main() {
    let mut stream = TcpStream::connect("localhost:8080").unwrap();

    // Get keybinds
    let mut binds_raw = String::new();
    let mut reader = BufReader::new(&stream);
    reader.read_line(&mut binds_raw).unwrap();
    binds_raw.pop(); // Remove newline
    for item in binds_raw.split(";") {
        if item.len() == 0 {
            continue;
        }

        let vals: Vec<&str> = item.split(": ").collect();
        let key = Key::from_string(&vals[0].to_string());
        let action = vals[1].to_string();
        println!("{} -> {}", key.to_string(), action);
    }

    // Send key
    write!(stream, "W:true\n").unwrap();
    stream.flush().unwrap();
}
