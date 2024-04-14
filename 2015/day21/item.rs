#[derive(Clone)]
pub struct Item {
    pub name: String,
    pub cost: i32,
    pub damage: i32,
    pub defense: i32,
}

impl Item {
    pub fn new(name: &str, cost: i32, damage: i32, defense: i32) -> Self {
        Item {
            name: name.to_string(),
            cost,
            damage,
            defense,
        }
    }
}
