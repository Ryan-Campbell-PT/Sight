export function roundToDecimal(num: number, decimalPlace: number): number {
	const factor = 10 ** decimalPlace;
	return Math.round(num * factor) / factor;
}