terraform {
  required_providers {
    honeycombio = {
      source  = "honeycombio/honeycombio"
      version = "~> 0.29.0"
    }
  }
}

provider "honeycombio" {
  # Configuration options
  # set the API key using the $HONEYCOMB_API_KEY environment variable
}
