use crate::item::Item;

pub struct Player {
    pub hitpoints: i32,
    pub weapon: Item,
    pub armor: Item,
    pub rings: [Item; 2],
    pub damage: i32,
    pub defense: i32,
    pub value: i32,
}

impl Player {
    pub fn new(hitpoints: i32, weapon: Item, armor: Item, rings: [Item; 2]) -> Self {
        Self {
            // Base Stats
            hitpoints,

            // Items
            weapon: weapon.clone(),
            armor: armor.clone(),
            rings: rings.clone(),

            // Rating
            damage: weapon.damage + rings[0].damage + rings[1].damage,
            defense: armor.defense + rings[0].defense + rings[1].defense,
            value: weapon.cost + armor.cost + rings[0].cost + rings[1].cost,
        }
    }

    pub fn apply_damage(&mut self, damage: i32) {
        let hit = (damage - self.defense).max(1);
        self.hitpoints = (self.hitpoints - hit).max(0);
    }
}
