package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStrongPasswordSteps(t *testing.T) {
	r := gin.Default()
	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		var req PasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		numOfSteps := calculateNumOfSteps(req.InitPassword)
		c.JSON(http.StatusOK, gin.H{"num_of_steps": numOfSteps})
	})

	testCases := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Example 1",
			requestBody:    `{"init_password": "aA1"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"num_of_steps":3}`,
		},
		{
			name:           "Example 2",
			requestBody:    `{"init_password": "1445D1cd"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"num_of_steps":0}`,
		},
		{
			name:           "Example 3",
			requestBody:    `{"init_password": "AAABB"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"num_of_steps":2}`,
		},
		{
			name:           "Example 4",
			requestBody:    `{"init_password": "aA1bB2cC3dD4eE5fF6gG7hH8iI9jJ0kK"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"num_of_steps":12}`,
		},
		{
			name:           "Example 5",
			requestBody:    `{"init_password": "aaaaa"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"num_of_steps":2}`,
		},
		{
			name:           "Example 6",
			requestBody:    `{"init_password": "bbbbbb"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"num_of_steps":2}`,
		},
		{
			name:           "Example 7",
			requestBody:    `{"init_password": "ccccccc"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"num_of_steps":2}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/api/strong_password_steps", strings.NewReader(tc.requestBody))
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}
