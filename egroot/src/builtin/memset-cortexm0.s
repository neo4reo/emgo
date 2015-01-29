// +build cortexm0

.syntax unified

// func Memset(s unsafe.Pointer, b byte, n uintptr)
.global builtin$Memset

// unsafe$Pointer memset(unsafe$Pointer s, byte b, uintptr n)
.global memset

.thumb_func
builtin$Memset:
.thumb_func
memset:
	cmp  r2, 0
	bne  1f

	mov  r3, r0
0:
	strb  r1, [r3]
	adds  r3, 1
	subs  r2, 1
	bne   0b
1:
	bx  lr