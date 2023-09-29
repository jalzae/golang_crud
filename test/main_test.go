package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	helper "rest/helper"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRootEndpoint(t *testing.T) {
	r := gin.Default()
	r.Use(helper.Cors())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connected",
		})
	})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expected := `{"message":"Connected"}`
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
		fmt.Println("Test failed: status code does not match")
	} else if w.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, w.Body.String())
		fmt.Println("Test failed: response body does not match")
	} else {
		fmt.Println("Test passed")
	}
}
