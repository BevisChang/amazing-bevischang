package dao

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("MysqlMemberDAO", func() {

	var dao MySqlMemberDAO

	BeforeEach(func() {
		dao = NewMySqlMemberDAO(testConn)
	})

	Describe("CreateMember", func() {

		var err error
		var member *Member

		JustBeforeEach(func() {
			err = dao.CreateMember(testCtx, member)
		})

		Describe("success", func() {

			BeforeEach(func() {
				now := time.Now()
				member = &Member{Name: "John", Birthday: &now}
			})

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("member created", func() {
				Expect(member.ID).NotTo(BeNil())
				Expect(member.CreatedAt).NotTo(BeNil())
				Expect(member.UpdatedAt).NotTo(BeNil())
			})

		})

	})

	Describe("UpdateMember", func() {

		var err error
		var member *Member
		now := time.Now()

		JustBeforeEach(func() {
			member, err = dao.UpdateMember(testCtx, member)
		})

		Describe("success", func() {

			BeforeEach(func() {
				dao.CreateMember(testCtx, &Member{ID: 1, Name: "John"})
				member = &Member{ID: 1, Name: "Sean", Birthday: &now}
			})

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("member updated", func() {
				Expect(member.ID).To(Equal(int64(1)))
				Expect(member.Name).To(Equal("Sean"))
				Expect(member.Birthday).To(Equal(&now))
				Expect(member.CreatedAt).NotTo(BeNil())
				Expect(member.UpdatedAt).NotTo(BeNil())
			})

		})

	})

	Describe("LisMembers", func() {

		var err error
		var members []Member
		now := time.Now()

		JustBeforeEach(func() {
			members, err = dao.ListMembers(testCtx, &now)
		})

		Describe("success", func() {

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("members list", func() {
				Expect(len(members)).To(BeNumerically(">=", 1))
			})

		})

	})

	Describe("DeleteMember", func() {

		var err error

		JustBeforeEach(func() {
			dao.DeleteMember(testCtx, int64(1))
		})

		Describe("success", func() {

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

})
