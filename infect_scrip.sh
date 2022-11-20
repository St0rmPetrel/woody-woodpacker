#!/bin/bash

VICTIM_NAME="victim"
PARASITE_NAME="parasit"

echo '
#include <stdio.h>

int main() {
    printf("Hello World\n");
}
' | gcc -x c - -o $VICTIM_NAME

xxd -g 1 $VICTIM_NAME $VICTIM_NAME.dump

echo '
    global _start

	section .text

_start:
	push 0x0a293b    ; const char* rsp = ";)\n"
	mov rdi, 1
	mov rsi, rsp     ; string on stack
	mov rdx, 3       ; number of chars
	mov rax, 1       ; write syscall
    syscall
	pop rax

	push 0x401000 ; tricky jump
	ret           ;
' > $PARASITE_NAME.asm

nasm -f elf64 $PARASITE_NAME.asm

xxd -g 1 $PARASITE_NAME.o $PARASITE_NAME.dump

rm $PARASITE_NAME.o $PARASITE_NAME.asm
