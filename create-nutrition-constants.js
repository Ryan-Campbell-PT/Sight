// generate this file by the command:
// node create-nutrition-constants.js

import fs from "fs";
const data = JSON.parse(fs.readFileSync("nutrition-constants.json", "utf8"));

// --- Generate TS ---
let ts = `// AUTO-GENERATED — DO NOT EDIT\n`;
ts += `export enum NutritionValues {\n`;
for (const n of data) ts += `  ${n.key} = "${n.key}",\n`;
ts += "}\n\n";
ts += `export const NutritionMap = {\n`;
for (const n of data)
    ts += `  [NutritionValues.${n.key}]: { id: ${n.id}, unit: "${n.unit}", dailyValue: ${n.dailyValue ?? "null"}, dbName: ${n.dbName ? `"${n.dbName}"` : null}, displayName:  "${n.displayName}"},\n`;
ts += "};\n\n";
ts += `export function getNutritionId(key: NutritionValues): number {\n  return NutritionMap[key].id;\n}\n`;

fs.writeFileSync("./src/NutritionConstants.ts", ts);

// --- Generate Crystal ---
let cr = "# AUTO-GENERATED — DO NOT EDIT\n";
cr += "module NutritionValues\n";
for (const n of data) cr += `  ${n.key} = "${n.key}"\n`;
cr += "end\n\n";
cr += "NUTRITION_MAP = {\n";
for (const n of data)
    cr += `  NutritionValues::${n.key} => {id: ${n.id}, unit: "${n.unit}", daily_value: ${n.dailyValue ?? "nil"}, dbName: ${n.dbName ? `"${n.dbName}"` : "nil"}, display_name: "${n.displayName}"},\n`;
cr += "}\n\n";
cr += "def get_nutrition_id(key)\n  NUTRITION_MAP[key][:id]\nend\n";

fs.writeFileSync("./backend/crystal-backend/src/models/nutrition-constants.cr", cr);

console.log("-- Generated NutritionConstants.ts and nutrition-constants.cr --");
