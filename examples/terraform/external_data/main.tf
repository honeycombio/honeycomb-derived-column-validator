# Example 1, using a file input
data "external" "example1" {
  program = ["${path.root}/../../../honeycomb-derived-column-validator"]

  query = {
    expression = file("${path.cwd}/sli.request_latency.honeycomb")
  }
}

resource "honeycombio_derived_column" "dc1" {
  alias       = "sli.request_latency"
  description = "SLI: request latency less than 300ms"
  dataset     = "__all__"

  expression = data.external.example1.result.expression
}

# Example 2, using a heredoc

data "external" "example2" {
  program = ["${path.root}/../../../honeycomb-derived-column-validator"]

  query = {
    expression = <<-EOT
    IF(
      AND(
        NOT(
          EXISTS(
            $trace.parent_id
          )
        ),
        EXISTS(
          $duration_ms
        )
      ),
      LTE(
        $duration_ms,
        300
      )
    )
    EOT
  }
}

resource "honeycombio_derived_column" "dc2" {
  alias       = "sli.request_latency"
  description = "SLI: request latency less than 300ms"
  dataset     = "__all__"

  expression = data.external.example2.result.expression
}

# Example 3, using an inline expression
data "external" "example3" {
  program = ["${path.root}/../../../honeycomb-derived-column-validator"]

  query = {
    expression = "IF(EQUALS($http.response.status_code, 200), 1)"
  }
}

resource "honeycombio_derived_column" "dc3" {
  alias       = "sli.request_latency"
  description = "SLI: request latency less than 300ms"
  dataset     = "__all__"

  expression = data.external.example3.result.expression
}
