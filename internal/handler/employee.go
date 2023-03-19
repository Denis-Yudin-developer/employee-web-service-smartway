package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"smartway-test-task/internal/model"
	"strconv"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	logrus.Print("Запрос на маршрут /create")
	var employee model.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректный запрос")
	}
	defer r.Body.Close()

	err := model.Validate(employee)
	if err != nil {
		logrus.Printf("ДАННЫЕ СОТРУДНИКА НЕ ПРОШЛИ ВАЛИДАЦИЮ. ПРИЧИНА: %s", err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Employee.Create(employee)
	if err != nil {
		logrus.Printf("НЕ УДАЛОСЬ СОЗДАТЬ СОТРУДНИКА. ПРИЧИНА: %s", err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := make(map[string]int, 1)
	res["id"] = id
	logrus.Printf("id транзакции: %d", res["id"])
	respondWithJSON(w, http.StatusOK, res)
}

func (h *Handler) GetAllByCompany(w http.ResponseWriter, r *http.Request) {
	logrus.Print("Запрос на маршрут /get?company=id")
	queryParams := r.URL.Query()
	companyId, err := strconv.Atoi(queryParams.Get("company"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректный запрос")
		return
	}
	defer r.Body.Close()

	employees, err := h.services.Employee.GetAllByCompany(companyId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Printf("ПОЛУЧЕНЫ СОТРУДНИКИ КОМПАНИИ ПО ID %d", companyId)
	respondWithJSON(w, http.StatusOK, employees)
}

func (h *Handler) GetAllByDepartment(w http.ResponseWriter, r *http.Request) {
	logrus.Print("Запрос на маршрут /get?department=name")
	queryParams := r.URL.Query()
	department := queryParams.Get("department")
	employees, err := h.services.Employee.GetAllByDepartment(department)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	logrus.Printf("ПОЛУЧЕНЫ СОТРУДНИКИ ДЕПАРТАМЕНТА ПО ID %s", department)
	respondWithJSON(w, http.StatusOK, employees)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	logrus.Print("Запрос на маршрут /update/{id}")
	vars := mux.Vars(r)
	employeeIdString := vars["id"]
	employeeId, err := strconv.Atoi(employeeIdString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректный запрос")
		return
	}
	var updatedEmployee model.UpdateEmployee

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedEmployee); err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректный запрос")
		return
	}
	defer r.Body.Close()

	err = h.services.Employee.Update(updatedEmployee, employeeId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := make(map[string]string, 1)
	res["status"] = "success"
	logrus.Printf("статус обновления пользователя с id %d: %s", employeeId, res["status"])
	respondWithJSON(w, http.StatusOK, res)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	logrus.Print("Запрос на маршрут /delete/{id}")
	vars := mux.Vars(r)
	userIdString := vars["id"]
	employeeId, err := strconv.Atoi(userIdString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректный запрос")
		return
	}

	err = h.services.Employee.Delete(employeeId)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	res := make(map[string]string, 1)
	res["status"] = "success"
	logrus.Printf("Пользователь с id %d успешно удален", employeeId)
	respondWithJSON(w, http.StatusOK, res["status"])
}
