#!/bin/bash

# A wrapper for `go tool cgo -godefs`, to generate architecture specific
# file with information in system headers.

# Copyright (C) 2017  Laboratory of ACM/ICPC, Xidian University

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published
# by the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.

# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

# Author: Xi Ruoyao <ryxi@stu.xidian.edu.cn>

# This just works.  DO NOT use this in your project.

inputfile="$1"
outputfile=$(echo $inputfile | sed "s/_${GOOS}\\.go/.go/; s/\\.go/_${GOOS}_${GOARCH}.go/; s/^.*$/z&/")
go tool cgo -godefs $inputfile > $outputfile
rm -rf _obj
gofmt -w $outputfile
