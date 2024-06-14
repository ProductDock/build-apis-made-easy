import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "@typespec/versioning";

import "./events.models.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.Rest.Resource;
using TypeSpec.Versioning;


@service({
  title: "Event Service",
})
@versioned(Versions)
namespace EventService;

@route("/events")
@tag("Events")
interface Events {
  @added(Versions.v1)
  op create(@body event: Event): Event | EventError;

  @added(Versions.v1)
  op list(@query filter: string): Event[];

  @added(Versions.v1)
  op read(@path id: string): Event;

  @delete
  @added(Versions.v1)
  op delete(@path id: string): Event;
}

enum Versions {
  v1
}