package rpc

import (
	"context"
	"errors"
	"github.com/AmazingTalker/bevis-chang/internal/daomock"
	"github.com/AmazingTalker/bevis-chang/pkg/dao"
	"github.com/AmazingTalker/bevis-chang/pkg/pb"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
	"testing"
	"time"
)

func ExpectRecordMatcher(target pb.Record, expectation pb.Record) {
	Expect(target.TheStr).To(Equal(expectation.TheStr))
	Expect(target.TheNum).To(Equal(expectation.TheNum))
}

var _ = Describe("BevisChangServer", func() {

	var serv BevisChangServer
	var ctrl *gomock.Controller
	var mockRecordDao *daomock.MockRecordDAO
	var mockMemberDao *daomock.MockMemberDAO

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRecordDao = daomock.NewMockRecordDAO(ctrl)
		serv = NewBevisChangServer(BevisChangServerOpt{
			RecordDao: mockRecordDao,
			MemberDao: mockMemberDao,
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

	var _ = Describe("Member", func() {
		var _ = Describe("create", func() {

			var err error
			var req *pb.CreateMemberReq
			var res *pb.CreateMemberRes
			now := time.Now()

			JustBeforeEach(func() {
				res, err = serv.CreateMember(context.Background(), req)
			})

			Describe("success", func() {
				req = &pb.CreateMemberReq{
					Name:     "John",
					Birthday: &now,
				}

				expectCreateArg1 := &dao.Member{
					Name:     req.Name,
					Birthday: req.Birthday,
				}

				mockMemberDao.EXPECT().CreateMember(mockCtx, expectCreateArg1).Return(nil).Times(1)
			})

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("return member", func() {

				expectation := pb.Member{
					Name:     req.Name,
					Birthday: req.Birthday,
				}
				Expect(res).NotTo(BeNil())

				result := res.Member

				Expect(result).NotTo(BeNil())
				Expect(result.ID).NotTo(BeNil())
				Expect(result.Name).To(Equal(expectation.Name))
				Expect(result.Birthday).To(Equal(expectation.Birthday))
			})
		})

		var _ = Describe("update", func() {

			var err error
			var req *pb.UpdateMemberReq
			var res *pb.UpdateMemberRes
			now := time.Now()

			JustBeforeEach(func() {
				res, err = serv.UpdateMember(context.Background(), req)
			})

			Describe("success", func() {
				req = &pb.UpdateMemberReq{
					ID:       "1",
					Name:     "John",
					Birthday: &now,
				}

				expectUpdateArg1 := &dao.Member{
					Name:     req.Name,
					Birthday: req.Birthday,
				}

				mockMemberDao.EXPECT().UpdateMember(mockCtx, expectUpdateArg1).Return(nil).Times(1)
			})

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("return member", func() {

				expectation := pb.Member{
					Name:     req.Name,
					Birthday: req.Birthday,
				}
				Expect(res).NotTo(BeNil())

				result := res.Member

				Expect(result).NotTo(BeNil())
				Expect(result.ID).NotTo(BeNil())
				Expect(result.Name).To(Equal(expectation.Name))
				Expect(result.Birthday).To(Equal(expectation.Birthday))
			})
		})

		var _ = Describe("delete member", func() {

			var err error
			var req *pb.DeleteMemberReq

			JustBeforeEach(func() {
				_, err = serv.DeleteMember(context.Background(), req)
			})

			Describe("success", func() {
				req = &pb.DeleteMemberReq{
					ID: "1",
				}

				id, _ := strconv.ParseInt(req.ID, 10, 64)
				expectDeleteArg1 := &dao.Member{
					ID: id,
				}

				mockMemberDao.EXPECT().DeleteMember(mockCtx, expectDeleteArg1).Return(nil).Times(1)
			})

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})

func TestBevisChangService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BevisChangService Suite")
}
