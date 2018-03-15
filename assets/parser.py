import csv
import json

heroes = {}

def parse_name(name):
    name = name.replace(" ", "_")
    name = name.replace("!", "")
    name = name.replace("(", "")
    name = name.replace(")", "")
    return name.lower()

with open("base_stats.csv", "r") as f:
    reader = csv.reader(f)
    for row in reader:
        name = parse_name(row[0])
        heroes[name] = {
            "name": row[0],
            "base": {
                "hp": row[3],
                "atk": row[4],
                "spd": row[5],
                "def": row[6],
                "res": row[7]
            }
        }

with open("growth_points.csv", "r") as f:
    reader = csv.reader(f)
    for row in reader:
        name = parse_name(row[0])
        heroes[name]["growth_points"] = {
            "hp": row[3],
            "atk": row[4],
            "spd": row[5],
            "def": row[6],
            "res": row[7]
        }

for key in heroes:
    with open("heroes/" + key + ".json", "w") as f:
        json.dump(heroes[key], f, indent=2)
