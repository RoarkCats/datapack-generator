# LCG Seed implementation
#
# x_(n+1) = x_(n)*a + c
#
# a = 1103515245, c = 12345

scoreboard players operation #lcg %namespace%.rng *= #lcg_constant %namespace%.rng
scoreboard players add #lcg %namespace%.rng 12345
scoreboard players operation out %namespace%.rng = #lcg %namespace%.rng