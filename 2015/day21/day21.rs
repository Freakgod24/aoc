#![allow(dead_code)]
#![allow(unused_variables)]
mod boss;
mod item;
mod player;

use crate::boss::Boss;
use crate::item::Item;
use crate::player::Player;
use genawaiter::{sync::gen, yield_};
use itertools::iproduct;
use std::future::Future;

fn main() {
    part1();
    part2();
}

fn part1() {
    let mut gold_spent = std::i32::MAX;
    for (weapon, armor, rings) in create_all_configurations() {
        let mut boss = Boss::new(100, 8, 2);
        let mut player = Player::new(100, weapon, armor, rings);

        while boss.hitpoints > 0 && player.hitpoints > 0 {
            boss.apply_damage(player.damage);
            if boss.hitpoints > 0 {
                player.apply_damage(boss.damage);
            }
        }

        if player.hitpoints > 0 {
            gold_spent = gold_spent.min(player.value);
        }
    }

    println!("{}", gold_spent);
}

fn part2() {
    let mut gold_spent = 0;
    for (weapon, armor, rings) in create_all_configurations() {
        let mut boss = Boss::new(100, 8, 2);
        let mut player = Player::new(100, weapon, armor, rings);

        while boss.hitpoints > 0 && player.hitpoints > 0 {
            boss.apply_damage(player.damage);
            if boss.hitpoints > 0 {
                player.apply_damage(boss.damage);
            }
        }

        if player.hitpoints == 0 {
            gold_spent = gold_spent.max(player.value);
        }
    }

    println!("{}", gold_spent);
}

fn create_all_configurations(
) -> genawaiter::sync::Gen<(Item, Item, [Item; 2]), (), impl Future<Output = ()>> {
    let weapons = [
        Item::new("Dagger", 8, 4, 0),
        Item::new("Shortsword", 10, 5, 0),
        Item::new("Warhammer", 25, 6, 0),
        Item::new("Longsword", 40, 7, 0),
        Item::new("Greataxe", 74, 8, 0),
    ];
    let armors = [
        Item::new("None", 0, 0, 0),
        Item::new("Leather", 13, 0, 1),
        Item::new("Chainmail", 31, 0, 2),
        Item::new("Splintmail", 53, 0, 3),
        Item::new("Bandedmail", 75, 0, 4),
        Item::new("Platemail", 102, 0, 5),
    ];
    let rings = [
        Item::new("Damage +0", 0, 0, 0),
        Item::new("Damage +1", 25, 1, 0),
        Item::new("Damage +2", 50, 2, 0),
        Item::new("Damage +3", 100, 3, 0),
        Item::new("Defense +0", 0, 0, 0),
        Item::new("Defense +1", 20, 0, 1),
        Item::new("Defense +2", 40, 0, 2),
        Item::new("Defense +3", 80, 0, 3),
    ];

    // For this puzzle, have decided to try the genawaiter crate. This allows moving alot of nested code
    // outside of this function while still going through the list only once.
    gen!({
        for (weapon, armor, ring1, ring2) in iproduct!(&weapons, &armors, &rings, &rings) {
            if ring1.name != ring2.name {
                yield_!((
                    weapon.clone(),
                    armor.clone(),
                    [ring1.clone(), ring2.clone()]
                ));
            }
        }
    })
}
