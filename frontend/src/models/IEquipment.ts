import { EmployeeInterface } from "./IEmployee";

export interface EquipmentInterface {
    ID:             number;
    Name:           string;
    Amount:         number;
    Time:           Date | null;

    CatagoryID:    number;
    // Catagory:       CatagoryInteraface;
    UnitID:        number;
    // Unit:           UnitInterface;
    
    EmployeeID:    number;
    Employee:       EmployeeInterface;
}