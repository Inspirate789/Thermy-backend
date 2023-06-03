LATEX_ARGS=--shell-escape

echo ${LATEX_ARGS}
rm *.aux *.toc *.out *.log *.bbl *.blg
xelatex ${LATEX_ARGS} main.tex
xelatex ${LATEX_ARGS} main.tex
