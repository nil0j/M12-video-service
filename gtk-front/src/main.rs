use gtk::prelude::*;
use gtk::{Align, glib};
use gtk::{Application, ApplicationWindow, Box, Image, Label, Orientation, SearchEntry};

pub fn build_ui(app: &Application) {
    // Create the main window
    let window = ApplicationWindow::builder()
        .application(app)
        .title("Rust GTK Media App")
        .default_width(800)
        .default_height(600)
        .build();

    // Main vertical box container
    let main_box = Box::new(Orientation::Vertical, 0);
    window.set_child(Some(&main_box));

    // Top bar with image, text, and search
    let top_bar = Box::new(Orientation::Horizontal, 10);
    top_bar.set_valign(Align::Start);
    top_bar.set_margin_bottom(10);

    // App logo/image
    let logo = Image::from_file("./assets/giraffe.png");
    logo.set_margin_end(10);
    top_bar.append(&logo);

    // App title
    let title = Label::new(Some("Media Player"));
    title.set_halign(Align::Start);
    title.set_hexpand(true);
    top_bar.append(&title);

    // Search bar
    let search_entry = SearchEntry::new();
    search_entry.set_placeholder_text(Some("Search..."));
    search_entry.set_width_chars(30);
    search_entry.set_margin_start(10);
    top_bar.append(&search_entry);

    main_box.append(&top_bar);

    // Video player at the bottom
    let video = gtk::Video::new();
    video.set_vexpand(true);

    // Set video from URL (GStreamer will handle the streaming)
    if let Some(file) = Some(gtk::gio::File::for_uri(
        "https://giraffe.niliara.net/api/video/source/1",
    )) {
        video.set_file(Some(&file));
    }

    main_box.append(&video);

    window.show();
}

fn main() -> glib::ExitCode {
    let app = Application::builder()
        .application_id("org.example.mediaplayer")
        .build();

    app.connect_activate(build_ui);

    app.run()
}
