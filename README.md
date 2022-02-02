# Micro Service Demo

To start this demo run `docker-compose up -d`. This will build the docker container. 

After they are build you can open your browser and navigate to `localhost:5000`.

## Controls
The GET input is used to retrieve a document by its author.

Underneath are the SET inputs. The Author sets the Author, the ID sets the ID field, and the template
number input is used to select the template which should be used to generate the document.

| Number | Template                                               |
|:------:|--------------------------------------------------------|
|   0    | Uses the file.template                                 |
|   1    | Uses the book.template                                 |
|   2    | Uses the pdf.template and generates a PDF via PdfLatex |
