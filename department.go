package wework

import json "github.com/json-iterator/go"

// Department 部门详情
type Department struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NameEn   string `json:"name_en"`
	ParentID int    `json:"parentid"`
	Order    int    `json:"order"`
}

// DepartmentList 部门详情列表
type DepartmentList struct {
	Errcode    int          `json:"errcode"`
	Errmsg     string       `json:"errmsg"`
	Department []Department `json:"department"`
}

// AllDept 获取企业所有部门
func (wx *Client) AllDept() ([]Department, error) {
	dept := DepartmentList{}
	res, err := wx.Get(ApiGetDept)
	if err != nil {
		return dept.Department, err
	}
	err = json.Unmarshal(res, &dept)
	if err != nil {
		return dept.Department, err
	}
	return dept.Department, nil
}

// DeptByID 通过部门id获取部门详情
func (wx *Client) DeptByID(departmentID string) ([]Department, error) {
	dept := DepartmentList{}
	res, err := wx.Get(ApiGetDept, "id", departmentID)
	if err != nil {
		return dept.Department, err
	}
	err = json.Unmarshal(res, &dept)
	if err != nil {
		return dept.Department, err
	}
	return dept.Department, nil
}
