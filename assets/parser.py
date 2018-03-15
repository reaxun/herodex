import csv
import json

heroes = {}

def parse_name(name):
    name = name.replace(" ", "_")
    name = name.replace("!", "")
    name = name.replace("(", "")
    name = name.replace(")", "")
    return name.lower()

with open("csv/heroes.csv", "r") as f:
    reader = csv.reader(f)
    for row in reader:
        name = parse_name(row[0])
        rarities = [int(r) for r in ",".join(row[5] + row[6] + row[7] + row[8] + row[9] + row[10]).split(",") if r != ""]
        heroes[name] = {
            "name": row[0],
            "title": row[1],
            "origin": row[2],
            "rarities": rarities
        }
        
with open("csv/base_stats.csv", "r") as f:
    reader = csv.reader(f)
    for row in reader:
        name = parse_name(row[0])
        heroes[name]["base"] = {
            "hp": int(row[3]),
            "atk": int(row[4]),
            "spd": int(row[5]),
            "def": int(row[6]),
            "res": int(row[7])
        }

with open("csv/growth_points.csv", "r") as f:
    reader = csv.reader(f)
    for row in reader:
        name = parse_name(row[0])
        heroes[name]["growth_points"] = {
            "hp": int(row[3]),
            "atk": int(row[4]),
            "spd": int(row[5]),
            "def": int(row[6]),
            "res": int(row[7])
        }

for key in heroes:
    with open("heroes/" + key + ".json", "w") as f:
        json.dump(heroes[key], f, indent=2)
