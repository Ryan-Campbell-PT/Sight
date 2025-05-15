import { NutritionLabelContent } from "./NutritionData";

export function roundToDecimal(num: number, decimalPlace: number): number {
	const factor = 10 ** decimalPlace;
	return Math.round(num * factor) / factor;
}

export function formatDateToYYYYMMDD(date: Date): string {
	const year = date.getFullYear();
	const month = (date.getMonth() + 1).toString().padStart(2, "0");
	const day = date.getDate().toString().padStart(2, "0");
	return `${year}-${month}-${day}`;
}

export function getNutritionValueFromName(macroId: number, nutritionMap: Map<number, number>): number {
	// const macroValue = NutritionLabelContent.find(m => m.macro_name === macro);
	// return macroValue ? nutritionMap.get(macroValue.id) ?? -1 : -1;

	// return nutritionMap.get(macroId) ?? -1;
	return -1
}