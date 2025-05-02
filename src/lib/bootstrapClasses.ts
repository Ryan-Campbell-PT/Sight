//this is a simple work around to allow bootstrap classes in custom css objects
//without having to add a new library or anything
export const bootstrap = {
    nutritionLabel: {
        //extended/macroRow defines the actual row that contains all the information 
        macroRow: "fs-6 d-flex border-top border-dark border-1 row justify-content-between", 
        extendedMacroRow: "ps-3 border-top border-1 border-dark row justify-content-between",
        //macroNameValue defines the left side of the macro information, containing the name (carbs) and the value (28g)
        macroNameValue: "col-10 align-self-start",
        //percentage defines the percentage calculaged from macroValue, aligned to the end of the container
        percentage: "col-2 align-self-end text-end",
    }
}