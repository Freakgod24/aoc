pub struct Boss {
    pub hitpoints: i32,
    pub damage: i32,
    pub defense: i32,
}

impl Boss {
    pub fn new(hitpoints: i32, damage: i32, defense: i32) -> Self {
        Self {
            hitpoints,
            damage,
            defense,
        }
    }

    pub fn apply_damage(&mut self, damage: i32) {
        let hit = (damage - self.defense).max(1);
        self.hitpoints = (self.hitpoints - hit).max(0);
    }
}
