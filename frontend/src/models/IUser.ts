export interface UsersInterface {
    ID:         number,

    // PrefixID:  number,
    Prefix:     string,

    FirstName: string,
    LastName:  string,
    Email:      string,
    Password:   string,
    Address:    string,
    Birthday:   Date | null,

    // GenderID:  number,
    Gender:     string,
    
    PersonalID: string,
    Mobile:      string;
}