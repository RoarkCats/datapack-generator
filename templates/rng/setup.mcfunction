scoreboard objectives add %namespace%.rng dummy

scoreboard players set in %namespace%.rng -100
scoreboard players set in1 %namespace%.rng 100

scoreboard players set #lcg_constant %namespace%.rng 1103515245
execute unless score #lcg %namespace%.rng matches ..0 unless score #lcg %namespace%.rng matches 1.. run function %namespace%:rng/uuid_reset