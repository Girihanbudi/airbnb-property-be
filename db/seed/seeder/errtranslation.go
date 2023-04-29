package seeder

import (
	middlewareerr "airbnb-property-be/internal/app/middleware/preset/error"
	propertyerr "airbnb-property-be/internal/app/property/preset/error"
	"airbnb-property-be/internal/pkg/stderror"
	"net/http"

	translationmodule "airbnb-property-be/internal/app/translation"

	"gorm.io/gorm"
)

func SeedErrTranslation(db gorm.DB) error {

	data := []translationmodule.ErrTranslation{
		/*
			Default
		*/
		// En translation
		MakeErrTranslation(stderror.DEF_SERVER_500, "en", http.StatusInternalServerError, "Request aborted due to server error"),
		MakeErrTranslation(stderror.DEF_AUTH_401, "en", http.StatusUnauthorized, "Cannot authorize user"),
		MakeErrTranslation(stderror.DEF_DATA_400, "en", http.StatusBadRequest, "Requested data is not valid"),
		// Id translation
		MakeErrTranslation(stderror.DEF_SERVER_500, "id", http.StatusInternalServerError, "Permintaan dibatalkan karena terjadi kesalahan server"),
		MakeErrTranslation(stderror.DEF_AUTH_401, "id", http.StatusUnauthorized, "Tidak dapat mengotorisasi user"),
		MakeErrTranslation(stderror.DEF_DATA_400, "id", http.StatusBadRequest, "Permintaan tidak valid"),

		/*
			Middleware
		*/
		// En translation
		MakeErrTranslation(middlewareerr.TokenNotFound, "en", http.StatusUnauthorized, "Authorization not found"),
		MakeErrTranslation(middlewareerr.TokenNotValid, "en", http.StatusUnauthorized, "Token is not valid"),
		MakeErrTranslation(middlewareerr.UserAlreadyVerified, "en", http.StatusForbidden, "User already verified"),
		// Id translation
		MakeErrTranslation(middlewareerr.TokenNotFound, "id", http.StatusUnauthorized, "Otorisasi tidak ditemukan"),
		MakeErrTranslation(middlewareerr.TokenNotValid, "id", http.StatusUnauthorized, "Token tidak valid"),
		MakeErrTranslation(middlewareerr.UserAlreadyVerified, "id", http.StatusForbidden, "User telah terverifikasi"),

		/*
			Locale
		*/
		// En translation
		MakeErrTranslation(propertyerr.DbServiceUnavailable, "en", http.StatusServiceUnavailable, "Failed to communicate with store server"),
		MakeErrTranslation(propertyerr.DbRecordNotFound, "en", http.StatusNotFound, "Requested data not found"),
		MakeErrTranslation(propertyerr.DbEmptyResult, "en", http.StatusNotFound, "Requested result nothing"),
		// Id translation
		MakeErrTranslation(propertyerr.DbServiceUnavailable, "id", http.StatusServiceUnavailable, "Gagal berkomunikasi dengan server penyimpanan"),
		MakeErrTranslation(propertyerr.DbRecordNotFound, "id", http.StatusNotFound, "Data tidak ditemukan"),
		MakeErrTranslation(propertyerr.DbEmptyResult, "id", http.StatusNotFound, "Tidak ada hasil apapun"),
	}

	var errTranslationRecords []translationmodule.ErrTranslation
	if err := db.Find(&errTranslationRecords).Error; err != nil {
		return err
	}

	if len(errTranslationRecords) > 0 {
		if err := db.Delete(&errTranslationRecords).Error; err != nil {
			return err
		}
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func MakeErrTranslation(code, localeCode string, httpCode int, message string) translationmodule.ErrTranslation {
	return translationmodule.ErrTranslation{
		Code:       code,
		LocaleCode: localeCode,
		HttpCode:   httpCode,
		Message:    message,
	}
}
