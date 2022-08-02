use chrono::Local;
use console::{Term, style};
use std::{thread, time};
use terminal_size::{terminal_size, Height, Width};

const INTERVAL: u64 = 950;
const LEN_STRFTIME: usize = 19;

fn main() {
    let size = terminal_size().unwrap();
    let (Width(width), Height(height)) = size;

    let term = Term::stdout();
    term.set_title("crock");
    term.hide_cursor().unwrap();
    term.clear_screen().unwrap();

    let interval = time::Duration::from_millis(INTERVAL);
    let mut previous = Local::now().format("%Y/%m/%d %H:%M:%S").to_string();

    let spaces = String::from(" ").repeat(width as usize - LEN_STRFTIME);

    loop {
        let now = Local::now();
        let fmt = now.format("%Y/%m/%d %H:%M:%S").to_string();
        let line_to_show = format!("{}{}", spaces, fmt);

        if previous == fmt {
            continue;
        }

        term.clear_last_lines(2).unwrap();

        println!("{}", style(&line_to_show).dim());
        render_clock(&width, &height);

        previous = fmt;

        thread::sleep(interval);
    }
}

fn render_clock(width: &u16, height: &u16) {







    
}
