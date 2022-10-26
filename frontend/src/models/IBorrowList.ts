import { EquipmentInterface } from "./IEquipment";
import { EmployeeInterface } from "./IEmployee";
import { UsersInterface } from "./IUser";

export interface BorrowListInterface{
    ID?: number;

    EquipmentID?: number;
    Equipment?: EquipmentInterface;

    MemberID?: number;
    Member?: UsersInterface;

    EmployeeID?: number;
    Employee?: EmployeeInterface;

    Amount?: number;
    BorrowTime?: Date | null;
}
