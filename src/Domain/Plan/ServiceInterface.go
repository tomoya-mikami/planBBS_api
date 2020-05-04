package Plan

type ServiceInterface interface {
	Add() error
	findAll() ([]Plan, error)
}
