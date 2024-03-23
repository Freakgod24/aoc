use std::collections::HashMap;

pub struct Graph {
    vertices: HashMap<String, Vec<(String, u32)>>,
}

impl Graph {
    fn new() -> Self {
        Graph {
            vertices: HashMap::new(),
        }
    }

    fn add_vertex(&mut self, vertex: &str) {
        if !self.vertices.contains_key(vertex) {
            self.vertices.insert(vertex.to_string(), Vec::new());
        }
    }

    fn add_edge(&mut self, source: &str, destination: &str, distance: u32) {
        self.vertices
            .entry(source.to_string())
            .or_insert(Vec::new())
            .push((destination.to_string(), distance));
        self.vertices
            .entry(destination.to_string())
            .or_insert(Vec::new())
            .push((source.to_string(), distance));
    }

    fn print(&self) {
        for (vertex, neighbors) in &self.vertices {
            print!("Vertex {}: ", vertex);
            for (neighbor, distance) in neighbors {
                print!("({}, {}), ", neighbor, distance);
            }
            println!();
        }
    }
}
