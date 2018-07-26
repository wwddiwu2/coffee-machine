package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mkqavi/coffee-machine/v0/pkg/coffeemachine"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	addr = "http://localhost:31565"
)

var _ = Describe("Coffeed", func() {
	var (
		stop chan os.Signal
		done <-chan bool
	)

	BeforeSuite(func() {
		stop = make(chan os.Signal, 1)
		done = serveAPI(stop)
	})

	AfterSuite(func() {
		stop <- os.Interrupt
		<-done
	})

	Describe("Serving the API", func() {
		It("Should serve /", func() {
			c := &http.Client{}
			resp, err := c.Get(addr)
			Ω(err).Should(BeNil())
			defer resp.Body.Close()
			var s status
			Ω(json.NewDecoder(resp.Body).Decode(s))
		})

		It("Should serve /brew", func() {
			c := &http.Client{}
			body := coffee{uint8(coffeemachine.Espresso)}
			bodyjson, err := json.Marshal(body)
			Ω(err).Should(BeNil())
			resp, err := c.Post(addr+"/brew", "application/json", bytes.NewBuffer(bodyjson))
			Ω(err).Should(BeNil())
			defer resp.Body.Close()
			resbody, err := ioutil.ReadAll(resp.Body)
			Ω(err).Should(BeNil())
			Ω(string(resbody)).Should(Equal("{}"))
		})

		It("Should serve /clean", func() {
			c := &http.Client{}
			resp, err := c.Post(addr+"/clean", "application/json", nil)
			Ω(err).Should(BeNil())
			defer resp.Body.Close()
			resbody, err := ioutil.ReadAll(resp.Body)
			Ω(err).Should(BeNil())
			Ω(string(resbody)).Should(Equal("{}"))
		})
	})
})
