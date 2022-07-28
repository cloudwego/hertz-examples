namespace go student.management

struct QueryStudentRequest {
    1: string Num;
}

struct QueryStudentResponse {
    1: bool   Exist;
    2: string Num;
    3: string Name;
    4: string Gender;
}

struct InsertStudentRequest {
    1: string Num;
    2: string Name;
    3: string Gender;
}

struct InsertStudentResponse {
    1: bool Ok;
    2: string Msg;
}

service StudentManagement {
    QueryStudentResponse queryStudent(1: QueryStudentRequest req);
    InsertStudentResponse insertStudent(1: InsertStudentRequest req);
}
