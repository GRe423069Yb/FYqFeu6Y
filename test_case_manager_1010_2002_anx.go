// 代码生成时间: 2025-10-10 20:02:42
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "log"
)

// TestCase struct to hold test case data
type TestCase struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Input     string `json:"input"`
    Expected  string `json:"expected"`
    Actual    string `json:"actual"`
    Status    string `json:"status"`
}

// TestCaseManager struct to manage test cases
type TestCaseManager struct {
    testCases []TestCase
}

// NewTestCaseManager creates a new instance of TestCaseManager
func NewTestCaseManager() *TestCaseManager {
    return &TestCaseManager{
        testCases: make([]TestCase, 0),
    }
}

// AddTestCase adds a new test case to the manager
func (m *TestCaseManager) AddTestCase(tc TestCase) {
    // Add the test case to the list
    m.testCases = append(m.testCases, tc)
}

// GetTestCases returns all test cases
func (m *TestCaseManager) GetTestCases() []TestCase {
    return m.testCases
}

// UpdateTestCase updates a test case by ID
func (m *TestCaseManager) UpdateTestCase(id string, tc TestCase) error {
    for i, existingTc := range m.testCases {
        if existingTc.ID == id {
            m.testCases[i] = tc
            return nil
        }
    }
    return gin.Error{
        Err:  ErrTestCaseNotFound,
        Type: gin.ErrorTypeAny,
        Meta: nil,
    }
}

// DeleteTestCase deletes a test case by ID
func (m *TestCaseManager) DeleteTestCase(id string) error {
    for i, existingTc := range m.testCases {
        if existingTc.ID == id {
            m.testCases = append(m.testCases[:i], m.testCases[i+1:]...)
            return nil
        }
    }
    return gin.Error{
        Err:  ErrTestCaseNotFound,
        Type: gin.ErrorTypeAny,
        Meta: nil,
    }
}

// Test Case Manager API Handlers
func testCasesHandler(manager *TestCaseManager) gin.HandlerFunc {
    return func(c *gin.Context) {
        switch c.Request.Method {
        case http.MethodGet:
            testCases := manager.GetTestCases()
            c.JSON(http.StatusOK, testCases)
        case http.MethodPost:
            var tc TestCase
            if err := c.ShouldBindJSON(&tc); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }
            manager.AddTestCase(tc)
            c.JSON(http.StatusOK, tc)
        default:
            c.Status(http.StatusMethodNotAllowed)
        }
    }
}

func main() {
    r := gin.Default()
    r.Use(gin.Recovery()) // Use Recovery middleware to handle errors

    manager := NewTestCaseManager()

    // Define test cases API route
    r.GET("/test_cases", testCasesHandler(manager))
    r.POST("/test_cases", testCasesHandler(manager))

    // Start the server
    r.Run() // listening and serving on 0.0.0.0:8080
}
