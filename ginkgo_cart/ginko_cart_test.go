package ginko_cart_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/eliteGoblin/code_4_blog/ginko_cart"
)

var _ = Describe("GinkoCart", func() {
	var (
		cart *ginko_cart.Cart
		err error
	)
	Context(`Start with empty cart`, func(){
		BeforeEach(func() {
			cart = ginko_cart.NewCart()
		})
		When(`Add One A item to cart`, func(){
			BeforeEach(func() {
				err = cart.AddItem(ginko_cart.Item{
					Name: "A",
					Price: 3.99,
					Qty  : 1,
				})
			})
			It(`Should no error`, func() {
				Expect(err).To(BeNil())
			})
			It(`Should display items count as 1`, func() {
				Expect(cart.TotalItems()).To(Equal(1))
			})
			It(`Should display items count as 1`, func() {
				Expect(cart.TotalUniqueItems()).To(Equal(1))
			})
			It(`Should display items total price as A`, func() {
				Expect(cart.TotalPrice()).To(Equal(3.99))
			})
		})
		When(`Add One A and Two B item to cart`, func(){
			BeforeEach(func() {
				err = cart.AddItem(ginko_cart.Item{
					Name: "A",
					Price: 3.99,
					Qty  : 1,
				})
				Expect(err).To(BeNil())
				err = cart.AddItem(ginko_cart.Item{
					Name: "B",
					Price: 12.99,
					Qty  : 2,
				})
				Expect(err).To(BeNil())
			})
			It(`Should display items count as 3`, func() {
				Expect(cart.TotalItems()).To(Equal(3))
			})
			It(`Should display unique items count as 2`, func() {
				Expect(cart.TotalUniqueItems()).To(Equal(2))
			})
			It(`Should display items total price as A+B`, func() {
				Expect(cart.TotalPrice()).To(Equal(3.99+2*12.99))
			})
		})
	})
})
