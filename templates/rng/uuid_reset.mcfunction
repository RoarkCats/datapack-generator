summon area_effect_cloud ~ ~ ~ {Tags:["get_uuid"]}
execute store result score #lcg %namespace%.rng run data get entity @e[tag=get_uuid,limit=1] UUID[0] 1
kill @e[tag=get_uuid]