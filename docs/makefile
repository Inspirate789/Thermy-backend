LATEX_ARGS := -shell-escape

.PHONY: build clean_trash clean

explanatory_note.pdf: tex/build.sh tex/*.tex tex/*.bib tex/*.sty tex/*.bst tex/*.cls tex/img/* # tex/*.inc tex/*.layout tex/*.module tex/*.clo tex/listings/*
	make build
	make clean_trash

build:
	cd tex && sh ./build.sh && cd ..
	pdftk A=pdf/Title_Sapozhkov.pdf B=pdf/SW_Sapozhkov.pdf C=tex/main.pdf cat A B C output pdf/explanatory_note.pdf
	rm -f tex/main.pdf

clean_trash:
	rm -f tex/*.aux tex/*.toc tex/*.out tex/*.log tex/*.bbl tex/*.blg

clean:
	make clean_trash
	rm -f explanatory_note.pdf
