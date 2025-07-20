###
#    public int nextInt(int bound) {
#        if (bound <= 0)
#            throw new IllegalArgumentException(BadBound);
#
#        int r = next(31);
#        int m = bound - 1;
#        if ((bound & m) == 0)  // i.e., bound is a power of 2
#            r = (int)((bound * (long)r) >> 31);
#        else {
#            for (int u = r; u - (r = u % bound) + m < 0; u = next(31));
#        }
#        return r;
#    }

function %namespace%:rng/lcg

scoreboard players operation #temp %namespace%.rng = out %namespace%.rng
scoreboard players operation out %namespace%.rng %= #range %namespace%.rng
scoreboard players operation #temp %namespace%.rng -= out %namespace%.rng
scoreboard players operation #temp %namespace%.rng += #m1 %namespace%.rng
execute if score #temp %namespace%.rng matches ..-1 run function %namespace%:rng/next_int_lcg