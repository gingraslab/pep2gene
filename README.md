# Gene-Peptide

Match peptide search results from TPP or MSPLIT to genes.

## Progress

* ✔︎ Parse peptides from pepXML files
* Next: parse peptides from MSPLIT

## Run script

```
go run main.go parseflags.go -db="testfiles/HEK293_RefV57_cRAPgene_20130129.fasta" -enzyme="trypsin" -fdr="0.01" -file="testfiles/4745_BirAFLAG_27Sept2012_combined.pepxml" 
```
