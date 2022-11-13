package api

import (
	"context"
	"log_task/dao"
	"log_task/service"
	"log_task/utils/errmsg"
)

type TaskService struct {
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (*TaskService) TaskCreate(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task dao.Task
	resp = new(service.CommonResponse)
	resp.Code = errmsg.SUCCESS
	err = task.Create(req)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetMsg(errmsg.ERROR)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = errmsg.GetMsg(uint(resp.Code))
	return resp, nil
}

func (*TaskService) TaskShow(ctx context.Context, req *service.TaskRequest) (resp *service.TasksDetailResponse, err error) {
	var t dao.Task
	resp = new(service.TasksDetailResponse)
	tRep, err := t.Show(req)
	resp.Code = errmsg.SUCCESS
	if err != nil {
		resp.Code = errmsg.ERROR
		return resp, err
	}
	resp.TaskDetail = dao.BuildTasks(tRep)
	return resp, nil
}

func (*TaskService) TaskUpdate(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task dao.Task
	resp = new(service.CommonResponse)
	resp.Code = errmsg.SUCCESS
	err = task.Update(req)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetMsg(errmsg.ERROR)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = errmsg.GetMsg(uint(resp.Code))
	return resp, nil
}

func (*TaskService) TaskDelete(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task dao.Task
	resp = new(service.CommonResponse)
	resp.Code = errmsg.SUCCESS
	err = task.Delete(req)
	if err != nil {
		resp.Code = errmsg.ERROR
		resp.Msg = errmsg.GetMsg(errmsg.ERROR)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = errmsg.GetMsg(uint(resp.Code))
	return resp, nil
}
