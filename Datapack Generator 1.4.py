import os
import time

# Rng import, mostly just writing large strings to a few files with the namespace filled in
def import_rng(ver):

    print("  Importing rng...")
    os.makedirs(f"./{pack_name}/data/{pack_namespace}/functions/rng/")
    
    with open(f"{pack_name}/data/{pack_namespace}/functions/rng/lcg.mcfunction", "w") as f:
        
        f.write(f"""# LCG Seed implementation
#
# x_(n+1) = x_(n)*a + c
#
# a = 1103515245, c = 12345

scoreboard players operation #lcg {pack_namespace}.rng *= #lcg_constant {pack_namespace}.rng
scoreboard players add #lcg {pack_namespace}.rng 12345
scoreboard players operation out {pack_namespace}.rng = #lcg {pack_namespace}.rng""")
    
    with open(f"{pack_name}/data/{pack_namespace}/functions/rng/next_int_lcg.mcfunction", "w") as f:
        
        f.write(f"""###
#    public int nextInt(int bound) {{
#        if (bound <= 0)
#            throw new IllegalArgumentException(BadBound);
#
#        int r = next(31);
#        int m = bound - 1;
#        if ((bound & m) == 0)  // i.e., bound is a power of 2
#            r = (int)((bound * (long)r) >> 31);
#        else {{
#            for (int u = r; u - (r = u % bound) + m < 0; u = next(31));
#        }}
#        return r;
#    }}

function {pack_namespace}:rng/lcg

scoreboard players operation #temp {pack_namespace}.rng = out {pack_namespace}.rng
scoreboard players operation out {pack_namespace}.rng %= #range {pack_namespace}.rng
scoreboard players operation #temp {pack_namespace}.rng -= out {pack_namespace}.rng
scoreboard players operation #temp {pack_namespace}.rng += #m1 {pack_namespace}.rng
execute if score #temp {pack_namespace}.rng matches ..-1 run function {pack_namespace}:rng/next_int_lcg""")
    
    with open(f"{pack_name}/data/{pack_namespace}/functions/rng/range_lcg.mcfunction", "w") as f:
        
        f.write(f"""scoreboard players add in1 {pack_namespace}.rng 1
scoreboard players operation #range {pack_namespace}.rng = in1 {pack_namespace}.rng
scoreboard players operation #range {pack_namespace}.rng -= in {pack_namespace}.rng

scoreboard players operation #m1 {pack_namespace}.rng = #range {pack_namespace}.rng
scoreboard players remove #m1 {pack_namespace}.rng 1
function {pack_namespace}:rng/next_int_lcg
scoreboard players operation out {pack_namespace}.rng += in {pack_namespace}.rng

scoreboard players reset #m1 {pack_namespace}.rng
scoreboard players remove in1 {pack_namespace}.rng 1""")
    
    with open(f"{pack_name}/data/{pack_namespace}/functions/rng/setup.mcfunction", "w") as f:
        
        f.write(f"""scoreboard objectives add {pack_namespace}.rng dummy

scoreboard players set in {pack_namespace}.rng -100
scoreboard players set in1 {pack_namespace}.rng 100

scoreboard players set #lcg_constant {pack_namespace}.rng 1103515245
execute unless score #lcg {pack_namespace}.rng matches ..0 unless score #lcg {pack_namespace}.rng matches 1.. run function {pack_namespace}:rng/uuid_reset""")
    
    with open(f"{pack_name}/data/{pack_namespace}/functions/rng/uuid_reset.mcfunction", "w") as f:
        
        # UUID data storage changed in 1.16+
        insert = "UUID[0] 1"
        if ver < 6 :
            insert = "UUIDMost 0.00000000023283064365386962890625"

        f.write(f"""summon area_effect_cloud ~ ~ ~ {{Tags:["get_uuid"]}}
execute store result score #lcg {pack_namespace}.rng run data get entity @e[tag=get_uuid,limit=1] {insert}
kill @e[tag=get_uuid]""")

    with open(f"{pack_name}/data/{pack_namespace}/functions/load.mcfunction", "r") as f:

        f_remp = f.read()

        with open(f"{pack_name}/data/{pack_namespace}/functions/load.mcfunction", "w") as f:

            f.write(f"{f_remp}\n\nfunction {pack_namespace}:rng/setup")      












print("\n  -- Thank you for using Datapack Generator by RoarkCats --\n")
pack_name = input("  Datapack Name: ")
pack_namespace = input("  Datapack Namespace: ")
pack_author = input("  Datapack Author: ")
pack_ver = input("  Datapack Version: ")
pack_rng = input("  Import Rng? (y/n) ")

try :
    pack_ver = int(pack_ver)
except :
    pack_ver = 10

proc_time_start = time.process_time()

print("\n  Generating folders...")
os.makedirs(f"./{pack_name}/data/minecraft/tags/functions/")
os.makedirs(f"./{pack_name}/data/{pack_namespace}/functions/")

print("  Generating files...")
with open(f"{pack_name}/pack.mcmeta", "w") as f:

    f.write(f"""{{
    \"pack\": {{
        \"pack_format\": {pack_ver},
        \"description\": \"{pack_name} by {pack_author}\"
    }}
}}""")

with open(f"{pack_name}/data/minecraft/tags/functions/tick.json", "w") as f:

    f.write(f"""{{
    \"values\":[
        \"{pack_namespace}:main\"
    ]
}}""")

with open(f"{pack_name}/data/minecraft/tags/functions/load.json", "w") as f:

    f.write(f"""{{
    \"values\":[
        \"{pack_namespace}:load\"
    ]
}}""")

with open(f"{pack_name}/data/{pack_namespace}/functions/main.mcfunction", "w") as f:

    f.write("")

with open(f"{pack_name}/data/{pack_namespace}/functions/load.mcfunction", "w") as f:

    f.write(f"tellraw @a {{\"text\":\"Thank you for downloading {pack_name} by {pack_author}!\",\"color\":\"gold\"}}")

if pack_rng == "y":
    import_rng(pack_ver)

proc_time_end = time.process_time()
print(f"  Done! ({(proc_time_end - proc_time_start)*1000} ms)")
input()