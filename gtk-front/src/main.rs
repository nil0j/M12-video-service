use gtk::prelude::*;
use gtk::{Align, glib};
use gtk::{Application, ApplicationWindow};

pub fn build_ui(app: &Application) {
    // Create the main window
    let window = ApplicationWindow::builder()
        .application(app)
        .title("Jirafeitor")
        .default_width(800)
        .default_height(600)
        .build();

    // Main vertical box container
    let main_box = gtk::Box::new(gtk::Orientation::Vertical, 0);
    window.set_child(Some(&main_box));

    main_box.append(&top_bar());
    main_box.append(&video_scene());

    window.show();
}

fn top_bar() -> gtk::Box {
    // Top bar with image, text, and search
    let top_bar = gtk::Box::new(gtk::Orientation::Horizontal, 10);
    top_bar.set_valign(Align::Start);
    top_bar.set_margin_top(5);

    // App logo/image
    let logo = gtk::Image::from_file("./assets/giraffe.png");
    logo.set_margin_end(10);
    logo.set_margin_start(20);
    logo.set_pixel_size(50);
    top_bar.append(&logo);

    // App title
    let title = gtk::Label::new(Some("Jirafeitor"));
    title.set_halign(Align::Start);
    title.set_hexpand(true);
    top_bar.append(&title);

    // Search bar
    let search_entry = gtk::SearchEntry::new();
    search_entry.set_placeholder_text(Some("Search..."));
    search_entry.set_max_width_chars(40);
    search_entry.set_margin_start(20);
    search_entry.set_margin_end(20);
    search_entry.set_margin_top(10);
    search_entry.set_margin_bottom(10);
    top_bar.append(&search_entry);

    return top_bar.upcast();
}

fn video_scene() -> gtk::Widget {
    // Video player at the bottom
    let video = gtk::Video::new();
    video.set_vexpand(true);

    // Set video from URL (GStreamer will handle the streaming)
    if let Some(file) = Some(gtk::gio::File::for_uri(
        "https://giraffe.niliara.net/api/video/source/1",
    )) {
        video.set_file(Some(&file));
    }

    return video.upcast();
}

fn main() -> glib::ExitCode {
    let app = Application::builder()
        .application_id("net.niliara.jirafeitor")
        .build();

    app.connect_activate(build_ui);

    app.run()
}
