package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /v1/ping Basic HandlePing
// ---
// produces:
//   - application/json
//   - application/text
//
// summary: Get basic service information to aid in debugging.
// description: returns clean JSON object to check that everything is fine
// parameters:
//   - name: Accept
//     in: header
//     description: standard "Accept" header values
//     type: string
//     required: true
//
// responses:
//
//	'200':
//		description: "OK; returns empty page or json structure"
func (api *API) HandlePing(c *gin.Context) {

	c.JSON(http.StatusOK, "pong :)")
}
