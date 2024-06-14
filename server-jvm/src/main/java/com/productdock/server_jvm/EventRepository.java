package com.productdock.server_jvm;

import com.productdock.openapi.model.Event;

import java.util.List;

public interface EventRepository {
    Event add(Event event);

    Event findById(String id);

    void deleteById(String id);

    List<Event> findByNameFuzzy(String name);
}
