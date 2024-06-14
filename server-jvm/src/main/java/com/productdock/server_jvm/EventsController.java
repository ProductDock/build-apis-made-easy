package com.productdock.server_jvm;

import com.productdock.openapi.api.EventsApi;
import com.productdock.openapi.model.Event;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
class EventsController implements EventsApi {

    private final EventRepository eventRepository;

    EventsController(EventRepository eventService) {
        this.eventRepository = eventService;
    }

    @Override
    public ResponseEntity<Event> eventsCreate(Event event) {
        Event createdEvent = eventRepository.add(event);
        return new ResponseEntity<>(createdEvent, HttpStatus.CREATED);
    }

    @Override
    public ResponseEntity<Void> eventsDelete(String id) {
        eventRepository.deleteById(id);
        return new ResponseEntity<>(HttpStatus.NO_CONTENT);
    }

    @Override
    public ResponseEntity<List<Event>> eventsFindByNameFuzzy(String name) {
        List<Event> matchingEvents = eventRepository.findByNameFuzzy(name);
        return new ResponseEntity<>(matchingEvents, HttpStatus.OK);
    }

    @Override
    public ResponseEntity<Event> eventsRead(String id) {
        Event foundEvent = eventRepository.findById(id);
        if (foundEvent == null) {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        }
        return new ResponseEntity<>(foundEvent, HttpStatus.OK);
    }
}