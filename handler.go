package main

import (
	"context"
	demo "github.com/SchrodingerwithCat/rpcsvr/kitex_gen/demo"

	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"github.com/SchrodingerwithCat/rpcsvr/model"
)

/////////////////////////////////////////////////////////

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	// add global database
	db  *gorm.DB
	err error
}

func NewStuService() *StudentServiceImpl {
	return &StudentServiceImpl{
		db:  nil,
		err: nil,
	}
}

func (s *StudentServiceImpl) InitStuService(name_db string) {
	// 返回值为  void 类型
	// add database
	s.db, s.err = gorm.Open(sqlite.Open(name_db), &gorm.Config{})

	if s.err != nil {
		panic("failed to connect to the database!")
	}

	// create the table
	s.err = s.db.Migrator().CreateTable(&model.Student{})

	if s.err != nil {
		panic("create the database failed")
	}
}

/////////////////////////////////////////////////////////

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	// 先查询数据库是否有该学生信息，如果没有将 Student 信息插入到数据库

	// 查询数据库
	var stuRes *model.Student
	id := student.Id
	query_result := s.db.Table("students").First(&stuRes, id)

	if errors.Is(query_result.Error, gorm.ErrRecordNotFound) {
		// not found, insert the info to the database
		s.db.Table("students").Create(student2Model(student))

		resp.Success = true
		resp.Message = "insert the info successfully!"
		return
	}

	resp.Success = false
	resp.Message = "the stu info has exited!"
	err = nil

	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...
	var stuRes *model.Student
	id := req.Id
	query_result := s.db.Table("students").First(&stuRes, id)

	if errors.Is(query_result.Error, gorm.ErrRecordNotFound) {
		// not found
		return
	}

	// found the stu info
	resp = model2Student(stuRes)
	return
}

// //////////////////////


func student2Model(student *demo.Student) *model.Student {
	return &model.Student{
		Id:             student.Id,
		Name:           student.Name,
		Email:          strings.Join(student.Email, ","),
		CollegeName:    student.College.Name,
		CollegeAddress: student.College.Address,
	}
}

func model2Student(student *model.Student) *demo.Student {
	return &demo.Student{
		Id:      student.Id,
		Name:    student.Name,
		Email:   strings.Split(student.Email, ","),
		College: &demo.College{Name: student.CollegeName, Address: student.CollegeAddress},
	}
}
