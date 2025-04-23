import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
)

func TestPingRoute(t *testing.T) {
    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    req, _ := http.NewRequest("GET", "/ping", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}
