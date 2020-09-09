namespace go qlnv

exception OverlapException{
    1: string value
}

exception NotFoundException{
    1: string value
}

exception AccessException{
    1: string type
    2: string value
}

struct User {
    1: string userId
    2: string name
    3: i64 dateOfBirth
    4: string address
    5: string unitId
}


service UserService{
    void createUser(1: User u) throws (1:OverlapException err)
    User readUser(1: string userId) throws (1:NotFoundException err)
    void updateUser(1: User u) throws (1:NotFoundException err)
    void deleteUser(1:string userId) throws(1:NotFoundException notFound,2:AccessException accessEx)
    Unit getUnitUser(1:string userId) throws(1:NotFoundException err)
    list<User> getUserSortedByPage(1:i32 numOfPages,2: i32 sizeOfpage,3:string sortType) throws (1:NotFoundException err)
}

struct Unit {
    1: string unitId
    2: string name
    3: string address
}

service UnitService{
    void createUnit(1: Unit u) throws (1:OverlapException err)
    Unit readUnit(1: string unitId) throws (1:NotFoundException err)
    void updateUnit(1: Unit u) throws (1:NotFoundException err)
    void deleteUnit(1:string unitId) throws(1:NotFoundException notFound,2:AccessException accessEx)
    list<User> getAllMemberOfUnit(1:string unitId) throws(1:NotFoundException err)
    list<User> getMembersByPage(1:string unitId,2:i32 numOfPage,3: i32 sizeOfPage) throws (1:NotFoundException err)
}
