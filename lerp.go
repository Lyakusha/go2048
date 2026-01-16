package main

func lerp(from float32, to float32, t float32) float32 {
	return (1-t)*from + t*to
}
