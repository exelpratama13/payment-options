package repository

import (
	"payment-options/internal/models"
	//"time"
)

type paymentRepo struct{}

func NewPaymentRepo() PaymentRepository {
	return &paymentRepo{}
}

func (r *paymentRepo) CallOVO() models.PaymentMethod {
// Simulate network delay
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "10000",
		Icon:    "https://sampleurl.com/ovo.jpg",
	}
}

func (r *paymentRepo) CallDANA() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "10000",
		Icon:    "https://sampleurl.com/dana.jpg",
	}
}

func (r *paymentRepo) CallGoPay() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "10000",
		Icon:    "https://sampleurl.com/gopay.jpg",
	}
}

func (r *paymentRepo) CallShopee() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "10000",
		Icon:    "https://sampleurl.com/shopee.jpg",
	}
}

func (r *paymentRepo) CallOneKlik() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "10000",
		Icon:    "https://sampleurl.com/oneklik.jpg",
	}
}

func (r *paymentRepo) CallBRIDD() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "10000",
		Icon:    "https://sampleurl.com/bridd.jpg",
	}

func (r *paymentRepo) CallLinkAja() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "0812xx", 
		Status:  "Active", 
		Balance: "150000", 
		Icon:    "https://sampleurl.com/linkaja.jpg",
	}
}


