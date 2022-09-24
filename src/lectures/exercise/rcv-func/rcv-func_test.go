package main

import "testing"

func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

}

func TestHealth(t *testing.T) {
	player := newPlayer()
	player.addHealth(999)

	if player.health > player.maxHealth {
		t.Fatalf("Health went over limit: %v, want: %v", player.health, player.maxHealth)
	}
	player.applyDamage(player.maxHealth + 1)
	if player.health < 0 {
		t.Fatalf("Health: %v. Minimum: 0", player.health)
	}

	if player.health > player.maxHealth {
		t.Fatalf("Health: %v. Maximum: %v", player.health, player.maxHealth)
	}

}
