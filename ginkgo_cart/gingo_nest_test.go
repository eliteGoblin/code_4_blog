package ginko_cart

import (
	"fmt"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Nest Test Demo", func() {
	Context("MyTest level1", func() {
		BeforeEach(func() {
			fmt.Println("beforeEach level 1")
		})
		It("spec 3-1 in level1", func(){
			fmt.Println("sepc on level 1")
		})
		Context("MyTest level2", func() {
			BeforeEach(func() {
				fmt.Println("beforeEach level 2")
			})
			Context("MyTest level3", func() {
				BeforeEach(func() {
					fmt.Println("beforeEach level 3")
				})
				It("spec 3-1 in level3", func() {
					fmt.Println("A simple spec in level 3")
				})
				It("3-2 in level3", func() {
					fmt.Println("A simple spec in level 3")
				})
			})
		})
	})
})