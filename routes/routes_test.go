package routes

import (
	"encoding/json"
	"fmt"
	"github.com/kentpon/LetsGO/routes/models"
	"github.com/kentpon/LetsGO/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

var _ = Describe("Test Root Routes", func() {
	It("GetHomePage", func() {
		res := utils.TestRequest(Router, "/", http.MethodGet, "", nil)
		Expect(res.Code).To(Equal(http.StatusOK))
		Expect(res.Body.String()).To(Equal(welcomeString))
	})
})

var _ = Describe("Test Notes", func() {
	const (
		RouteNoteList = "/notes"
		RouteNote     = RouteNoteList + "/%s"
	)

	testNote := models.Note{
		Title:  "test title",
		Detail: "test detail",
	}
	It("GetEmptyNoteList", func() {
		var notes []models.Note
		res := utils.TestRequest(Router, RouteNoteList, http.MethodGet, "", nil)
		json.Unmarshal(res.Body.Bytes(), &notes)
		fmt.Println("Get note list response", notes)
		Expect(res.Code).To(Equal(http.StatusOK))
		Expect(len(notes)).Should(BeZero())
	})

	It("AddNote", func() {
		res := utils.TestRequest(Router, RouteNoteList, http.MethodPost, "", testNote)
		fmt.Println("Add note")
		Expect(res.Code).To(Equal(http.StatusCreated))
	})

	It("CheckNote", func() {
		By("Get Note List")
		var notes []models.Note
		res := utils.TestRequest(Router, RouteNoteList, http.MethodGet, "", nil)
		json.Unmarshal(res.Body.Bytes(), &notes)
		fmt.Println("Get note list response", notes)
		Expect(res.Code).To(Equal(http.StatusOK))
		Expect(len(notes)).To(Equal(1))

		note := notes[0]
		Expect(note.ID).NotTo(Equal(uuid.Nil))
		Expect(note.Title).To(Equal(testNote.Title))
		Expect(note.Detail).To(Equal(testNote.Detail))
		nilTime := time.Time{}
		Expect(note.CreatedAt).NotTo(Equal(nilTime))
		Expect(note.UpdatedAt).NotTo(Equal(nilTime))
		Expect(note.DeletedAt).Should(BeNil())

		By("Get Note")
		note = models.Note{}
		res = utils.TestRequest(Router, fmt.Sprintf(RouteNote, notes[0].ID.String()), http.MethodGet, "", nil)
		json.Unmarshal(res.Body.Bytes(), &note)
		fmt.Println("Get note response", note)
		Expect(res.Code).To(Equal(http.StatusOK))
		Expect(note.ID).NotTo(Equal(uuid.Nil))
		Expect(note.Title).To(Equal(testNote.Title))
		Expect(note.Detail).To(Equal(testNote.Detail))
		Expect(note.CreatedAt).NotTo(Equal(nilTime))
		Expect(note.UpdatedAt).NotTo(Equal(nilTime))
		Expect(note.DeletedAt).Should(BeNil())

		testNote = note
	})

	It("ModifyNote", func() {
		By("Modify Note")
		newNoteReq := models.Note{
			Title:  "test title 2",
			Detail: "test detail 2",
		}
		res := utils.TestRequest(Router, fmt.Sprintf(RouteNote, testNote.ID.String()), http.MethodPatch, "", newNoteReq)

		fmt.Println("Get modify note response")
		Expect(res.Code).To(Equal(http.StatusOK))

		By("Get New Note")
		var notes []models.Note
		res = utils.TestRequest(Router, RouteNoteList, http.MethodGet, "", nil)
		json.Unmarshal(res.Body.Bytes(), &notes)
		fmt.Println("Get note list response", notes)
		Expect(res.Code).To(Equal(http.StatusOK))
		newNote := notes[0]
		Expect(newNote.ID).To(Equal(testNote.ID))
		Expect(newNote.Title).To(Equal(newNoteReq.Title))
		Expect(newNote.Detail).To(Equal(newNoteReq.Detail))
		Expect(newNote.CreatedAt).To(Equal(testNote.CreatedAt))
		Expect(newNote.UpdatedAt).NotTo(Equal(testNote.UpdatedAt))
		Expect(newNote.DeletedAt).Should(BeNil())

		testNote = newNote
	})

	It("DeleteNote", func() {
		res := utils.TestRequest(Router, fmt.Sprintf(RouteNote, testNote.ID.String()), http.MethodDelete, "", nil)
		fmt.Println("Get delete note response")
		Expect(res.Code).To(Equal(http.StatusOK))

		var notes []models.Note
		res = utils.TestRequest(Router, RouteNoteList, http.MethodGet, "", nil)
		json.Unmarshal(res.Body.Bytes(), &notes)
		fmt.Println("Get note list response", notes)
		Expect(res.Code).To(Equal(http.StatusOK))
		Expect(len(notes)).To(Equal(0))

		res = utils.TestRequest(Router, fmt.Sprintf(RouteNote, testNote.ID.String()), http.MethodGet, "", nil)
		Expect(res.Code).To(Equal(http.StatusNotFound))
	})
})
