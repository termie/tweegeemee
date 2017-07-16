import json
import os
import sys
import urlparse

import requests

TEST_URL="https://gist.github.com/rogerallen/e21aaff6d8cea388d316#file-1_archive-edn-L1193-L1195"
TEST_URL="https://gist.github.com/rogerallen/42ca73fee10cdbf0ab41905d0c80672a#file-1_archive-edn-L1469-L1471"
#URL="https://gist.github.com/rogerallen/e21aaff6d8cea388d316#file-1_archive-edn-L1607-L1609"
#URL=os.environ.get("TGM_URL", TEST_URL)

GIST_API="https://api.github.com/gists/%s"

def main():
  url = sys.argv[1]
  o = urlparse.urlparse(url)
  frag = o.fragment
  gist_id = o.path.split("/")[-1]

  parts = frag.split("-")
  lines = [x[1:] for x in parts[-2:]]
  filename = "_".join(parts[1:-2])

  #print "frag", frag
  #print "gist", gist_id
  #print "lines", lines
  #print "file", filename

  req = requests.get(GIST_API % gist_id)

  data = req.json()

  content = data["files"]["1_archive.edn"]["content"].split("\n")
  code_line = content[int(lines[1]) - 2]
  code = code_line[9:]

  output = """
(ns tweegeemee.onetime
  (:use [clisk live])
  (:require tweegeemee.core)
  (:import [java.io File]
           [javax.imageio ImageIO])
  (:gen-class))

(defn write-png
  "write the-image to filename (which should have a .png suffix)"
  [filename the-image]
  (ImageIO/write the-image "png" (File. filename)))

(defn mymain [& args]
  (def code '%s
  )
  (write-png "output.png" (image (eval code) :width 2560 :height 1440))
)
""" % code
  print output



if __name__ == "__main__":
  main()
