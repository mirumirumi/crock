use chrono::Local;
use console::Term;
use std::{thread, time};
use terminal_size::{terminal_size, Height, Width};

const INTERVAL: u64 = 444;

fn main() {
    let term = Term::stdout();
    term.set_title("crock");
    term.hide_cursor().unwrap();
    term.clear_screen().unwrap();

    let interval = time::Duration::from_millis(INTERVAL);
    let mut previous = Local::now().format("%Y/%m/%d %H:%M:%S").to_string();

    let size = terminal_size().unwrap();
    let (Width(width), Height(height)) = size;

    loop {
        let now = Local::now();
        let fmt = now.format("%Y/%m/%d %H:%M:%S").to_string();

        if previous == fmt {
            continue;
        }

        term.clear_last_lines(2).unwrap();

        println!("{}", fmt);
        previous = fmt;

        thread::sleep(interval);
    }
}
