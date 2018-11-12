# tinyVM
tinyVM

## cpu specifications
* little endian
* 32 bit instuction code

### registers
* $1
* $2
* $3
* $4
* $5
* $6
* sp - stackpointer
* ip - instruction pointer

## instcution code
* 32bit code  
* opcode 1byte(8bit)
* regs and immediate value(28bit)

## opcodes
### HLT
bytecode = 0

### MOV DST, SRC
bytecode = 1  
DST should be a register.  
SRC should be a register or immediate value.

### ADD DST, SRC1, SRC2
bytecode = 2  
DST should be a register.  
SRC1 should be a register or immediate value.  
SRC2 should be a register or immediate value.  

### SUB DST, SRC1, SRC2
bytecode = 3  
DST should be a register.  
SRC1 should be a register or immediate value.  
SRC2 should be a register or immediate value.  
