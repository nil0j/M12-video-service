use gtk::glib;
use gtk::prelude::*;
use gtk::{Application, ApplicationWindow};

mod components;
mod models;
mod scenes;

pub fn build_ui(app: &Application) {
    let window = ApplicationWindow::builder()
        .application(app)
        .title("Jirafeitor")
        .default_width(800)
        .default_height(600)
        .build();

    let main_box = gtk::Box::new(gtk::Orientation::Vertical, 0);
    window.set_child(Some(&main_box));

    main_box.append(&components::topbar::get());
    main_box.append(&scenes::get());

    window.show();
}

fn main() -> glib::ExitCode {
    let app = Application::builder()
        .application_id("net.niliara.jirafeitor")
        .build();

    app.connect_activate(build_ui);

    app.run()
}
