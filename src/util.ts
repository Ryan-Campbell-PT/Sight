export function roundToDecimal(num: number, decimalPlace: number) {
    const decimal = Math.pow(10, decimalPlace)
    return Math.round(num * decimal) / decimal
}