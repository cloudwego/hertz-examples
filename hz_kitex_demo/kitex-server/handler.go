package main

import (
	"context"
	"hertz-examples/hz_demo/kitex-server/kitex_gen/student/management"
)

// StudentManagementImpl implements the last service interface defined in the IDL.
type StudentManagementImpl struct{}

type StudentInfo struct {
	Num    string
	Name   string
	Gender string
}

var StudentData = make(map[string]StudentInfo, 5)

// QueryStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) QueryStudent(ctx context.Context, req *management.QueryStudentRequest) (resp *management.QueryStudentResponse, err error) {
	// TODO: Your code here...
	stu, exist := StudentData[req.Num]
	if !exist {
		return &management.QueryStudentResponse{
			Exist: false,
		}, nil
	}

	resp = &management.QueryStudentResponse{
		Exist:  true,
		Num:    stu.Num,
		Name:   stu.Name,
		Gender: stu.Gender,
	}

	return resp, nil
}

// InsertStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) InsertStudent(ctx context.Context, req *management.InsertStudentRequest) (resp *management.InsertStudentResponse, err error) {
	// TODO: Your code here...
	_, exist := StudentData[req.Num]
	if exist {
		return &management.InsertStudentResponse{
			Ok:  false,
			Msg: "the num has exists",
		}, nil
	}

	StudentData[req.Num] = StudentInfo{
		Num:    req.Num,
		Name:   req.Name,
		Gender: req.Gender,
	}

	return &management.InsertStudentResponse{
		Ok: true,
	}, nil
}
