package handler

import (
	"context"
	"github.com/wsjcko/shoppayment/common"
	"github.com/wsjcko/shoppayment/domain/model"
	"github.com/wsjcko/shoppayment/domain/service"
	"github.com/wsjcko/shoppayment/logger"
	pb "github.com/wsjcko/shoppayment/protobuf/pb"
)

type ShopPayment struct {
	PaymentService service.IPaymentService
}

func (e *ShopPayment) AddPayment(ctx context.Context, request *pb.PaymentInfo, response *pb.PaymentID) error {
	payment := &model.Payment{}
	if err := common.SwapTo(request, payment); err != nil {
		logger.Error(err)
	}
	paymentID, err := e.PaymentService.AddPayment(payment)
	if err != nil {
		logger.Error(err)
	}
	response.PaymentId = paymentID
	return nil
}

func (e *ShopPayment) UpdatePayment(ctx context.Context, request *pb.PaymentInfo, response *pb.Response) error {
	payment := &model.Payment{}
	if err := common.SwapTo(request, payment); err != nil {
		logger.Error(err)
	}
	return e.PaymentService.UpdatePayment(payment)
}

func (e *ShopPayment) DeletePaymentByID(ctx context.Context, request *pb.PaymentID, response *pb.Response) error {
	return e.PaymentService.DeletePayment(request.PaymentId)
}

func (e *ShopPayment) FindPaymentByID(ctx context.Context, request *pb.PaymentID, response *pb.PaymentInfo) error {
	payment, err := e.PaymentService.FindPaymentByID(request.PaymentId)
	if err != nil {
		logger.Error(err)
	}
	return common.SwapTo(payment, response)
}

func (e *ShopPayment) FindAllPayment(ctx context.Context, request *pb.All, response *pb.PaymentAll) error {
	allPayment, err := e.PaymentService.FindAllPayment()
	if err != nil {
		logger.Error(err)
	}

	for _, v := range allPayment {
		paymentInfo := &pb.PaymentInfo{}
		if err := common.SwapTo(v, paymentInfo); err != nil {
			logger.Error(err)
		}
		response.PaymentInfo = append(response.PaymentInfo, paymentInfo)
	}
	return nil
}
