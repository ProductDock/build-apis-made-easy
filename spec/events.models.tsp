import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

namespace EventService;

using TypeSpec.Http;
using TypeSpec.Rest;

@resource("events")
model Event {
  @key
  id: string;
  @doc("Name of the event")
  name: string;
  @doc("When the event takes place")
  date: utcDateTime;
  @doc("Where will it be held")
  location: Location;
}

model Location {
  name: string;
}

@doc("Error")
@error
model EventError {
  code: int32;
  message: string;
}