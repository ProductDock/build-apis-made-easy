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
  @doc("The name of the event")
  name: string;
  @doc("The scheduled date and time of the event")
  date: utcDateTime;
  @doc("Information about the venue where the event will be held")
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