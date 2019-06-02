module Factorial where

factorial :: Integer -> Integer
factorial x = if x == 0
then 1
else product [1..x]
