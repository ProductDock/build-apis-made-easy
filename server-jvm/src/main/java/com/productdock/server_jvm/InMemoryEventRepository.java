package com.productdock.server_jvm;

import com.productdock.openapi.model.Event;
import org.springframework.stereotype.Repository;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@Repository
class InMemoryEventRepository implements EventRepository {
    private List<Event> events = new ArrayList<>();

    @Override
    public Event add(Event event) {
        events.add(event);
        return event;
    }

    @Override
    public Event findById(String id) {
        return events.stream().filter(e -> e.getId().equals(id)).findFirst().orElse(null);
    }

    @Override
    public void deleteById(String id) {
        events.removeIf(e -> e.getId().equals(id));
    }

    @Override
    public List<Event> findByNameFuzzy(String name) {
        return events.stream().filter(e -> e.getName().contains(name)).collect(Collectors.toList());
    }
}
