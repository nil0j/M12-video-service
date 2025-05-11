use gtk::prelude::*;

pub mod video_scene;

pub fn get() -> gtk::Widget {
    let scene = gtk::Stack::new();
    scene.add_child(&video_scene::get());
    return scene.upcast();
}
