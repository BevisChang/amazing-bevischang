package dao

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MysqlRecordDAO", func() {

	var dao RecordDAO

	BeforeEach(func() {
		dao = NewMySqlRecordDAO(testConn)
	})

	Describe("CreateRecord", func() {

		var err error
		var record *Record

		JustBeforeEach(func() {
			err = dao.CreateRecord(testCtx, record)
		})

		Describe("success", func() {

			BeforeEach(func() {
				record = &Record{
					TheNum: 111,
					TheStr: "222",
				}
			})

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("record created", func() {
				Expect(record.ID).NotTo(BeNil())
				Expect(record.CreatedAt).NotTo(BeNil())
				Expect(record.UpdatedAt).NotTo(BeNil())
			})

		})

	})

})
