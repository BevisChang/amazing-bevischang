package rpc

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/AmazingTalker/go-amazing/internal/daomock"
	"github.com/AmazingTalker/go-amazing/pkg/dao"
	"github.com/AmazingTalker/go-amazing/pkg/pb"
)

func ExpectRecordMatcher(target pb.Record, expectation pb.Record) {
	Expect(target.TheStr).To(Equal(expectation.TheStr))
	Expect(target.TheNum).To(Equal(expectation.TheNum))
}

var _ = Describe("GoAmazingServer", func() {

	var serv GoAmazingServer
	var ctrl *gomock.Controller
	var mockRecordDao *daomock.MockRecordDAO

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRecordDao = daomock.NewMockRecordDAO(ctrl)
		serv = NewGoAmazingServer(GoAmazingServerOpt{
			RecordDao: mockRecordDao,
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	var _ = Describe("Record", func() {

		var _ = Describe("create", func() {

			var err error
			var req *pb.CreateRecordReq
			var res *pb.CreateRecordRes

			JustBeforeEach(func() {
				res, err = serv.CreateRecord(context.Background(), req)
			})

			Describe("success", func() {

				BeforeEach(func() {

					// my test case
					req = &pb.CreateRecordReq{
						TheNum: 111,
						TheStr: "xxx",
					}

					// the expectation of dao method arguments
					expectCreateArg1 := &dao.Record{
						TheNum: req.TheNum,
						TheStr: req.TheStr,
					}

					mockRecordDao.EXPECT().CreateRecord(mockCtx, expectCreateArg1).Return(nil).Times(1)
				})

				It("no error", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("return record", func() {

					expectation := pb.Record{
						TheNum: req.TheNum,
						TheStr: req.TheStr,
					}
					Expect(res).NotTo(BeNil())

					result := res.Record

					Expect(result).NotTo(BeNil())
					Expect(result.ID).NotTo(BeNil())
					ExpectRecordMatcher(*result, expectation)
				})

			})

			Describe("duplicate (demo, it is not real)", func() {

				var duplicateErr error

				BeforeEach(func() {

					// my test case
					req = &pb.CreateRecordReq{
						TheNum: 111,
						TheStr: "xxx",
					}

					// the expectation of dao method arguments
					expectCreateArg1 := &dao.Record{
						TheNum: req.TheNum,
						TheStr: req.TheStr,
					}

					duplicateErr = errors.New("the record is duplicated")

					// Use mocker to return error is easier for you guy to test the business logic you wrote.
					mockRecordDao.EXPECT().CreateRecord(mockCtx, expectCreateArg1).Return(duplicateErr).Times(1)
				})

				It("got error", func() {
					Expect(err).To(HaveOccurred())
					Expect(err).To(Equal(duplicateErr))
				})

				It("no response", func() {
					Expect(res).To(BeNil())
				})

			})
		})
	})
})

func TestGoAmazingService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoAmazingService Suite")
}
