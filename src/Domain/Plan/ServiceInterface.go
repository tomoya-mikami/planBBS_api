package Plan

type ServiceInterface interface {
	Add() error
	FindAll() ([]Plan, error)
}
