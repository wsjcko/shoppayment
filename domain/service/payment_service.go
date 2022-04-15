package service

import (
	"github.com/wsjcko/shoppayment/domain/model"
	"github.com/wsjcko/shoppayment/domain/repository"
)

type IPaymentService interface {
	AddPayment(*model.Payment) (int64, error)
	DeletePayment(int64) error
	UpdatePayment(*model.Payment) error
	FindPaymentByID(int64) (*model.Payment, error)
	FindAllPayment() ([]model.Payment, error)
}

//创建
func NewPaymentService(paymentRepository repository.IPaymentRepository) IPaymentService {
	return &PaymentService{paymentRepository}
}

type PaymentService struct {
	PaymentRepository repository.IPaymentRepository
}

//插入
func (u *PaymentService) AddPayment(payment *model.Payment) (int64, error) {
	return u.PaymentRepository.CreatePayment(payment)
}

//删除
func (u *PaymentService) DeletePayment(paymentID int64) error {
	return u.PaymentRepository.DeletePaymentByID(paymentID)
}

//更新
func (u *PaymentService) UpdatePayment(payment *model.Payment) error {
	return u.PaymentRepository.UpdatePayment(payment)
}

//查找
func (u *PaymentService) FindPaymentByID(paymentID int64) (*model.Payment, error) {
	return u.PaymentRepository.FindPaymentByID(paymentID)
}

//查找
func (u *PaymentService) FindAllPayment() ([]model.Payment, error) {
	return u.PaymentRepository.FindAll()
}
