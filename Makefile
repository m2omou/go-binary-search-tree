### Makefile --- 
## 
## Filename: Makefile
## Author: Mourad sabour
## Created: Jul 6 16:37:24 2014 (-0700)
## 

NAME=		bst

BIN=		go

RM=		rm -rf

SRC=		main.go\
		bst.go\
		delete.go\
		exists.go\
		insert.go\
		show.go

all:
		$(BIN) build -o $(NAME) $(SRC)

clean:		
		$(RM) $(NAME) *~ *#

re:		
		make clean
		make all
