#!/bin/bash -eu

bin/practice_1.10 $(curl "http://www.alexa.com/topsites" | \
  pup "div.td.DescriptionCell > p > a attr{href}" | \
  xargs -IXXX echo "http://www.alexa.com/topsitesXXX")
