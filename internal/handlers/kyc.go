package handlers

import (
	"net/http"

	"philcoin/internal/services"
	"philcoin/internal/utils"
)

// UploadID handles ID document upload and OCR processing
func UploadID(w http.ResponseWriter, r *http.Request) {
	// Limit upload size (5MB)
	r.Body = http.MaxBytesReader(w, r.Body, 5<<20)

	file, _, err := r.FormFile("id_image")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "ID image is required and must be <=5MB")
		return
	}
	defer file.Close()

	rec, err := services.ProcessIDImage(file)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, rec)
}

// SubmitSelfie handles selfie upload for liveness check
func SubmitSelfie(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 5<<20)

	file, _, err := r.FormFile("selfie_image")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Selfie image is required and must be <=5MB")
		return
	}
	defer file.Close()

	ok, err := services.VerifyLiveness(file)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "Liveness verification failed")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"status": "verified"})
}
