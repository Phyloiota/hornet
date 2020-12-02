#!/bin/bash
del /f /q /s "snapshots\alphanet1\full_export.bin"
del /f /q /s "snapshots\alphanet1\delta_export.bin"
del /f /q /s "snapshots\alphanet2\full_export.bin"
del /f /q /s "snapshots\alphanet2\delta_export.bin"
del /f /q /s "snapshots\alphanet3\full_export.bin"
del /f /q /s "snapshots\alphanet3\delta_export.bin"
del /f /q /s "snapshots\alphanet4\full_export.bin"
del /f /q /s "snapshots\alphanet4\delta_export.bin"
mkdir "snapshots\alphanet1\"
mkdir "snapshots\alphanet2\"
mkdir "snapshots\alphanet3\"
mkdir "snapshots\alphanet4\"
go run "..\main.go" tool snapgen alphanet1 6920b176f613ec7be59e68fc68f597eb3393af80f74c7c3db78198147d5f1f92 "snapshots\alphanet1\full_export.bin"
copy "snapshots\alphanet1\full_export.bin" "snapshots\alphanet2\"
copy "snapshots\alphanet1\full_export.bin" "snapshots\alphanet3\"
copy "snapshots\alphanet1\full_export.bin" "snapshots\alphanet4\"
