# gossip-glomers
My solutions to the Gossip Glomers distributed systems challenge - https://fly.io/dist-sys/

To avoid spoilers, I will just add some parts of the solutions here. See the directories or links for complete solutions.


## Challenge 1 - echo

This is just to set up the system and dependencies. They already give us the solution.

## Challenge 2 - unique ID generator

From the challenge:  "IDs may be of any type--strings, booleans, integers, floats, arrays, etc."
Though IDs can be of any type, let us stick with integer since they are easier to reason about (and can be converted to string or array if required). Booleans can be ruled out cuz duh, and floats are black magic to a computer so let's stay away as any sane person would.

[Solution README](maelstrom-unique-id/README.md)

[Solution code](maelstrom-unique-id/main.go)

