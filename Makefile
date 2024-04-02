LIB := libtest.a
BIN := main
HDRS := $(wildcard *.h)
SRCS := $(wildcard *.cpp)
OBJS := $(SRCS:.cpp=.o)

CC := gcc
CFLAGS := -fPIC -g -Wall -O2 -L. -L/usr/local/lib
CXX := g++
CXXFLAGS := -fPIC -g -Wall -O2 -std=c++20
AR := ar

all: lib bin

bin: $(BIN)

lib: $(LIB)

$(BIN): $(LIB) main.c
	$(CC) $(CFLAGS) main.c -o $(BIN) -ltest -lstdc++ -lngt

$(LIB): $(OBJS)
	$(AR) rvs $(LIB) $(OBJS)

.cpp.o:
	$(CXX) $(CXXFLAGS) -c $<

.PHONY: clean
clean:
	$(RM) $(BIN) $(LIB) $(OBJS)

runc: $(BIN)
	./$(BIN)

rungo: $(LIB)
	go run main.go

runrust: $(LIB)
	cargo run

runall: runc rungo runrust
