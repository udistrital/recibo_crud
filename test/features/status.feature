Feature: Validate API responses
  RECIBO_CRUD
  probe JSON responses


Scenario Outline: To probe response route /
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"
  And the response should match json "<bodyres>"

  Examples:
  |method|route|bodyreq               |codres|bodyres              |
  |GET   |/    |./files/req/Vacio.json|200 OK|./files/res0/Vok.json|
