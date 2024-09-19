module module3

go 1.23.0

replace module1 => ../module1

replace module2 => ../module2

require (
	module1 v0.0.0-00010101000000-000000000000
	module2 v0.0.0-00010101000000-000000000000
)
