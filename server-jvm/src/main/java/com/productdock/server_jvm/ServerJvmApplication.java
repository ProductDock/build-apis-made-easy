package com.productdock.server_jvm;

import com.productdock.openapi.api.EventsApi;
import com.productdock.openapi.model.Event;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;

import java.util.Collections;
import java.util.List;

@SpringBootApplication
public class ServerJvmApplication {

	public static void main(String[] args) {
		SpringApplication.run(ServerJvmApplication.class, args);
	}

}