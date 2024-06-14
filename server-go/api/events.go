//go:generate oapi-codegen --config=events.gen.yml ../../tsp-output/@typespec/openapi3/openapi.v1.yaml
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewEventService(store EventStorer) *EventService {
	return &EventService{
		store,
	}
}

type EventService struct {
	Store EventStorer
}

func (es *EventService) EventsList(c *gin.Context, params EventsListParams) {
	c.JSON(http.StatusOK, es.Store.FindByNameFuzzy(params.Filter))
}

func (es *EventService) EventsCreate(c *gin.Context) {
	var e Event
	c.Bind(&e)
	newEvent, err := es.Store.Add(e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Message: "failed to create event",
		})
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}

func (es *EventService) EventsDelete(c *gin.Context, id string) {
	if ok := es.Store.DeleteById(id); !ok {
		c.JSON(http.StatusInternalServerError, Response{
			Message: "failed to delete event",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Message: "event deleted",
	})
}

func (es *EventService) EventsRead(c *gin.Context, id string) {
	event, err := es.Store.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, event)
}

type Response struct {
	Message string `json:"message"`
}
