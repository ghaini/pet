package http

import (
	"encoding/json"
	"fmt"
	"go-crud/src/domain"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	Svc domain.PetSvc
}

func NewHandler(svc domain.PetSvc) *Handler {
	return &Handler{
		Svc: svc,
	}
}

// swagger:route GET /pets/{id} Pet GetPet
// provides the detail of the pet with the given id
// responses:
//
//	400: ErrorResponse
//	500: ErrorResponse
//	200: GetPetResponse
func (h *Handler) GetPet(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value("id").(string)
	fmt.Println("pet id : ", ID)
	petID, err := uuid.Parse(ID)
	if err != nil {
		resp := Resp{
			Code: http.StatusBadRequest,
			Msg:  "invalid pet_id provided in url param",
		}
		respond(w, r, &resp)
		return
	}
	if petID == uuid.Nil {
		resp := Resp{
			Code: http.StatusBadRequest,
			Msg:  "please provide the pet id to retrieve",
		}
		respond(w, r, &resp)
		return
	}
	pet, err := h.Svc.Get(petID)
	if err != nil {
		fmt.Println(fmt.Errorf("error - fetching pet detail from db failed, err : %v", err))
		resp := Resp{
			Code: http.StatusInternalServerError,
			Msg:  "fetching pet detail failed. please try again later",
		}
		respond(w, r, &resp)
		return
	}
	resp := Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: pet,
	}
	respond(w, r, &resp)
}

// swagger:route GET /pets Pet ListPets
// provides the details of all pets
// responses:
//
//	500: ErrorResponse
//	200: ListPetsResponse
func (h *Handler) ListPets(w http.ResponseWriter, r *http.Request) {
	pets, err := h.Svc.List()
	if err != nil {
		fmt.Println(fmt.Errorf("error - fetching pet details for given category from db failed, err : %v", err))
		resp := Resp{
			Code: http.StatusInternalServerError,
			Msg:  "fetching all pets detail failed. please try again later",
		}
		respond(w, r, &resp)
		return
	}
	resp := Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: pets,
	}
	respond(w, r, &resp)
}

// swagger:route POST /pets Pet AddPet
// add a new pet detail
// responses:
//
//	500: ErrorResponse
//	200: SuccessRespWithoutData
func (h *Handler) AddPet(w http.ResponseWriter, r *http.Request) {

	var req *domain.Pet
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	err := h.Svc.Create(req)
	if err != nil {
		fmt.Println(fmt.Errorf("error - adding new pet detail to db failed, err : %v", err))
		resp := Resp{
			Code: http.StatusInternalServerError,
			Msg:  "adding new pet failed. please try again later",
		}
		respond(w, r, &resp)
		return
	}
	resp := Resp{
		Code: http.StatusOK,
		Msg:  "success",
	}
	respond(w, r, &resp)
}

// swagger:route DELETE /pets/{id} Pet DeletePet
// delete the pet detail with given id
// responses:
//
//	400: ErrorResponse
//	500: ErrorResponse
//	200: SuccessRespWithoutData
func (h *Handler) DeletePet(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value("id").(string)
	fmt.Println("pet id : ", ID)
	petID, err := uuid.Parse(ID)
	if err != nil {
		resp := Resp{
			Code: http.StatusBadRequest,
			Msg:  "invalid pet_id provided in url param",
		}
		respond(w, r, &resp)
		return
	}
	if petID == uuid.Nil {
		resp := Resp{
			Code: http.StatusBadRequest,
			Msg:  "please provide the pet id to retrieve",
		}
		respond(w, r, &resp)
		return
	}
	err = h.Svc.Delete(petID)
	if err != nil {
		fmt.Println(fmt.Errorf("error - deleting pet record from db failed, err: %v", err))
		resp := Resp{
			Code: http.StatusInternalServerError,
			Msg:  "unable to delete the pet details. please try again later",
		}
		respond(w, r, &resp)
		return
	}
	resp := Resp{
		Code: http.StatusOK,
		Msg:  "success",
	}
	respond(w, r, &resp)
}
