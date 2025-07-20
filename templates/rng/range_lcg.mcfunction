scoreboard players add in1 %namespace%.rng 1
scoreboard players operation #range %namespace%.rng = in1 %namespace%.rng
scoreboard players operation #range %namespace%.rng -= in %namespace%.rng

scoreboard players operation #m1 %namespace%.rng = #range %namespace%.rng
scoreboard players remove #m1 %namespace%.rng 1
function %namespace%:rng/next_int_lcg
scoreboard players operation out %namespace%.rng += in %namespace%.rng

scoreboard players reset #m1 %namespace%.rng
scoreboard players remove in1 %namespace%.rng 1