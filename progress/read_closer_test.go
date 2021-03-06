package progress_test

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/pivotal-cf/om/progress"
	"github.com/pivotal-cf/om/progress/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReadCloser", func() {
	Describe("Read", func() {
		var (
			bar               *fakes.ProgressBar
			readCloser        *progress.ReadCloser
			callbackWasCalled bool
		)

		BeforeEach(func() {
			reader := strings.NewReader("abcdefghij")
			bar = &fakes.ProgressBar{}
			bar.NewProxyReaderReturns(ioutil.NopCloser(reader))

			callbackWasCalled = false
			readCloser = progress.NewReadCloser(reader, bar, func() { callbackWasCalled = true })
		})

		It("reads from the wrapped reader", func() {
			buffer := make([]byte, 10)
			n, err := readCloser.Read(buffer)
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(10))
			Expect(buffer).To(Equal([]byte("abcdefghij")))

			Expect(callbackWasCalled).To(BeFalse())

			n, err = readCloser.Read(buffer)
			Expect(err).To(MatchError(io.EOF))
			Expect(n).To(Equal(0))

			Expect(callbackWasCalled).To(BeTrue())
		})

		It("starts the progress bar only once", func() {
			_, err := readCloser.Read([]byte{})
			Expect(err).NotTo(HaveOccurred())

			Expect(bar.StartCallCount()).To(Equal(1))

			_, err = readCloser.Read([]byte{})
			Expect(err).NotTo(HaveOccurred())

			Expect(bar.StartCallCount()).To(Equal(1))
		})

		Context("when io.EOF is returned from the reader", func() {
			It("stops the progress bar", func() {
				_, err := readCloser.Read(make([]byte, 10))
				Expect(err).NotTo(HaveOccurred())

				Expect(bar.FinishCallCount()).To(Equal(0))

				_, err = readCloser.Read(make([]byte, 10))
				Expect(err).To(MatchError(io.EOF))

				Expect(bar.FinishCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Close", func() {
		var (
			bar               *fakes.ProgressBar
			closer            *fakeCloser
			readCloser        *progress.ReadCloser
			callbackWasCalled bool
		)

		BeforeEach(func() {
			bar = &fakes.ProgressBar{}
			closer = &fakeCloser{}
			bar.NewProxyReaderReturns(closer)

			readCloser = progress.NewReadCloser(closer, bar, func() { callbackWasCalled = true })
		})

		It("closes the underlying closer and calls Finish", func() {
			Expect(closer.closeWasCalled).To(BeFalse())
			Expect(bar.FinishCallCount()).To(Equal(0))
			Expect(callbackWasCalled).To(BeFalse())

			err := readCloser.Close()
			Expect(err).NotTo(HaveOccurred())

			Expect(closer.closeWasCalled).To(BeTrue())
			Expect(bar.FinishCallCount()).To(Equal(1))
			Expect(callbackWasCalled).To(BeTrue())
		})

		Context("when there is no callback", func() {
			It("does not panic", func() {
				readCloser = progress.NewReadCloser(closer, bar, nil)

				Expect(func() {
					readCloser.Close()
				}).NotTo(Panic())
			})
		})
	})
})

type fakeCloser struct {
	closeWasCalled bool
}

func (f *fakeCloser) Read([]byte) (int, error) {
	return -1, nil
}

func (f *fakeCloser) Close() error {
	f.closeWasCalled = true
	return nil
}
