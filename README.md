# pep2gene

Match peptide search results from TPP or MSPLIT to genes.

* [Motivation](#motivation)
* [Rules](#rules)
* [Installation](#installation)
* [Usage](#usage)
* [Output](#output)

## Motivation

Matching peptides to the correct protein and protein isoform can be a challenging task requiring complicated rulesets that can create a confusing picture of the actual sample composition. If the principal interest is actually at the gene level, one way to simplify the interpretation of results is to match peptides to genes, since much of the complexity at the protein level is due to different transcripts arising from a single gene sequence. pep2gene was created to perform this task of matching peptides to genes.

## Rules

Peptides are matched via protein sequence to the corresponding gene identifier. A peptide that matches to multiple proteins that arise from a single gene will count as a single unique peptide match to the gene and any spectral counts for that peptide will be assigned to the gene. If a peptide matches to more than a single gene, whether the peptide gets assigned to the gene, and what portion of its spectral counts get assigned depend on the following rules.

1. If a peptide matches to multiple genes (shared peptide), and there is evidence that each of those genes is present in the sample, i.e. each gene has at least one unique peptide, then the shared peptide will be assigned to all of these matching genes and each gene will get a portion of the shared peptide's spectral counts relative to the evidence for each gene's existence. For example, if one gene has two unique peptides and another has one unique peptide, the first will get twice as many spectral counts from the shared peptide.

2. If a peptide matches to multiple genes but only a subset of those genes have unique peptides, the shared peptide will only be assigned to those genes for which there is definite evidence they are in the sample, i.e. have at least one unique peptide. The spectral counts for the shared peptide will be apportioned as in rule 1.

3. If a gene (A) matches to a peptide or peptides, but those same peptides also match to another gene (B) and that other gene has additional evidence for its existence, gene A is considered to be `subsumed` by B and will be listed as such in the output summary for gene B.

4. If two (or more) genes match to the exact same set of peptides and their is no evidence favouring the presence of one gene over the other, both genes are considered to be present and will evenly split the spectral counts of the shared peptides.

## Installation

This was built as a GO module using go1.12.7. If you have GO installed, you can build/install the binary.
Otherwise it can be run as a container, using Docker for example.

### GO executable

1. Ensure [GO](https://golang.org/doc/install) is installed.

2. Clone repo
```
git clone https://github.com/knightjdr/pep2gene.git
cd pep2gene
```

3. Build executable
```
go build
```

The executable will be called `pep2gene`.

### Docker

#### Pull image (and rename - optional)

```
docker pull knightjdr/pep2gene:v1.2.0
docker tag knightjdr/pep2gene:v1.2.0 pep2gene
```

Check for [versions](https://cloud.docker.com/repository/registry-1.docker.io/knightjdr/pep2gene/tags).

#### Build

1. Clone repo
```
git clone https://github.com/knightjdr/pep2gene.git
cd pep2gene
```

2. Build the image

For Docker:

```
docker build -t pep2gene -f docker/standard/Dockerfile .
```

For Singularity:

```
docker build -t pep2genesing -f docker/singularity/Dockerfile .
```

We do not provide a Singularity definition file but the Docker image can be used with Singularity provided it is built from the correct source. The Dockerfile found in `docker/standard/` was designed for Docker itself. While the image is small (~7mb), it does not work with Singularity. The Dockerfile found in `docker/singularity/` will build an image compatibly with Singularity although it is about twice the size (13MB).

The images are also hosted at DockerHub in separate repos: [Docker](https://cloud.docker.com/repository/docker/knightjdr/pep2gene) and [Singularity](https://cloud.docker.com/repository/docker/knightjdr/pep2genesing).

## Usage

### GO executable

```
pep2gene -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin"
```

### Docker

```
docker run -v $(pwd):/files/ pep2gene -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin"
```

### Singularity

```
singularity run -B ./:/files/ docker://knightjdr/pep2genesing:v1.2.0 -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin"
```

The database and peptide file must be located in the working directory Docker/Singularity is called from. Relative or nested paths will not work, i.e. `./some-directory/database.fasta` or `../database.fasta`. The output file will also be written to the working directory.

### Flags

| Name | Description | Required | Default |
|------|-------------|----------|---------|
| -db | FASTA database | true | |
| -enzyme | digestion enzyme | false | |
| -fdr | MSPLIT peptide FDR | false | 0.01 |
| -file | peptide file | true | |
| -ignoreinvalid | ignore sequences with an invalid header | false | true |
| -inferenzyzme | infer the digestive enzyme | false | false |
| -missedcleavages | number of missed cleavages | false | 0 |
| -output | output file format | false | json |
| -pepprob | TPP peptide probability | false | 0.85 |
| -pipeline | search pipeline | false | TPP |

#### Notes

**_-db (database)_**

The search database is expected to be in FASTA format, with headers containing the following string
> gn|\<gene symbol>:\<Entrez gene ID>

E.G:

> \>gi|22538794|gn|PDCD10:11235| programmed cell death protein 10 [Homo sapiens]

**_-enzyme_**

If an enzyme is specified, the sequence database will be digested before peptide matching begins. This significantly speeds up the matching process. If no enzyme is used, peptides are matched against the any protein subsequence.

The available enzymes are:
* arg-c
* asp-n
* asp-n_ambic
* chymotrypsin
* cnbr
* lys-c
* lys-c/p
* lys-n
* pepsina
* trypchymo
* trypsin
* trypsin/p
* v8-de
* v8-e

**_-fdr_**

The FDR is used for parsing high-quality peptides from MSPLIT results, both DDA and DIA. It is ignored when parsing TPP results.

**_-file_**

pepXML files from TPP are supported, as are DDA and DIA output files from MSPLIT.

**_-inferenzyme_**

pep2gene can infer the enzyme used to digest the sample, rather that requiring it to be input as an argument. However, currently
the enzyme name can only be parsed from pepXML files that contain the `sample_enzyme` field:

> <sample_enzyme name="trypsin">

The name of the enzyme must match one of the names listed above.

**_-ignoreinvalid_**

Sequences that do not conform to the required header format

> gn|\<gene symbol>:\<Entrez gene ID>

will be ignored by default since pep2gene will not know how to parse the gene symbol and gene ID, both of which are required.
This can be overridden by setting this argument to `false`. When this argument is set to false, any sequences for which a symbol
and ID can not be determined will be identified by any leading non-whitespace characters in the header, and will be prefixed with
`p-` to indicate they do not conform.  

**_-output_**

Results can be output in either json (default) or txt format. The txt format is a legacy format that we do not recommend using. See the [Output](#output) section for a detailed description of each format.

**_-pepprob_**

The peptide probability for parsing high-quality peptides from TPP results. It is ignored when parsing MSPLIT results.

**_-pipeline_**

The analysis pipeline used for searching peptides. The options are:
* MSPLIT_DDA
* MSPLIT_DIA
* TPP

## Output

### json

The json format will contain fields for user-supplied command line arguments, for example
the database and file names, and a `genes` object indexed by gene ID for each gene
identified in the sample.

#### gene fields

| gene field | definition |
|------------|------------|
| name | gene name/symbol |
| peptides | peptides assigned to the gene |
| sharedIDs | any other genes (by ID) it shares peptides with |
| sharedNames | any other genes (by name) it shares peptides with |
| spectralCount | total spectral count for the gene |
| subsumed | subsumed genes |
| unique | peptides unique to the gene |
| uniqueShared | peptides unique to the gene group, if the gene shares peptides |

#### peptide fields

| peptide field | definition |
|---------------|------------|
| allottedSpectralCount | the portion of the peptide's spectral count allotted to the gene |
| totalSpectralCount | the total spectral count for the peptide in the sample |
| unique | a boolean indicating if the peptide is unique to the gene |
| uniqueShared | a boolean indicating if the peptide is unique to the group the gene shares peptides with |

```
{
  "database": "database.fasta",
  "enzyme": "trypsin",
  "file": "sample.pepxml",
  "genes": {
    "5825": {
      "name": "ABCD3",
      "peptides": {
        "DQVIYPDGR": {
          "allottedSpectralCount": 1,
          "totalSpectralCount": 1,
          "unique": true,
          "uniqueShared": false
        },
        "FDHVPLATPN[115]GDVLIR": {
          "allottedSpectralCount": 1,
          "totalSpectralCount": 1,
          "unique": true,
          "uniqueShared": false
        }
      },
      "sharedIDs": [],
      "sharedNames": [],
      "spectralCount": 2,
      "subsumed": [],
      "unique": 2,
      "uniqueShared": 0
    },
    "60": {
      "name": "ACTB",
      "peptides": {
        "AGFAGDDAPR": {
          "allottedSpectralCount": 2.5,
          "totalSpectralCount": 5,
          "unique": false,
          "uniqueShared": true
        },
        "DLTDYLMK": {
          "allottedSpectralCount": 2.5,
          "totalSpectralCount": 5,
          "unique": false,
          "uniqueShared": false
        }
      },
      "sharedIDs": ["71"],
      "sharedNames": ["ACTG1"],
      "spectralCount": 5,
      "subsumed": ["100996820", "345651", "445582", "653269", "653781", "728378"],
      "unique": 0,
      "uniqueShared": 1
    }
  }
}
````

### txt

The txt format contains less information than the json format and is not recommended.

The first two lines are headers, followed by gene entries separated by newlines. The first header line
contains the keys for the summary line of each hit. In the example below the `HitNumber` for the first hit
is Hit_1, the `Gene` is ABCD3, the `GeneID` is 5825, the `SpectralCount` is 4.00, the number of `Unique`
peptides is 4 and there are no `Subsumed` genes for the hit. Since spectral counts for peptides can
be divided between genes, the spectral count is reported as a floating-point number.

The second gene entry is for a shared group, i.e. the members or this group perfectly share a
set of peptides: in this example ACTB and ACTG1, corresponding to the gene IDs 60 and 71 respectively.
This group subsumes several other genes indicated by their IDs.

The summary line for each hit is followed by its assigned peptides. Each peptide has a `TotalSpectralCount`
referring to the total number of spectral counts detected for it in the sample and a yes/no indicator to
declare its uniqueness to the gene hit.

```
HitNumber;;Gene;;GeneID;;SpectralCount;;Unique;;Subsumed
Peptide;;TotalSpectralCount;;IsUnique

Hit_1;;ABCD3;;5825;;4.00;;4;;
DQVIYPDGR;;1;;yes
FDHVPLATPN[115]GDVLIR;;1;;yes
IANPDQLLTQDVEK;;1;;yes
ITELMQVLK;;1;;yes

Hit_2;;ACTB, ACTG1;;60, 71;;8.56;;0;;100996820, 345651, 445582, 653269, 653781, 728378
AGFAGDDAPR;;2;;no
DLTDYLMK;;2;;no
DLYANTVLSGGTTMYPGIADR;;3;;no
DLYANTVLSGGTTM[147]YPGIADR;;1;;no
DSYVGDEAQSK;;2;;no
EITALAPSTMK;;1;;no
```