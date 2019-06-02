module Radians_Degrees where

degToRad :: Floating a => a -> a
degToRad x = x * (pi / 180)

radToDeg :: Floating a => a -> a
radToDeg x = x * (180 / pi)
